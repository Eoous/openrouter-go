package openroutergo

import (
	"github.com/go-resty/resty/v2"
)

const (
	baseUrl = "https://openrouter.ai/api/v1"
)

type Client struct {
	resty *resty.Client
}

func NewClient() *Client {
	client := resty.New().SetBaseURL(baseUrl)
	//client.SetHeader("Authorization", "Bearer "+apiKey)

	return &Client{
		resty: client,
	}
}

func (c *Client) WithAuth(key string) *Client {
	c.resty.SetHeader("Authorization", "Bearer "+key)

	return c
}

// WithSiteUrl sets site url for rankings on openrouter.ai.
func (c *Client) WithSiteUrl(url string) *Client {
	c.resty.SetHeader("HTTPReferer", url)

	return c
}

// WithSiteName sets site name for rankings on openrouter.ai.
func (c *Client) WithSiteName(name string) *Client {
	c.resty.SetHeader("X-Title", name)

	return c
}

func (c *Client) Completion() (*Client, error) {
	a :=
		`{
			"model": "google/gemini-2.5-flash-image-preview",
			"messages": [
				{
					"role": "user",
					"content": "Generate a beautiful sunset over mountains"
				}
			],
			"modalities": ["image", "text"]
	}`

	r, err := c.resty.R().SetBody(a).Post("/chat/completions")
	if err != nil {
		return nil, err
	}

	println(r.String())

	return c, nil
}
