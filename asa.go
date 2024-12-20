package asa

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ropon/requests/v2"
	"net/url"
	"time"
)

const (
	defaultBaseURL        = "https://api.searchads.apple.com/api/v5/"
	defaultAuthURL        = "https://appleid.apple.com/auth/oauth2/token"
	defaultTimeout        = 30 * time.Second
	defaultExpireDuration = 180 * 24 * time.Hour
	userAgent             = "apple-search-api-go"
)

// Client 客户端
type Client struct {
	auth   *TokenConfig
	client *requests.Request
	common service

	Campaigns         *CampaignService
	AdGroups          *AdGroupService
	Reporting         *ReportingService
	Keywords          *KeywordService
	AccessControlList *AccessControlListService
}

// service 服务
type service struct {
	client *Client
}

// SetDefault 设置默认请求配置
func SetDefault(req *requests.Request) {
	_ = req.SetBaseUrl(defaultBaseURL)
	req.SetTimeout(defaultTimeout)
	req.SetHeader("User-Agent", userAgent)
}

// NewClient 创建客户端
func NewClient(httpReq *requests.Request, accessToken ...string) *Client {
	if httpReq == nil {
		httpReq = requests.New()
	}
	SetDefault(httpReq)
	c := &Client{
		client: httpReq,
	}
	if len(accessToken) > 0 {
		c.client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken[0]))
	}
	c.common.client = c
	c.Campaigns = (*CampaignService)(&c.common)
	c.AdGroups = (*AdGroupService)(&c.common)
	c.Reporting = (*ReportingService)(&c.common)
	c.Keywords = (*KeywordService)(&c.common)
	c.AccessControlList = (*AccessControlListService)(&c.common)
	return c
}

func NewClientWithAuth(auth *TokenConfig) *Client {
	c := &Client{
		auth: auth,
	}

	c.common.client = c
	c.Campaigns = (*CampaignService)(&c.common)
	c.AdGroups = (*AdGroupService)(&c.common)
	c.Reporting = (*ReportingService)(&c.common)
	c.Keywords = (*KeywordService)(&c.common)
	c.AccessControlList = (*AccessControlListService)(&c.common)
	return c
}

// SetHTTPTimeout 设置http请求超时时间
func (c *Client) SetHTTPTimeout(n time.Duration) {
	c.client.SetTimeout(n)
}

// SetHTTPDebug 设置http请求debug
func (c *Client) SetHTTPDebug(flag bool) {
	c.client.Debug = flag
}

// SetHTTPProxy 设置http请求代理
func (c *Client) SetHTTPProxy(proxyUrl string) {
	c.client.SetProxy(proxyUrl)
}

// SetOrgID 设置组织ID
func (c *Client) SetOrgID(orgID int64) {
	c.auth.SetOrgID(orgID)
}

func (c *Client) HttpClient() (*requests.Request, error) {
	httpReq, err := c.auth.Client()
	if err != nil {
		return nil, err
	}
	httpReq.Debug = true
	SetDefault(httpReq)
	return httpReq, nil
}

// rawJson 处理json响应
func (c *Client) rawJson(res *requests.Response, resp interface{}) error {
	if res.Json().Get("code").Int64() != 0 {
		return errors.New(res.Json().Get("message").String())
	}
	err := res.RawJson(resp)
	if err != nil {
		return err
	}
	return nil
}

// get 处理get请求
func (c *Client) get(apiUrl string, resp interface{}, params ...interface{}) error {
	if len(params) > 0 {
		param := params[0]
		// 构建 URL
		u, err := url.Parse(apiUrl)
		if err != nil {
			return err
		}

		query := u.Query()
		err = addParamsToQuery(query, param)
		if err != nil {
			return err
		}
		u.RawQuery = query.Encode()
		apiUrl = u.String()
	}
	client, err := c.HttpClient()
	if err != nil {
		return err
	}
	res, err := client.Get(apiUrl)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

// post 处理post请求
func (c *Client) post(url string, resp interface{}, data ...interface{}) error {
	client, err := c.HttpClient()
	if err != nil {
		return err
	}
	if len(data) > 0 {
		postData := data[0]
		bS, err := json.Marshal(postData)
		if err != nil {
			return err
		}
		res, err := client.Post(url, string(bS))
		if err != nil {
			return err
		}
		return c.rawJson(res, resp)
	}
	res, err := client.Post(url, data...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}
