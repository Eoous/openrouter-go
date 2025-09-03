package openroutergo

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

func (r Role) New(content string) Message {
	return Message{
		Role:    r,
		Content: content,
	}
}

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type OpenRouterParams struct {
	Model      string    `json:"model"`
	Messages   []Message `json:"messages"`
	Modalities []string  `json:"modalities"`
}

type OpenRouterImage struct {
	Type     string `json:"type"`
	ImageUrl struct {
		Url string `json:"url"`
	} `json:"image_url"`
	Index int `json:"index"`
}

type OpenRouterMessage struct {
	Message

	Refusal   interface{}       `json:"refusal"`
	Reasoning interface{}       `json:"reasoning"`
	Images    []OpenRouterImage `json:"images"`
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
		Logprobs           interface{}       `json:"logprobs"`
		FinishReason       string            `json:"finish_reason"`
		NativeFinishReason string            `json:"native_finish_reason"`
		Index              int               `json:"index"`
		Message            OpenRouterMessage `json:"message"`
	} `json:"choices"`

	Usage OpenRouterUsage `json:"usage"`
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
