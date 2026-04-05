# Architecture Overview

SkillAuditAI is designed as a modular auditing engine.

## The Pipeline

1. **Parse**: `parser.ts` converts a Markdown skill into a `SkillContext`.
2. **Context**: Enrichment with Git metadata (author, maintenance).
3. **Check**: The `Auditor` executes a registry of `BaseCheck` implementations.
4. **Behavioral**: (Optional) `BehavioralService` performs live LLM "Red Teaming".
5. **Report**: Results are aggregated and formatted by the CLI.

## Packages

- `@skillauditai/core`: The stateless auditing engine and check logic.
- `@skillauditai/cli`: The terminal wrapper for local file I/O.
