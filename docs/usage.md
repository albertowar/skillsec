# CLI Usage Guide

SkillAuditAI provides a command-line interface to audit AI skills against security best practices.

## Installation

```bash
npm install
npm run build
```

## Running an Audit

Audit a local `SKILL.md` file:

```bash
npx ts-node packages/cli/src/index.ts path/to/skill.md
```

## Flags & Options

| Flag | Description | Options |
|------|-------------|---------|
| `--format` | Output format | `table` (default), `json` |
| `--api-key` | API key for behavioral simulations | String |
| `--model` | LLM model name | e.g., `gemini-1.5-pro` |
| `--provider` | LLM provider | `google`, `openai`, `anthropic`, `custom` |
| `--base-url` | Custom API endpoint | URL string |

## Environment Variables

You can set these in a `.env` file in the root directory:

- `GOOGLE_API_KEY`
- `OPENAI_API_KEY`
- `ANTHROPIC_API_KEY`
