# Spec: SkillAuditAI Go Migration

**Date:** 2026-04-06
**Status:** Draft
**Topic:** Migrating the TypeScript codebase to an idiomatic Go CLI tool using `langchaingo`.

## 1. Purpose
The goal is to provide a standalone, high-performance executable for auditing AI "Skills" (defined in Markdown). The migration will improve distribution (single binary), performance (concurrency), and leverage the `langchaingo` ecosystem for LLM-based behavioral testing.

## 2. Architecture

### 2.1 Project Structure (Internal-First)
Following Go best practices (Approach 2 from brainstorming):

```text
skillauditai/
├── cmd/
│   └── skillaudit/
│       └── main.go          # CLI Entry point, flag parsing, formatting
├── internal/
│   ├── engine/
│   │   ├── auditor.go       # Orchestration & scoring logic
│   │   └── parser.go        # Markdown parsing logic
│   ├── checks/
│   │   ├── check.go         # Base interface and registry
│   │   ├── dangerous_tools.go
│   │   ├── exfiltration.go
│   │   └── ...              # Other security checks
│   └── behavioral/
│       └── service.go       # langchaingo wrapper
├── pkg/
│   └── api/
│       └── types.go         # Public report and result types
├── go.mod
└── go.sum
```

### 2.2 Key Components

#### Auditor (`internal/engine`)
- Manages the lifecycle of an audit.
- Uses `Parser` to transform Markdown into a `SkillContext`.
- Executes `Checks` concurrently using goroutines.
- Calculates a weighted final score (0-10).

#### Parser (`internal/engine`)
- Extracts sections (Tools, Instructions, Examples) from Markdown.
- Identifies tool names from bulleted lists.
- Sanitizes system prompts and examples for analysis.

#### Checks (`internal/checks`)
- **Interface**:
  ```go
  type Check interface {
      ID() string
      Name() string
      Weight() float64
      Run(ctx context.Context, skill pkg.SkillContext, b *behavioral.Service) (pkg.CheckResult, error)
  }
  ```
- **Static Checks**: Dangerous tools, secret scanning, dependency audit, verified author, maintenance.
- **Behavioral Checks**: Prompt injection, exfiltration vectors, indirect injection, least privilege, tool chaining.

#### Behavioral Service (`internal/behavioral`)
- Utilizes `langchaingo` to interface with LLMs.
- Supports Google (Vertex/AI Studio), OpenAI, and Anthropic.
- Provides a standard interface for checks to run "probes" against the skill's logic.

## 3. Data Flow
1. **Input**: User runs `skillaudit skill.md --api-key xxx`.
2. **Setup**: `main.go` parses flags and initializes the `Auditor` with a `BehavioralService`.
3. **Parse**: `Parser` reads `skill.md` and populates `SkillContext`.
4. **Execute**: `Auditor` spawns goroutines for each `Check`.
5. **Score**: `Auditor` waits for all results, applies weights, and builds the `AuditReport`.
6. **Output**: `main.go` prints a `cli-table` or JSON.

## 4. Error Handling
- Checks that fail (e.g., LLM timeout) return a result with `RiskLevel: Critical` and a justification explaining the failure.
- CLI provides clear diagnostics for file access issues or missing environment variables/flags.

## 5. Success Criteria
- [ ] Successfully parses existing `vulnerable-skill.md`.
- [ ] Reproduces scores similar to the TypeScript implementation for static checks.
- [ ] Successfully connects to at least one LLM provider via `langchaingo`.
- [ ] Produces a single, portable binary for Linux/macOS/Windows.
