# Architecture Overview

SkillSec is designed as a modular auditing engine implemented in Go.

## The Pipeline

1. **Parse**: `internal/engine/parser.go` converts a Markdown skill into a `SkillContext`.
2. **Context**: Enrichment with Git metadata (author, maintenance) in `internal/engine/git.go`.
3. **Check**: The `Auditor` executes a registry of `Check` implementations concurrently.
4. **Behavioral**: `internal/behavioral/service.go` performs live LLM "Red Teaming" using `langchaingo`.
5. **Report**: Results are aggregated and formatted by the CLI in `cmd/skillsec/main.go`.

## Key Packages

- `pkg/api`: Public types for audit reports and results.
- `internal/engine`: Core logic for parsing, Git enrichment, and audit orchestration.
- `internal/checks`: Modular implementations of individual security checks.
- `internal/behavioral`: LLM abstraction layer using `langchaingo`.
