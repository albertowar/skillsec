# Multi-Platform Provider Support Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Enable SkillSec to detect and audit skills for OpenAI (ChatGPT) and Anthropic (Claude) by introducing provider-specific profiles and danger registries.

**Architecture:** A new `internal/provider` package will house a `Provider` interface and concrete implementations for each platform. The `Parser` will use these profiles to auto-detect the provider based on "signature tools," and security checks will use the detected provider to fetch platform-specific risks.

**Tech Stack:** Go (Standard Library)

---

### Task 1: Define Provider Interface and API Types

**Files:**
- Create: `internal/provider/provider.go`
- Modify: `pkg/api/types.go`

- [ ] **Step 1: Define the Provider interface**

```go
package provider

type Provider interface {
	ID() string
	Name() string
	SignatureTools() []string
	DangerousTools() []string
}
```

- [ ] **Step 2: Add Provider field to SkillContext**

```go
// pkg/api/types.go

type SkillContext struct {
	Raw          string         `json:"raw"`
	Tools        []string       `json:"tools"`
	Provider     string         `json:"provider"` // Add this
	SystemPrompt string         `json:"systemPrompt"`
	Examples     []string       `json:"examples"`
	Metadata     *SkillMetadata `json:"metadata,omitempty"`
}
```

- [ ] **Step 3: Commit interface and type changes**

```bash
git add internal/provider/provider.go pkg/api/types.go
git commit -m "feat: define Provider interface and update SkillContext"
```

### Task 2: Implement Provider Profiles and Registry

**Files:**
- Create: `internal/provider/gemini.go`
- Create: `internal/provider/openai.go`
- Create: `internal/provider/anthropic.go`
- Create: `internal/provider/generic.go`
- Create: `internal/provider/registry.go`

- [ ] **Step 1: Implement Gemini Profile**

```go
package provider

type GeminiProvider struct{}

func (p *GeminiProvider) ID() string   { return "google" }
func (p *GeminiProvider) Name() string { return "Gemini" }
func (p *GeminiProvider) SignatureTools() []string {
	return []string{"run_shell_command", "write_file", "google_search"}
}
func (p *GeminiProvider) DangerousTools() []string {
	return []string{"run_shell_command", "write_file", "delete_file"}
}
```

- [ ] **Step 2: Implement OpenAI Profile**

```go
package provider

type OpenAIProvider struct{}

func (p *OpenAIProvider) ID() string   { return "openai" }
func (p *OpenAIProvider) Name() string { return "ChatGPT" }
func (p *OpenAIProvider) SignatureTools() []string {
	return []string{"code_interpreter", "dalle", "browser"}
}
func (p *OpenAIProvider) DangerousTools() []string {
	return []string{"code_interpreter"}
}
```

- [ ] **Step 3: Implement Anthropic Profile**

```go
package provider

type AnthropicProvider struct{}

func (p *AnthropicProvider) ID() string   { return "anthropic" }
func (p *AnthropicProvider) Name() string { return "Claude" }
func (p *AnthropicProvider) SignatureTools() []string {
	return []string{"computer", "bash", "text_editor"}
}
func (p *AnthropicProvider) DangerousTools() []string {
	return []string{"computer", "bash", "text_editor"}
}
```

- [ ] **Step 4: Implement Generic Profile**

```go
package provider

type GenericProvider struct{}

func (p *GenericProvider) ID() string               { return "generic" }
func (p *GenericProvider) Name() string             { return "Generic" }
func (p *GenericProvider) SignatureTools() []string { return nil }
func (p *GenericProvider) DangerousTools() []string { return nil }
```

- [ ] **Step 5: Create Provider Registry**

```go
package provider

import "strings"

var registry = []Provider{
	&GeminiProvider{},
	&OpenAIProvider{},
	&AnthropicProvider{},
}

func Get(id string) Provider {
	for _, p := range registry {
		if p.ID() == id {
			return p
		}
	}
	return &GenericProvider{}
}

func Detect(tools []string) string {
	for _, t := range tools {
		lowerT := strings.ToLower(t)
		for _, p := range registry {
			for _, sig := range p.SignatureTools() {
				if lowerT == sig {
					return p.ID()
				}
			}
		}
	}
	return "generic"
}
```

- [ ] **Step 6: Commit Registry**

```bash
git add internal/provider/*.go
git commit -m "feat: implement provider profiles and registry"
```

### Task 3: Update Parser with Detection Logic

**Files:**
- Modify: `internal/engine/parser.go`

- [ ] **Step 1: Update ParseSkill to include Detect call**

```go
// internal/engine/parser.go

import "github.com/albertowar/skillsec/internal/provider"

func ParseSkill(content string) api.SkillContext {
    // ... existing extraction logic ...
    
    return api.SkillContext{
		Raw:          content,
		Tools:        tools,
		Provider:     provider.Detect(tools), // Add this line
		SystemPrompt: getSection("Instructions"),
		Examples:     examples,
	}
}
```

- [ ] **Step 2: Commit Parser changes**

```bash
git add internal/engine/parser.go
git commit -m "feat: add provider detection to parser"
```

### Task 4: Refactor Dangerous Tools Check

**Files:**
- Modify: `internal/checks/dangerous_tools.go`
- Test: `internal/checks/dangerous_tools_test.go` (verify existing still works, add new cases)

- [ ] **Step 1: Refactor DangerousToolsCheck**

```go
// internal/checks/dangerous_tools.go

import "github.com/albertowar/skillsec/internal/provider"

func (c *DangerousToolsCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	p := provider.Get(skill.Provider)
	dangerousList := p.DangerousTools()
	
	dangerousMap := make(map[string]bool)
	for _, dt := range dangerousList {
		dangerousMap[strings.ToLower(dt)] = true
	}
    
    // ... logic to check skill.Tools against dangerousMap ...
}
```

- [ ] **Step 2: Add test cases for OpenAI and Anthropic tools**

```go
// internal/checks/dangerous_tools_test.go

func TestDangerousTools_OpenAI(t *testing.T) {
    skill := api.SkillContext{
        Provider: "openai",
        Tools: []string{"code_interpreter"},
    }
    // Assert Critical
}
```

- [ ] **Step 3: Run tests and commit**

```bash
go test ./internal/checks/...
git add internal/checks/dangerous_tools.go internal/checks/dangerous_tools_test.go
git commit -m "feat: make dangerous tools check provider-aware"
```

### Task 5: Update CLI Header

**Files:**
- Modify: `cmd/skillsec/main.go`

- [ ] **Step 1: Display provider in the report**

```go
// cmd/skillsec/main.go

func renderTable(report api.AuditReport, skillCtx api.SkillContext) {
    p := provider.Get(skillCtx.Provider)
    color.New(color.Bold).Printf("\nSkillSec Report (%s) - Score: %.1f/10\n\n", p.Name(), report.FinalScore)
    // ...
}
```

- [ ] **Step 2: Commit CLI changes**

```bash
git add cmd/skillsec/main.go
git commit -m "feat: show detected provider in CLI report"
```
