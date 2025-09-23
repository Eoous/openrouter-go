package openroutergo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleMessage(t *testing.T) {
	messages := []Msg{
		RoleUser.New("My name is user."),
		RoleSystem.New("My name is system."),
		RoleAssistant.New("My name is assistant."),
		RoleDeveloper.New("My name is developer."),
		RoleTool.New("My name is tool."),
	}

	for _, msg := range messages {
		b, err := json.Marshal(msg)
		assert.NoError(t, err)

		var m Msg
		assert.NoError(t, json.Unmarshal(b, &m))
		assert.Equal(t, msg, m)
	}
}
