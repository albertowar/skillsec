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
}
