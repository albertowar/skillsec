# SkillAuditAI Go Migration Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Migrate the SkillAuditAI TypeScript codebase to an idiomatic, standalone Go CLI tool using `langchaingo` for behavioral analysis.

**Architecture:** Internal-first structure with a core engine (`internal/engine`), modular security checks (`internal/checks`), and an LLM-agnostic behavioral layer (`internal/behavioral`) using `langchaingo`. Public types are exposed in `pkg/api`.

**Tech Stack:** Go 1.21+, `langchaingo`, `cli-table`, `chalk`-like library for Go (e.g., `github.com/fatih/color`).

---

### Task 1: Project Initialization

**Files:**
- Create: `go.mod`
- Create: `pkg/api/types.go`

- [ ] **Step 1: Initialize Go module**

Run: `go mod init github.com/albertowar/skillauditai`
Expected: `go.mod` file created.

- [ ] **Step 2: Install core dependencies**

Run: `go get github.com/tmc/langchaingo github.com/fatih/color github.com/olekukonko/tablewriter`
Expected: `go.sum` updated.

- [ ] **Step 3: Define public types**

```go
package api

import "time"

type RiskLevel string

const (
	Critical RiskLevel = "Critical"
	High     RiskLevel = "High"
	Medium   RiskLevel = "Medium"
	Low      RiskLevel = "Low"
	Info     RiskLevel = "Info"
)

type CheckResult struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Score         float64   `json:"score"` // 0-10
	Level         RiskLevel `json:"level"`
	Justification string    `json:"justification"`
}

type AuditReport struct {
	SkillHash  string        `json:"skillHash"`
	FinalScore float64       `json:"finalScore"`
	Results    []CheckResult `json:"results"`
	Timestamp  time.Time     `json:"timestamp"`
}

type SkillContext struct {
	Raw          string   `json:"raw"`
	Tools        []string `json:"tools"`
	SystemPrompt string   `json:"systemPrompt"`
	Examples     []string `json:"examples"`
}
```

- [ ] **Step 4: Commit**

```bash
git add go.mod go.sum pkg/api/types.go
git commit -m "chore: initialize go project and public types"
```

---

### Task 2: Implement Markdown Parser

**Files:**
- Create: `internal/engine/parser.go`
- Create: `internal/engine/parser_test.go`

- [ ] **Step 1: Write failing test for parser**

```go
package engine

import (
	"testing"
	"github.com/albertowar/skillauditai/pkg/api"
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
```

- [ ] **Step 2: Run test to verify failure**

Run: `go test ./internal/engine/...`
Expected: FAIL (ParseSkill undefined)

- [ ] **Step 3: Implement Parser**

```go
package engine

import (
	"regexp"
	"strings"
	"github.com/albertowar/skillauditai/pkg/api"
)

func ParseSkill(content string) api.SkillContext {
	getSection := func(name string) string {
		re := regexp.MustCompile(`(?i)## ` + name + `\s+([\s\S]*?)(?=##|$)`)
		match := re.FindStringSubmatch(content)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		}
		return ""
	}

	toolsSection := getSection("Tools")
	var tools []string
	if toolsSection != "" {
		lines := strings.Split(toolsSection, "\n")
		re := regexp.MustCompile(`-\s*` + "`" + `?([\w_]+)` + "`" + `?`)
		for _, line := range lines {
			match := re.FindStringSubmatch(line)
			if len(match) > 1 {
				tools = append(tools, match[1])
			}
		}
	}

	examplesSection := getSection("Examples")
	var examples []string
	if examplesSection != "" {
		parts := strings.Split(examplesSection, "\n---\n")
		for _, p := range parts {
			if trimmed := strings.TrimSpace(p); trimmed != "" {
				examples = append(examples, trimmed)
			}
		}
	}

	return api.SkillContext{
		Raw:          content,
		Tools:        tools,
		SystemPrompt: getSection("Instructions"),
		Examples:     examples,
	}
}
```

- [ ] **Step 4: Verify tests pass**

Run: `go test ./internal/engine/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add internal/engine/parser.go internal/engine/parser_test.go
git commit -m "feat: implement markdown skill parser"
```

---

### Task 3: Define Check Interface and Registry

**Files:**
- Create: `internal/checks/check.go`

- [ ] **Step 1: Define Check interface**

```go
package checks

import (
	"context"
	"github.com/albertowar/skillauditai/pkg/api"
	"github.com/albertowar/skillauditai/internal/behavioral"
)

type Check interface {
	ID() string
	Name() string
	Weight() float64
	Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error)
}
```

- [ ] **Step 2: Add check registry**

```go
func AllChecks() []Check {
	return []Check{
		&DangerousToolsCheck{},
		// ... more will be added
	}
}
```

- [ ] **Step 3: Commit**

```bash
git add internal/checks/check.go
git commit -m "feat: define check interface and registry"
```

---

### Task 4: Implement Behavioral Service (langchaingo)

**Files:**
- Create: `internal/behavioral/service.go`

- [ ] **Step 1: Implement service wrapper**

```go
package behavioral

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

type Service struct {
	Model llms.Model
}

func NewService(provider, apiKey, modelName, baseURL string) (*Service, error) {
	var model llms.Model
	var err error

	switch provider {
	case "google":
		model, err = googleai.New(context.Background(), googleai.WithAPIKey(apiKey), googleai.WithDefaultModel(modelName))
	case "openai":
		model, err = openai.New(openai.WithToken(apiKey), openai.WithModel(modelName), openai.WithBaseURL(baseURL))
	default:
		return nil, nil // Not configured
	}

	if err != nil {
		return nil, err
	}
	return &Service{Model: model}, nil
}

func (s *Service) Test(ctx context.Context, systemPrompt, userMessage string) (string, error) {
	if s == nil || s.Model == nil {
		return "Behavioral testing not configured.", nil
	}
	
	resp, err := s.Model.Call(ctx, []llms.MessageContent{
		llms.TextPart(llms.ChatMessageTypeSystem, systemPrompt),
		llms.TextPart(llms.ChatMessageTypeHuman, userMessage),
	})
	if err != nil {
		return "", err
	}
	return resp, nil
}
```

- [ ] **Step 2: Commit**

```bash
git add internal/behavioral/service.go
git commit -m "feat: implement behavioral service with langchaingo"
```

---

### Task 5: Implement Static Checks (Example: Dangerous Tools)

**Files:**
- Create: `internal/checks/dangerous_tools.go`

- [ ] **Step 1: Implement DangerousToolsCheck**

```go
package checks

import (
	"context"
	"fmt"
	"strings"
	"github.com/albertowar/skillauditai/pkg/api"
	"github.com/albertowar/skillauditai/internal/behavioral"
)

type DangerousToolsCheck struct{}

func (c *DangerousToolsCheck) ID() string      { return "dangerous-tools" }
func (c *DangerousToolsCheck) Name() string    { return "Dangerous Tools Audit" }
func (c *DangerousToolsCheck) Weight() float64 { return 1.0 }

func (c *DangerousToolsCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	dangerous := []string{"run_shell_command", "write_file", "delete_file"}
	var found []string
	
	for _, t := range skill.Tools {
		for _, d := range dangerous {
			if t == d {
				found = append(found, t)
			}
		}
	}

	if len(found) > 0 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         0,
			Level:         api.Critical,
			Justification: fmt.Sprintf("Skill requests highly dangerous tools: %s", strings.Join(found, ", ")),
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         10,
		Level:         api.Low,
		Justification: "No dangerous tools detected.",
	}, nil
}
```

- [ ] **Step 2: Update registry in `internal/checks/check.go`**

```go
func AllChecks() []Check {
	return []Check{
		&DangerousToolsCheck{},
	}
}
```

- [ ] **Step 3: Commit**

```bash
git add internal/checks/dangerous_tools.go internal/checks/check.go
git commit -m "feat: add dangerous tools check"
```

---

### Task 6: Implement Auditor Engine

**Files:**
- Create: `internal/engine/auditor.go`

- [ ] **Step 1: Implement Auditor**

```go
package engine

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
	"github.com/albertowar/skillauditai/pkg/api"
	"github.com/albertowar/skillauditai/internal/checks"
	"github.com/albertowar/skillauditai/internal/behavioral"
)

type Auditor struct {
	Checks     []checks.Check
	Behavioral *behavioral.Service
}

func NewAuditor(b *behavioral.Service) *Auditor {
	return &Auditor{
		Checks:     checks.AllChecks(),
		Behavioral: b,
	}
}

func (a *Auditor) Audit(ctx context.Context, skill api.SkillContext) (api.AuditReport, error) {
	var wg sync.WaitGroup
	results := make([]api.CheckResult, len(a.Checks))
	
	for i, c := range a.Checks {
		wg.Add(1)
		go func(idx int, check checks.Check) {
			defer wg.Done()
			res, err := check.Run(ctx, skill, a.Behavioral)
			if err != nil {
				results[idx] = api.CheckResult{
					ID:            check.ID(),
					Name:          check.Name(),
					Score:         0,
					Level:         api.Critical,
					Justification: fmt.Sprintf("Check failed: %v", err),
				}
			} else {
				results[idx] = res
			}
		}(i, c)
	}

	wg.Wait()

	var weightedScore float64
	var totalWeight float64
	for i, r := range results {
		w := a.Checks[i].Weight()
		weightedScore += r.Score * w
		totalWeight += w
	}

	finalScore := 10.0
	if totalWeight > 0 {
		finalScore = weightedScore / totalWeight
	}

	hash := sha256.Sum256([]byte(skill.Raw))

	return api.AuditReport{
		SkillHash:  fmt.Sprintf("%x", hash),
		FinalScore: finalScore,
		Results:    results,
		Timestamp:  time.Now(),
	}, nil
}
```

- [ ] **Step 2: Commit**

```bash
git add internal/engine/auditor.go
git commit -m "feat: implement auditor orchestration engine"
```

---

### Task 7: Implement CLI (main.go)

**Files:**
- Create: `cmd/skillaudit/main.go`

- [ ] **Step 1: Implement main function with flags and table output**

```go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"github.com/albertowar/skillauditai/internal/engine"
	"github.com/albertowar/skillauditai/internal/behavioral"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func main() {
	format := flag.String("format", "table", "Output format (table|json)")
	apiKey := flag.String("api-key", os.Getenv("SKILLAUDIT_API_KEY"), "LLM API Key")
	model := flag.String("model", "gemini-1.5-pro", "LLM Model Name")
	provider := flag.String("provider", "google", "LLM Provider (google|openai)")
	baseURL := flag.String("base-url", "", "Custom LLM base URL")
	flag.Parse()

	filePath := flag.Arg(0)
	if filePath == "" {
		fmt.Println("Usage: skillaudit <skill.md> [flags]")
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	skillCtx := engine.ParseSkill(string(content))
	bService, _ := behavioral.NewService(*provider, *apiKey, *model, *baseURL)
	auditor := engine.NewAuditor(bService)
	
	report, err := auditor.Audit(context.Background(), skillCtx)
	if err != nil {
		fmt.Printf("Audit failed: %v\n", err)
		os.Exit(1)
	}

	if *format == "json" {
		data, _ := json.MarshalIndent(report, "", "  ")
		fmt.Println(string(data))
	} else {
		renderTable(report)
	}
}

func renderTable(report interface{}) {
	// ... implementation using tablewriter and color
}
```

- [ ] **Step 2: Commit**

```bash
git add cmd/skillaudit/main.go
git commit -m "feat: implement CLI interface and output formatting"
```
