package engine

import (
	"testing"
)

func TestParseSkill(t *testing.T) {
	content := `## Tools
- run_shell_command
- ` + "`" + `write_file` + "`" + `

## Instructions
Execute tasks safely.

## Examples
User: hi
---
Assistant: hello`

	ctx := ParseSkill(content)
	if len(ctx.Tools) != 2 || ctx.Tools[0] != "run_shell_command" {
		t.Errorf("expected 2 tools, got %v", ctx.Tools)
	}
	if ctx.SystemPrompt != "Execute tasks safely." {
		t.Errorf("wrong system prompt: %s", ctx.SystemPrompt)
	}
	if len(ctx.Examples) != 2 {
		t.Errorf("expected 2 example parts, got %v", ctx.Examples)
	}
}

func TestParseSkillProviderDetection(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name: "Google (Gemini)",
			content: `## Tools
- run_shell_command`,
			expected: "google",
		},
		{
			name: "Anthropic (Claude)",
			content: `## Tools
- bash`,
			expected: "anthropic",
		},
		{
			name: "OpenAI (ChatGPT)",
			content: `## Tools
- code_interpreter`,
			expected: "openai",
		},
		{
			name: "Generic",
			content: `## Tools
- unknown_tool`,
			expected: "generic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := ParseSkill(tt.content)
			if ctx.Provider != tt.expected {
				t.Errorf("expected provider %s, got %s", tt.expected, ctx.Provider)
			}
		})
	}
}
