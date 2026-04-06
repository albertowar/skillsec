# Spec: Enhanced Release Process and Documentation Site

**Date:** 2026-04-06
**Status:** Approved
**Topic:** Improving the GitHub Release Action and setting up MkDocs for documentation deployment.

## 1. Purpose
Provide a professional, secure, and well-documented release cycle for SkillSec. This includes self-aware versioning, integrity verification (checksums), and an automated searchable documentation site.

## 2. Architecture

### 2.1 CLI Versioning
- **Mechanism**: Use Go's `ldflags` to inject a version string at build time.
- **Variable**: `main.Version` in `cmd/skillsec/main.go`.
- **User Interface**: `skillsec --version` flag.

### 2.2 GitHub Release Workflow (`.github/workflows/release.yml`)
The workflow will be updated to include the following stages:
1.  **Run Checks**: Existing `go test` and `go vet` logic.
2.  **Build and Release**:
    - Build for multiple OS/Architectures.
    - Inject `Version` variable.
    - Generate `sha256sum` for all artifacts.
    - Create GitHub Release and upload binaries + `checksums.txt`.
3.  **Deploy Docs**:
    - Triggered only after a successful release.
    - Build static site using MkDocs (Material theme).
    - Deploy to `gh-pages` branch.

### 2.3 Documentation Site (MkDocs)
- **Engine**: MkDocs with `mkdocs-material`.
- **Structure**:
    - Home: `README.md`
    - Guides: `docs/usage.md`, `docs/interpreting-results.md`.
    - Reference: `docs/architecture.md`, `docs/extending.md`.
- **Search**: Built-in full-text search.

## 3. Implementation Details

### 3.1 `mkdocs.yml` Configuration
```yaml
site_name: SkillSec
theme:
  name: material
  palette:
    primary: indigo
    accent: indigo
nav:
  - Home: index.md
  - Usage: usage.md
  - Architecture: architecture.md
  - Results: interpreting-results.md
  - Extending: extending.md
```

### 3.2 Release Artifacts
Users will download:
- `skillsec-<os>-<arch>` (Binary)
- `checksums.txt` (Integrity verification)

## 4. Success Criteria
- [ ] `skillsec --version` correctly shows the release tag.
- [ ] GitHub Release contains a `checksums.txt` file.
- [ ] Documentation site is live at `https://<user>.github.io/skillsec/` after a release.
