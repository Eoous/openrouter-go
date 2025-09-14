package openroutergo

import (
	"errors"
)

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleDeveloper Role = "developer"
	RoleTool      Role = "tool"
)

func (r Role) New(content string) Msg {
	return Msg{
		Role:    r,
		Content: content,
	}
}

type Msg struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type Params struct {
	Model string `json:"model"`

	// Alternate list of models for routing overrides.
	Models []string `json:"models,omitempty"`
	// Provider
	// Reasoning
	// Usage
	// Transforms
	Stream      bool    `json:"stream,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Seed        int     `json:"seed,omitempty"`

	TopP float64 `json:"top_p,omitempty"`
	TopK float64 `json:"top_k,omitempty"`
	TopA float64 `json:"top_a,omitempty"`
	MinP float64 `json:"min_p,omitempty"`

	FrequencyPenalty  float64 `json:"frequency_penalty,omitempty"`
	PresencePenalty   float64 `json:"presence_penalty,omitempty"`
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty"`

	// LogitBias
	TopLogprobs int    `json:"top_logprobs,omitempty"`
	User        string `json:"user,omitempty"`

	Modalities []string `json:"modalities,omitempty"`
}

type request struct {
	// Must contain model.
	Params
	Messages []Msg `json:"messages"`
}

type OpenRouterImages struct {
	Type     string `json:"type"`
	ImageUrl struct {
		Url string `json:"url"`
	} `json:"image_url"`
	Index int `json:"index"`
}

type OpenRouterMsg struct {
	Msg

	Refusal          interface{}   `json:"refusal,omitempty"`
	Reasoning        string        `json:"reasoning,omitempty"`
	ReasoningDetails []interface{} `json:"reasoning_details,omitempty"`

	Images []OpenRouterImages `json:"images,omitempty"`
}

type OpenRouterUsage struct {
	PromptTokens        int `json:"prompt_tokens"`
	CompletionTokens    int `json:"completion_tokens"`
	TotalTokens         int `json:"total_tokens"`
	PromptTokensDetails struct {
		CachedTokens int `json:"cached_tokens"`
	} `json:"prompt_tokens_details"`
	CompletionTokensDetails struct {
		ReasoningTokens int `json:"reasoning_tokens"`
		ImageTokens     int `json:"image_tokens"`
	} `json:"completion_tokens_details"`
}

type OpenRouterResponse struct {
	Id       string `json:"id"`
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Object   string `json:"object"`
	Created  int64  `json:"created"`
	Choices  []struct {
		Logprobs           interface{}   `json:"logprobs"`
		FinishReason       string        `json:"finish_reason"`
		NativeFinishReason string        `json:"native_finish_reason"`
		Index              int           `json:"index"`
		Message            OpenRouterMsg `json:"message"`
	} `json:"choices"`

	Usage OpenRouterUsage `json:"usage"`
}

type Message interface {
	Message() (*Msg, error)
	Image() (*OpenRouterImages, error)
}

func (r *OpenRouterResponse) Message() (*Msg, error) {
	if len(r.Choices) == 0 {
		return nil, errors.New("no choices found")
	}

	return &r.Choices[0].Message.Msg, nil
}

func (r *OpenRouterResponse) Image() (*OpenRouterImages, error) {
	if len(r.Choices) == 0 {
		return nil, errors.New("no choices found")
	}

	if len(r.Choices[0].Message.Images) == 0 {
		return nil, errors.New("no images found")
	}

	return &r.Choices[0].Message.Images[0], nil
}

type OpenRouterResponseErr struct {
	Error struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		Metadata struct {
			Raw          string `json:"raw"`
			ProviderName string `json:"provider_name"`
		} `json:"metadata,omitempty"`
	} `json:"error,omitempty"`
	UserId string `json:"user_id,omitempty"`
}
