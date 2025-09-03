package openroutergo

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

const (
	baseUrl = "https://openrouter.ai/api/v1"

	defaultUserAgent = "openrouter-go (https://github.com/Eoous/openrouter-go)"
	defaultSiteName  = "openrouter-go)"
	defaultSiteUrl   = "https://github.com/Eoous/openrouter-go)"
)

type Client struct {
	Client *resty.Client
}

func NewClient() *Client {
	client := resty.New().
		SetBaseURL(baseUrl).
		SetHeader("User-Agent", defaultUserAgent)

	return (&Client{Client: client}).
		WithSiteName(defaultSiteName).
		WithSiteUrl(defaultSiteUrl)
}

func (c *Client) WithAuth(key string) *Client {
	c.Client.SetHeader("Authorization", "Bearer "+key)

	return c
}

// WithSiteUrl sets site url for rankings on openrouter.ai.
func (c *Client) WithSiteUrl(url string) *Client {
	c.Client.SetHeader("HTTP-Referer", url)

	return c
}

// WithSiteName sets site name for rankings on openrouter.ai.
func (c *Client) WithSiteName(name string) *Client {
	c.Client.SetHeader("X-Title", name)

	return c
}

func (c *Client) Completion(param OpenRouterParams) (*OpenRouterResponse, error) {
	r, err := c.Client.R().
		SetBody(param).
		Post("/chat/completions")
	if err != nil {
		return nil, err
	}

	var resp openRouterResponse
	err = json.Unmarshal(r.Body(), &resp)
	if err != nil {
		return nil, err
	}

	if resp.OpenRouterResponseErr != nil {
		return nil, errors.New(resp.Error.Message)
	}

	return resp.OpenRouterResponse, nil
}

type openRouterResponse struct {
	*OpenRouterResponse
	*OpenRouterResponseErr
}
