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

func TestParseSkillComplex(t *testing.T) {
	content := `## Tools
- read_file

## Instructions
Keep it simple.

## Examples
User: test
---
Assistant: ok
---
User: more
---
Assistant: done

## Metadata
Some extra data here.`

	ctx := ParseSkill(content)
	if len(ctx.Tools) != 1 || ctx.Tools[0] != "read_file" {
		t.Errorf("wrong tools: %v", ctx.Tools)
	}
	if ctx.SystemPrompt != "Keep it simple." {
		t.Errorf("wrong instructions: %s", ctx.SystemPrompt)
	}
	if len(ctx.Examples) != 4 {
		t.Errorf("expected 4 example parts, got %v", ctx.Examples)
	}
}
