# Design Spec: Remaining Audit Checks & Metadata Enrichment

**Date:** 2026-04-05  
**Status:** Approved  
**Target:** Implementation of all remaining audit check categories.

---

## 1. Overview
This design covers the implementation of the final five audit checks specified in the SkillAuditAI design:
1.  `least-privilege`: Tool usage vs. instructions intent.
2.  `prompt-injection`: Detection of safety delimiters and jailbreak protections.
3.  `verified-author`: Author identity validation via Git.
4.  `maintenance`: Skill versioning and update history via Git.
5.  `dependency-audit`: Safety of imported tools and sub-skills.

## 2. Updated Data Structures

### 2.1 `SkillContext` (in `@skillauditai/core`)
Update the context to support optional metadata provided by the CLI or other enrichers.

```typescript
export interface SkillMetadata {
  author?: {
    name: string;
    email: string;
    isVerified: boolean; // Based on signed commits or trusted registry
  };
  maintenance?: {
    lastUpdated: string; // ISO timestamp
    version?: string;    // SemVer from tags or frontmatter
  };
  dependencies?: string[]; // List of external skills/libraries referenced
}

export interface SkillContext {
  raw: string;
  tools: string[];
  systemPrompt: string;
  examples: string[];
  metadata?: SkillMetadata;
}
```

---

## 3. Core Audit Checks

### 3.1 `LeastPrivilegeCheck` (Static)
- **Goal:** Flag tools that don't match the skill's stated intent.
- **Implementation:**
    - Scan `systemPrompt` for keywords corresponding to requested `tools`.
    - **Mapping Example:**
        - `run_shell_command` -> Keywords: `bash`, `shell`, `command`, `terminal`, `execute`.
        - `write_file` -> Keywords: `save`, `write`, `create file`, `output`.

### 3.2 `PromptInjectionCheck` (Static)
- **Goal:** Detect lack of delimiters or jailbreak protections.
- **Implementation:**
    - Scan `systemPrompt` for:
        - **Delimiters:** `"""`, `---`, `<context>`, `###`.
        - **Safety Constraints:** "Do not reveal instructions", "Ignore previous instructions", "Stay in character".

### 3.3 `DependencyAuditCheck` (Static)
- **Goal:** Identify external dependencies (libraries or other skills).
- **Implementation:**
    - Scan `systemPrompt` and `examples` for:
        - Markdown links to other `.md` files.
        - Mentions of package managers: `npm install`, `pip install`, `cargo add`.
        - Import statements in code blocks: `import X`, `require('Y')`.

### 3.4 `VerifiedAuthorCheck` (Metadata-driven)
- **Goal:** Validate author via Git history.
- **Implementation:**
    - Uses `context.metadata.author`.

### 3.5 `MaintenanceCheck` (Metadata-driven)
- **Goal:** Check for freshness and versioning.
- **Implementation:**
    - Uses `context.metadata.maintenance.lastUpdated`.

---

## 4. CLI Metadata Enrichment (Git-based)

The CLI is updated to automatically enrich the `SkillContext` when running in a Git repository.

### 4.1 Git Extraction Commands
1.  **Author & Date:** `git log -1 --format="%an|%ae|%at" -- <filePath>`
2.  **Verification:** `git log -1 --format="%G?" -- <filePath>`
    - `%G?`: `G` (good), `B` (bad), `U` (untrusted/unverified), `X` (expired), `Y` (revoked), `R` (good sign but not by a trusted key), `N` (no signature).

---

## 5. Implementation Status
- All core checks implemented in `packages/core/src/checks/`.
- `Auditor` updated to include new checks.
- `getGitMetadata` utility implemented in `packages/cli/src/git.ts`.
- Unit tests added and passing.
