package asa

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"log"
	"time"

	"github.com/ropon/requests/v2"
)

// ErrMissingPEM happens when the bytes cannot be decoded as a PEM block.
var ErrMissingPEM = errors.New("no PEM blob found")

// ErrInvalidPrivateKey happens when a key cannot be parsed as a ECDSA PKCS8 private key.
var ErrInvalidPrivateKey = errors.New("key could not be parsed as a valid ecdsa.PrivateKey")

// ErrHTTPTokenBadRequest happens when apple generate token http request failed.
var ErrHTTPTokenBadRequest = errors.New("generate auth token failed with")

// TokenConfig 获取token配置
type TokenConfig struct {
	jwtGenerator *standardJWTGenerator
	httpReq      *requests.Request
	orgID        int64
}

type accessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`

	expiresAfter time.Time
}

type standardJWTGenerator struct {
	keyID          string
	issuerID       string
	clientID       string
	expireDuration time.Duration
	privateKey     *ecdsa.PrivateKey
	accessToken    *accessToken
	token          string
	httpReq        *requests.Request
}

func (g *standardJWTGenerator) AccessToken() (string, error) {
	if g.IsAccessTokenValid() {
		return g.accessToken.AccessToken, nil
	}

	token, err := g.Token()
	if err != nil {
		return "", err
	}

	accessTkn, err := g.generateAccessToken(token)
	if err != nil {
		return "", err
	}

	g.accessToken = accessTkn

	return accessTkn.AccessToken, nil
}

func (g *standardJWTGenerator) IsAccessTokenValid() bool {
	if g.accessToken == nil || g.accessToken.AccessToken == "" {
		return false
	}

	if g.accessToken.expiresAfter.Before(time.Now()) {
		return false
	}

	return true
}

func (g *standardJWTGenerator) generateAccessToken(token string) (*accessToken, error) {
	url := fmt.Sprintf("%s/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=searchadsorg", defaultAuthURL, g.clientID, token)
	res, err := requests.Post(url)

	if err != nil {
		return nil, err
	}

	accToken := &accessToken{}
	err = res.RawJson(accToken)
	if err != nil {
		return nil, err
	}
	accToken.expiresAfter = time.Now().Add(time.Second * time.Duration(accToken.ExpiresIn))
	return accToken, nil
}

func (g *standardJWTGenerator) Token() (string, error) {
	if g.IsTokenValid() {
		return g.token, nil
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, g.claims())
	t.Header["kid"] = g.keyID

	token, err := t.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	g.token = token

	return token, nil
}

func (g *standardJWTGenerator) IsTokenValid() bool {
	if g.token == "" {
		return false
	}

	parsed, err := jwt.Parse(
		g.token,
		jwt.KnownKeyfunc(jwt.SigningMethodES256, g.privateKey),
		jwt.WithAudience("https://appleid.apple.com"),
		jwt.WithIssuer(g.issuerID),
	)
	if err != nil {
		return false
	}

	return parsed.Valid
}

func (g *standardJWTGenerator) claims() jwt.Claims {
	expiry := time.Now().Add(g.expireDuration)

	return jwt.StandardClaims{
		Audience:  jwt.ClaimStrings{"https://appleid.apple.com"},
		Subject:   g.clientID,
		Issuer:    g.issuerID,
		ExpiresAt: jwt.At(expiry),
	}
}

// NewTokenConfig 创建token配置
func NewTokenConfig(clientID, teamID, keyID, privateKey string) (*TokenConfig, error) {
	key, err := parsePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	gen := &standardJWTGenerator{
		keyID:          keyID,
		issuerID:       teamID,
		clientID:       clientID,
		privateKey:     key,
		expireDuration: defaultExpireDuration,
	}
	return &TokenConfig{
		jwtGenerator: gen,
		httpReq:      requests.New(),
	}, nil
}

// SetHTTPDebug 设置http请求debug
func (t *TokenConfig) SetHTTPDebug(flag bool) {
	t.httpReq.Debug = flag
}

// SetHTTPProxy 设置http请求代理
func (t *TokenConfig) SetHTTPProxy(proxyUrl string) {
	t.httpReq.SetProxy(proxyUrl)
}

func (t *TokenConfig) SetOrgID(orgID int64) {
	t.orgID = orgID
}

// GenerateClientSecret 生成client secret https://developer.apple.com/documentation/apple_search_ads/implementing_oauth_for_the_apple_search_ads_api
func (t *TokenConfig) GenerateClientSecret() (string, error) {
	return t.jwtGenerator.Token()
}

func (t *TokenConfig) Client() *requests.Request {
	tokenStr, err := t.jwtGenerator.AccessToken()
	if err != nil {
		log.Println("TokenConfig Client error:", err)
		return nil
	}
	t.httpReq.SetHeader("Authorization", fmt.Sprintf("Bearer %s", tokenStr))
	if t.orgID > 0 {
		t.httpReq.SetHeader("X-AP-Context", fmt.Sprintf("orgId=%v", t.orgID))
	}
	return t.httpReq
}
