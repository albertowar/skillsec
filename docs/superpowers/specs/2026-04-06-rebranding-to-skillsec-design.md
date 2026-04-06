# Design Spec: Rebranding to SkillSec

This document outlines the full technical and visual rebranding of the project from **SkillSec** to **SkillSec**.

## 1. Core Identity Changes

| Entity | Old Value | New Value |
| :--- | :--- | :--- |
| Project Name | SkillSec | SkillSec |
| Domain | N/A | skillsec.dev |
| Go Module | `github.com/albertowar/skillsec` | `github.com/albertowar/skillsec` |
| Binary Name | `skillsec` | `skillsec` |
| Repository URL | `https://github.com/albertowar/skillsec` | `https://github.com/albertowar/skillsec` |
| Primary Env Var | `SKILLSEC_API_KEY` | `SKILLSEC_API_KEY` |

## 2. Technical Implementation Plan

### 2.1 Go Module & Imports
- Update `go.mod`: Change `module github.com/albertowar/skillsec` to `module github.com/albertowar/skillsec`.
- Mass update all `.go` files: Replace all internal import paths from `github.com/albertowar/skillsec/...` to `github.com/albertowar/skillsec/...`.

### 2.2 Filesystem Changes
- Rename `cmd/skillsec` directory to `cmd/skillsec`.
- Update `main.go` within the renamed directory to reflect the new binary name in usage strings and version output.

### 2.3 Documentation & UI
- Perform a case-sensitive find-and-replace across all `.md`, `.yml`, and `.go` files:
    - `SkillSec` -> `SkillSec`
    - `skillsec` -> `skillsec`
    - `skillsec` -> `skillsec`
    - `SKILLSEC` -> `SKILLSEC`
- Update `mkdocs.yml`:
    - `site_name: SkillSec`
    - `repo_url: https://github.com/albertowar/skillsec`
    - `site_url: https://skillsec.dev` (if applicable)
- Update `README.md` and `docs/*.md` to reflect the new brand and installation/usage commands.

### 2.4 CI/CD & GitHub Actions
- Update `.github/workflows/release.yml`:
    - Change binary output names from `skillsec-*` to `skillsec-*`.
    - Update repository-specific environment variables if any.

## 3. Verification Strategy
- **Build Verification**: Run `go build -o skillsec ./cmd/skillsec` and ensure the binary functions correctly.
- **Import Verification**: Run `go test ./...` to ensure all internal imports are correctly resolved.
- **Documentation Verification**: Preview the `mkdocs` site (if possible) to ensure no "SkillSec" remains in titles or links.
- **Usage Verification**: Run `./skillsec --help` to verify the output displays the correct project name and flags.
