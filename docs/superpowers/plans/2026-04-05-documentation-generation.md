# Documentation Generation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Generate a contributor-focused `README.md` and an AI-focused `AGENTS.md` in the project root to align with the SkillAuditAI design.

**Architecture:** Create two primary documentation files in the repository root. `README.md` targets human contributors with architecture and extension guides, while `AGENTS.md` provides architectural mandates and type references for AI agents.

**Tech Stack:** Markdown (GitHub-flavored).

---

### Task 1: Generate README.md

**Files:**
- Create: `README.md`

- [ ] **Step 1: Write README.md content**

```markdown
# SkillAuditAI

> **Security Auditor for AI Skills**  
> Evaluates AI skills against the [agentskills.io](https://agentskills.io) specification to provide a unified "Safety Score" (0-10).

## 🏗 Architecture

SkillAuditAI is built as a TypeScript monorepo to ensure the auditing engine remains portable and "edge-ready."

- **`@skillauditai/core`**: The stateless engine. Handles parsing, check orchestration, and scoring.
- **`@skillauditai/cli`**: The user-facing tool for file I/O and terminal rendering.

### The Audit Pipeline
1. **Parse**: Converts `SKILL.md` into a structured `SkillAST`.
2. **Context**: (Optional) Enriches the audit with external metadata.
3. **Check**: Executes a registry of `BaseCheck` implementations.
4. **Score**: Aggregates check results into a weighted scorecard.
5. **Report**: Generates output in Table, JSON, or SARIF formats.

## 🛠 Getting Started

### Prerequisites
- Node.js 18+
- npm

### Installation
```bash
npm install
npm run build
```

## 🧩 Extending SkillAuditAI

To add a new security check, implement the `BaseCheck` interface in `@skillauditai/core`:

1. Define your check in `packages/core/src/checks/`.
2. Assign a weight: `Critical (1.0)`, `High (0.8)`, `Medium (0.5)`, or `Low (0.3)`.
3. Register the check in the `Auditor` engine.

## 🧪 Development

- **Test**: `npm test` (Uses Vitest)
- **Build**: `npm run build`
- **Lint**: `npm run lint`
```

- [ ] **Step 2: Verify file creation**

Run: `ls -l README.md`
Expected: File exists with the content above.

- [ ] **Step 3: Commit**

```bash
git add README.md
git commit -m "docs: add contributor-focused README.md"
```

### Task 2: Generate AGENTS.md

**Files:**
- Create: `AGENTS.md`

- [ ] **Step 1: Write AGENTS.md content**

```markdown
# Agent Instructions: SkillAuditAI

This file provides architectural mandates and type references for AI agents contributing to this repository.

## ⚖️ Core Mandates

1. **Stateless Core**: Everything in `packages/core` must be stateless and "edge-ready." No direct filesystem or network I/O. Use dependency injection or the `Context` object.
2. **Type Safety**: No `any` types. No @ts-ignore. Use Zod for runtime validation where boundaries are crossed.
3. **Spec Alignment**: All parsing and auditing logic must strictly follow the [agentskills.io](https://agentskills.io) specification.

## 🔑 Key Symbols & Types

- `SkillAST`: The unified AST generated from `SKILL.md`.
- `BaseCheck`: Abstract class for all audit checks. Must implement `run(ast: SkillAST): Promise<CheckResult>`.
- `AuditReport`: The final output structure containing the scorecard and findings.

## 📊 Scoring Weights

Consistency in scoring is critical. Use these weights for all `BaseCheck` implementations:
- **Critical (1.0)**: Direct security vulnerabilities (e.g., hardcoded secrets, dangerous shell tools).
- **High (0.8)**: Potential for misuse or prompt injection risks.
- **Medium (0.5)**: Reputation, metadata, or maintenance issues.
- **Low (0.3)**: Stylistic or minor best-practice deviations.

## 🔄 Workflow Requirements

- **TDD**: Every bug fix requires a reproduction test case. Every new check requires a `.spec.ts` file.
- **Parser Logic**: The Markdown parser must target sections: `Tools`, `Instructions`, and `Examples`.
- **Validation**: Run `npm test` before proposing any implementation.
```

- [ ] **Step 2: Verify file creation**

Run: `ls -l AGENTS.md`
Expected: File exists with the content above.

- [ ] **Step 3: Commit**

```bash
git add AGENTS.md
git commit -m "docs: add AI-focused AGENTS.md"
```

---

### Plan Self-Review
1. **Spec Coverage:** Covers all sections defined in the documentation design spec.
2. **Placeholder Scan:** None. Actual Markdown content provided.
3. **Type Consistency:** Method signatures and weights match the core design spec.
