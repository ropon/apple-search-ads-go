package asa

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/ropon/requests/v2"
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
	App               *AppService
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
	req.SetHeader("Content-Type", "application/json")
}

// NewClient 创建客户端
func NewClient(httpClient interface{}, accessToken ...string) *Client {
	var c *Client

	switch v := httpClient.(type) {
	case *requests.Request:
		req := v
		if req == nil {
			req = requests.New()
		}
		SetDefault(req)
		c = &Client{
			client: req,
		}
		if len(accessToken) > 0 {
			c.client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken[0]))
		}

	case *TokenConfig:
		if v == nil {
			return nil
		}
		httpReq, err := v.Client()
		if err != nil {
			return nil
		}
		c = &Client{
			auth:   v,
			client: httpReq,
		}

	case nil:
		if len(accessToken) == 0 {
			return nil
		}
		req := requests.New()
		SetDefault(req)
		req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken[0]))
		c = &Client{
			client: req,
		}

	default:
		return nil
	}

	// 初始化各个服务
	c.common.client = c
	c.Campaigns = (*CampaignService)(&c.common)
	c.AdGroups = (*AdGroupService)(&c.common)
	c.Reporting = (*ReportingService)(&c.common)
	c.Keywords = (*KeywordService)(&c.common)
	c.AccessControlList = (*AccessControlListService)(&c.common)
	c.App = (*AppService)(&c.common)

	return c
}

// SetHTTPTimeout 设置http请求超时时间
func (c *Client) SetHTTPTimeout(n time.Duration) error {
	client, err := c.HttpClient()
	if err != nil {
		return err
	}
	client.SetTimeout(n)
	return nil
}

// SetHTTPDebug 设置http请求debug
func (c *Client) SetHTTPDebug(flag bool) error {
	client, err := c.HttpClient()
	if err != nil {
		return err
	}
	client.Debug = flag
	return nil
}

// SetHTTPProxy 设置http请求代理
func (c *Client) SetHTTPProxy(proxyUrl string) error {
	client, err := c.HttpClient()
	if err != nil {
		return err
	}
	client.SetProxy(proxyUrl)
	return nil
}

// SetOrgID 设置组织ID
func (c *Client) SetOrgID(orgID int64) error {
	// 如果有 auth 配置,直接设置
	if c.auth != nil {
		c.auth.SetOrgID(orgID)
		return nil
	}

	// 如果没有 auth 配置,但有 client
	if c.client != nil {
		// 更新或添加 X-AP-Context header
		c.client.SetHeader("X-AP-Context", fmt.Sprintf("orgId=%v", orgID))
		return nil
	}

	return errors.New("client not initialized")
}

// HttpClient 获取请求客户端
func (c *Client) HttpClient() (*requests.Request, error) {
	// 如果client为空,需要初始化
	if c.client == nil {
		// 优先使用TokenConfig方式
		if c.auth != nil {
			client, err := c.auth.Client()
			if err != nil {
				return nil, err
			}
			c.client = client
		} else {
			// 默认初始化
			c.client = requests.New()
		}
	}

	// 设置默认配置
	SetDefault(c.client)

	// 如果有auth配置,需要设置认证头
	if c.auth != nil {
		token, err := c.auth.jwtGenerator.AccessToken()
		if err != nil {
			return nil, err
		}
		c.client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
		// auth方式的orgID通过auth对象获取
		c.client.SetHeader("X-AP-Context", fmt.Sprintf("orgId=%d", c.auth.orgID))
	}
	// 注意:非auth方式的orgID已经在SetOrgID中设置了header,这里不需要重复设置
	return c.client, nil
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

// put 处理put请求
func (c *Client) put(url string, resp interface{}, data ...interface{}) error {
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
		res, err := client.Put(url, string(bS))
		if err != nil {
			return err
		}
		return c.rawJson(res, resp)
	}
	res, err := client.Put(url, data...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}

// delete 处理delete请求
func (c *Client) delete(url string, resp interface{}, data ...interface{}) error {
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
		res, err := client.Delete(url, string(bS))
		if err != nil {
			return err
		}
		return c.rawJson(res, resp)
	}
	res, err := client.Delete(url, data...)
	if err != nil {
		return err
	}
	return c.rawJson(res, resp)
}
