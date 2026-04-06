# CLI Usage Guide

SkillAuditAI provides a command-line interface to audit AI skills against security best practices.

## Installation

Building from source requires Go 1.21+:

```bash
go build -o skillaudit ./cmd/skillaudit
```

## Running an Audit

Audit a local `SKILL.md` file:

```bash
./skillaudit path/to/skill.md
```

## Flags & Options

| Flag | Description | Options |
|------|-------------|---------|
| `-format` | Output format | `table` (default), `json` |
| `-api-key` | API key for behavioral simulations | String |
| `-model` | LLM model name | e.g., `gemini-1.5-pro` |
| `-provider` | LLM provider | `google`, `openai` |
| `-base-url` | Custom API endpoint | URL string |

## Environment Variables

You can set these in your environment or a `.env` file:

- `SKILLAUDIT_API_KEY` (Can be used instead of `-api-key`)
