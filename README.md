# openrouter-go

API Client in golang for [OpenRouter](https://openrouter.ai/)

## Example

```go
c := NewClient()
c.WithAuth("your-api-key")
c.WithModel("deepseek/deepseek-r1")

msgs := []Msg{RoleUser.New("Do you know openrouter.ai?")}
completions, err := c.ChatCompletions(msgs)
```
