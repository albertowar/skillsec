# Design Spec: Developer & User Documentation Structure

**Date:** 2026-04-05
**Status:** Approved
**Target:** Implementation of comprehensive developer and user documentation for SkillSec.

---

## 1. Overview
This design covers the creation of a modular documentation suite for SkillSec, providing both user-facing guides for auditing AI skills and developer-facing documentation for contributing to the core engine.

## 2. Documentation Structure

All documentation will be located in the `docs/` directory, following a modular approach to separate usage, interpretation, and architecture.

### 2.1 User Documentation (`docs/usage.md`)
- **CLI Commands:** Instructions for running the auditor via `npx` or locally.
- **Flags & Options:** Detailed table of all supported flags:
    - `--format`: `table` (default), `json`.
    - `--api-key`: API key for behavioral simulations.
    - `--model`: Model name (e.g., `gemini-1.5-pro`).
    - `--provider`: `google`, `openai`, `anthropic`, `custom`.
    - `--base-url`: Custom endpoint for local or proxy models.
- **Environment Setup:** Recommended `.env` configuration for persistent API keys.

### 2.2 Interpretation Guide (`docs/interpreting-results.md`)
- **Scoring Logic:** Explanation of how the final score (0-10) is calculated using weighted averages.
- **Risk Level Definitions:**
    - `Critical`: Direct code execution (RCE) or confirmed data exfiltration.
    - `High`: Dangerous tools present without safety constraints or human-in-the-loop (HITL) instructions.
    - `Medium`: Significant vulnerabilities that require manual review.
    - `Low`: Minimal risks; follow best practices for delimiters and safety phrases.
    - `Info`: Metadata-only findings (maintenance, author verification).
- **Check Catalog:** A glossary of all 10 active checks and their specific "Fail" criteria.

### 2.3 Architecture Overview (`docs/architecture.md`)
- **The Pipeline:** Diagram and description of the `Parse -> Context -> Check -> Score -> Report` flow.
- **Package Relationships:** Separation of concerns between `@skillsec/core` (stateless engine) and `@skillsec/cli` (I/O wrapper).
- **Behavioral Service:** How the optional LLM-powered "Red Teaming" layer is injected into the audit process.

### 2.4 Extension Guide (`docs/extending.md`)
- **Implementing `BaseCheck`:** Step-by-step instructions and a boilerplate example for a new check.
- **Using `SkillContext`:** Documentation for available properties (`tools`, `systemPrompt`, `examples`, `metadata`).
- **Testing Requirements:** Guidelines for using `vitest` to verify check logic with mock contexts.
- **Registration:** How to add the new check to the `Auditor` constructor.

---

## 3. Implementation Plan

1.  Create the `docs/` directory (if not exists) and populate with the four Markdown files.
2.  Update `README.md` at the project root to link to the new documentation.
3.  Ensure all code examples in the documentation are accurate and reflect the current state of the repository.
4.  Commit the documentation to the `main` branch.
