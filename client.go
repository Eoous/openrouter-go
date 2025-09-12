package openroutergo

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

const (
	baseUrl = "https://openrouter.ai/api/v1"

	defaultUserAgent = "openrouter-go (https://github.com/Eoous/openrouter-go)"
	defaultSiteName  = "openrouter-go"
	defaultSiteUrl   = "https://github.com/Eoous/openrouter-go)"
)

type Client struct {
	client *resty.Client
	params Params
}

func NewClient() *Client {
	client := resty.New().
		SetBaseURL(baseUrl).
		SetHeader("User-Agent", defaultUserAgent)

	c := Client{client: client}
	return c.WithSiteName(defaultSiteName).WithSiteUrl(defaultSiteUrl)
}

func (c *Client) WithAuth(key string) *Client {
	c.client.SetHeader("Authorization", "Bearer "+key)

	return c
}

// WithSiteUrl sets site url for rankings on openrouter.ai.
func (c *Client) WithSiteUrl(url string) *Client {
	c.client.SetHeader("HTTP-Referer", url)

	return c
}

// WithSiteName sets site name for rankings on openrouter.ai.
func (c *Client) WithSiteName(name string) *Client {
	c.client.SetHeader("X-Title", name)

	return c
}

// WithModel sets the model to use.
func (c *Client) WithModel(model string) *Client {
	c.params.Model = model

	return c
}

// WithStream enables or disables streaming mode.
func (c *Client) WithStream(stream bool) *Client {
	c.params.Stream = stream
	c.client.SetDoNotParseResponse(stream)

	return c
}

func (c *Client) ChatCompletions(msgs []Message) (*OpenRouterResponse, error) {
	body := request{
		Params:   c.params,
		Messages: msgs,
	}
	r, err := c.client.R().
		SetBody(body).
		Post("/chat/completions")
	if err != nil {
		return nil, err
	}

	// todo: handle stream response
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
