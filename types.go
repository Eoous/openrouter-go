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
