# Design Spec: Documentation (README.md & AGENTS.md)

**Date:** 2026-04-05  
**Status:** Draft  
**Target:** Project Root Documentation  

---

## 1. Overview
This spec defines the structure and content for the primary repository documentation: `README.md` (targeted at contributors) and `AGENTS.md` (targeted at AI coding agents).

## 2. README.md (Contributor-Focused)

### 2.1 Goals
- Provide a clear entry point for developers.
- Explain the monorepo structure and the relationship between `@skillauditai/core` and `@skillauditai/cli`.
- Document the extension mechanism for adding new security checks.

### 2.2 Content Structure
1.  **Header:** Project name, brief description (Security Auditor for AI Skills), and link to [agentskills.io](https://agentskills.io).
2.  **Architecture:**
    - High-level diagram (text-based) of the pipeline: Parse → Context → Check → Score → Report.
    - Package breakdown: `packages/core` (stateless engine) vs. `packages/cli` (I/O & rendering).
3.  **Getting Started:**
    - Prerequisites (Node.js, npm).
    - Installation and build instructions (`npm install && npm run build`).
4.  **Extending SkillAuditAI:**
    - Walkthrough of the `BaseCheck` interface.
    - How to define check metadata (ID, Description, Weight).
    - Registering a check in the core engine.
5.  **Development:**
    - Testing with Vitest.
    - Linting and formatting rules.
    - Contribution guidelines (briefly).

## 3. AGENTS.md (AI-Focused)

### 3.1 Goals
- Optimize the developer experience for AI agents working on the codebase.
- Enforce architectural constraints and scoring consistency.
- Provide a quick reference for core types and symbols.

### 3.2 Content Structure
1.  **Core Mandates:**
    - **Stateless Core:** `packages/core` must remain edge-ready (no filesystem/network I/O; use dependency injection or context).
    - **Type Safety:** Strict TypeScript usage; no `any` or suppressed warnings.
2.  **Key Symbols:**
    - `SkillAST`: The structured representation of an audited skill.
    - `BaseCheck`: The abstract class for all audit logic.
    - `AuditReport`: The final output structure.
3.  **Scoring Logic:**
    - Definition of weights: `Critical (1.0)`, `High (0.8)`, `Medium (0.5)`, `Low (0.3)`.
    - Rules for score aggregation.
4.  **Workflow Requirements:**
    - Every bug must have a reproduction test case.
    - Every new check must have a corresponding `.spec.ts` file.
    - Every implementation must align with the `agentskills.io` specification.
5.  **Parsing Rules:**
    - Specific regex or parser instructions for the Markdown-to-AST conversion.

---

## 4. Technical Constraints
- Use standard Markdown (GitHub-flavored).
- Keep `AGENTS.md` concise and machine-parsable.
- Ensure all links (e.g., to packages) are relative and correct.
