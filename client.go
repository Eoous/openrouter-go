package openroutergo

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"

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

// WithMaxTokens sets the maximum number of tokens to generate.
func (c *Client) WithMaxTokens(maxTokens int) *Client {
	c.params.MaxTokens = maxTokens

	return c
}

// WithTemperature sets the temperature for the model.
func (c *Client) WithTemperature(t float64) *Client {
	c.params.Temperature = t

	return c
}

// WithStream enables or disables streaming mode.
func (c *Client) WithStream(stream bool) *Client {
	c.params.Stream = stream
	c.client.SetDoNotParseResponse(stream)

	return c
}

// validate checks if the client is properly configured.
func (c *Client) validate() error {
	if c.params.Model == "" {
		return errors.New("model is required")
	}

	return nil
}

func readOpenRouterResponse(data []byte) (*OpenRouterResponse, error) {
	var resp openRouterResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	if resp.OpenRouterResponseErr != nil {
		return nil, errors.New(resp.Error.Message)
	}

	return resp.OpenRouterResponse, nil
}

func (c *Client) ChatCompletions(msgs []Msg) (*OpenRouterResponse, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	data := request{
		Params:   c.params,
		Messages: msgs,
	}
	resp, err := c.client.R().
		SetBody(data).
		Post("/chat/completions")
	if err != nil {
		return nil, err
	}

	if c.params.Stream {
		buf := bufio.NewReader(resp.RawBody())
		for {
			chunk := make([]byte, 4096)
			n, err := buf.Read(chunk)
			if n > 0 {
				// todo: handle stream response
			}

			if err != nil {
				if err == io.EOF {
					// todo: last chunk
				}

				break
			}
		}

		return nil, nil
	}

	var orp openRouterResponse
	err = json.Unmarshal(resp.Body(), &orp)
	if err != nil {
		return nil, err
	}

	if orp.OpenRouterResponseErr != nil {
		return nil, errors.New(orp.Error.Message)
	}

	return orp.OpenRouterResponse, nil
}

type openRouterResponse struct {
	*OpenRouterResponse
	*OpenRouterResponseErr
}
