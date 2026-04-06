# Documentation Generation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Create a comprehensive suite of developer and user documentation for SkillSec.

**Architecture:** A modular documentation structure in `docs/` containing usage guides, scoring interpretation, architecture overviews, and extension guides.

**Tech Stack:** Markdown.

---

### Task 1: Initialize Docs Directory and Usage Guide

**Files:**
- Create: `docs/usage.md`

- [ ] **Step 1: Create `docs/usage.md` with CLI instructions**

```markdown
# CLI Usage Guide

SkillSec provides a command-line interface to audit AI skills against security best practices.

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
```

- [ ] **Step 2: Commit**

```bash
git add docs/usage.md
git commit -m "docs: add CLI usage guide"
```

---

### Task 2: Create Interpretation Guide

**Files:**
- Create: `docs/interpreting-results.md`

- [ ] **Step 1: Create `docs/interpreting-results.md` with scoring and risk details**

```markdown
# Interpreting Audit Results

SkillSec provides a Safety Score from 0 to 10. A higher score indicates a safer skill.

## Scoring Logic

The final score is a weighted average of all active checks. Weights are assigned based on the severity of the potential vulnerability:

- **Critical (1.0)**: Direct threats to system integrity.
- **High (0.8)**: Serious misconfigurations.
- **Medium (0.5)**: Significant risks requiring review.
- **Low (0.3)**: Minor best-practice violations.

## Risk Level Definitions

- **🔴 Critical**: Direct Remote Code Execution (RCE) or confirmed data exfiltration path.
- **🟠 High**: Usage of dangerous tools without any safety delimiters or human-in-the-loop (HITL) constraints.
- **🟡 Medium**: Potentially dangerous patterns that should be manually audited.
- **🟢 Low**: Minor issues like missing delimiters but having safety phrases.
- **⚪ Info**: Lack of metadata (author, maintenance date).

## Audit Catalog

| Check ID | Description |
|----------|-------------|
| `dangerous-tools` | Detects high-risk tools like `run_shell_command`. |
| `tool-chaining` | Detects "Read -> Move" patterns (e.g., File Read to API Post). |
| `prompt-injection` | Checks for structural delimiters and safety phrases. |
| `exfiltration` | Scans for network sinks and tests robustness via LLM. |
| `indirect-injection` | Audits trust boundaries for external content readers. |
| `secret-scanning` | Searches for hardcoded API keys or secrets. |
```

- [ ] **Step 2: Commit**

```bash
git add docs/interpreting-results.md
git commit -m "docs: add interpretation guide"
```

---

### Task 3: Create Architecture and Extension Guides

**Files:**
- Create: `docs/architecture.md`
- Create: `docs/extending.md`

- [ ] **Step 1: Create `docs/architecture.md`**

```markdown
# Architecture Overview

SkillSec is designed as a modular auditing engine.

## The Pipeline

1. **Parse**: `parser.ts` converts a Markdown skill into a `SkillContext`.
2. **Context**: Enrichment with Git metadata (author, maintenance).
3. **Check**: The `Auditor` executes a registry of `BaseCheck` implementations.
4. **Behavioral**: (Optional) `BehavioralService` performs live LLM "Red Teaming".
5. **Report**: Results are aggregated and formatted by the CLI.

## Packages

- `@skillsec/core`: The stateless auditing engine and check logic.
- `@skillsec/cli`: The terminal wrapper for local file I/O.
```

- [ ] **Step 2: Create `docs/extending.md`**

```markdown
# Extending SkillSec

## Adding a New Check

1. Create a new file in `packages/core/src/checks/`.
2. Implement the `BaseCheck` interface:

```typescript
import { BaseCheck, SkillContext, CheckResult } from '../types';

export class MyNewCheck implements BaseCheck {
  id = 'my-check';
  name = 'My Security Audit';
  weight = 0.5;

  async run(context: SkillContext): Promise<CheckResult> {
    // Audit logic here
  }
}
```

3. Register the check in `packages/core/src/auditor.ts`.
4. Add tests in `packages/core/tests/checks/`.
```

- [ ] **Step 3: Commit**

```bash
git add docs/architecture.md docs/extending.md
git commit -m "docs: add architecture and extension guides"
```

---

### Task 4: Update Root README

**Files:**
- Modify: `README.md`

- [ ] **Step 1: Link documentation in root README**

```markdown
# SkillSec

... existing content ...

## 📖 Documentation

- [CLI Usage Guide](docs/usage.md)
- [Interpreting Results](docs/interpreting-results.md)
- [Architecture Overview](docs/architecture.md)
- [Extending the Auditor](docs/extending.md)

... rest of content ...
```

- [ ] **Step 2: Commit**

```bash
git add README.md
git commit -m "docs: link modular documentation in README"
```
