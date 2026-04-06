# Design Spec: Multi-Platform Provider Support (Gemini, OpenAI, Anthropic)

## Overview
Currently, SkillSec is heavily biased toward Gemini-specific tools (`run_shell_command`, `write_file`). This spec outlines a strategy to introduce "Provider Profiles" that allow the tool to detect and audit skills for OpenAI (ChatGPT) and Anthropic (Claude) using the same `SKILL.md` format.

## Goals
- **Auto-Detection**: Identify the AI provider based on signature tools or instructions.
- **Platform-Specific Auditing**: Customize security checks (like `DangerousToolsCheck`) based on the detected provider.
- **Extensibility**: Make it easy to add new providers (e.g., DeepSeek, MCP) in the future.
- **Generic Fallback**: Support platform-agnostic auditing for unknown or custom providers.

## Architecture

### 1. The Provider Registry (`internal/provider/`)
A new package will manage provider definitions. Each provider will implement a standard interface:

```go
type Provider interface {
    ID() string             // e.g., "google", "openai", "anthropic"
    Name() string           // e.g., "Gemini", "ChatGPT", "Claude"
    SignatureTools() []string // Tools unique to this provider used for detection
    DangerousTools() []string // Tools that trigger a Critical security alert
}
```

### 2. Detection Logic
The `Parser` will use a `Registry` to match extracted tools against `SignatureTools()`:
- `run_shell_command` -> **Gemini**
- `code_interpreter`, `dalle`, `browser` -> **OpenAI**
- `computer`, `bash`, `text_editor` -> **Anthropic**

If no signatures match, the provider defaults to `"generic"`.

### 3. Data Model Changes (`pkg/api/types.go`)
Update `SkillContext` to include the detected provider:

```go
type SkillContext struct {
    Raw          string         `json:"raw"`
    Tools        []string       `json:"tools"`
    Provider     string         `json:"provider"` // New field
    SystemPrompt string         `json:"systemPrompt"`
    // ... rest of fields
}
```

### 4. Check Refactoring (`internal/checks/`)
The `DangerousToolsCheck` will be updated to be provider-aware:

```go
func (c *DangerousToolsCheck) Run(ctx context.Context, skill api.SkillContext, ...) {
    p := provider.Get(skill.Provider)
    dangerous := p.DangerousTools()
    // Compare skill.Tools against dangerous
}
```

## User Experience
The `skillsec` CLI will display the detected provider in the report header:
`SkillSec Report (OpenAI) - Score: 8.5/10`

## Testing Plan
- **Detection Tests**: Provide snippets with `code_interpreter` and verify `SkillContext.Provider == "openai"`.
- **Audit Consistency**: Verify that `run_shell_command` is Critical for Gemini but not for OpenAI (unless explicitly added to OpenAI's danger list).
- **Generic Handling**: Ensure secret scanning and other behavioral checks still run even if the provider is `"generic"`.
