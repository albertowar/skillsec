# Enhanced Release and Documentation Site Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Improve the release process with versioning and checksums, and automate documentation deployment using MkDocs.

**Architecture:**
- CLI versioning via Go `ldflags`.
- integrity verification via SHA256 checksums in GitHub Releases.
- Automated documentation site built with MkDocs (Material theme) and deployed to GitHub Pages.

**Tech Stack:** Go, GitHub Actions, MkDocs, mkdocs-material.

---

### Task 1: Support Versioning in CLI

**Files:**
- Modify: `cmd/skillsec/main.go`

- [ ] **Step 1: Add Version variable and flag handling**

```go
package main

import (
    // ... existing imports
)

var Version = "dev"

func main() {
    versionFlag := flag.Bool("version", false, "Print version and exit")
    // ... existing flags
    flag.Parse()

    if *versionFlag {
        fmt.Printf("skillsec version %s\n", Version)
        os.Exit(0)
    }
    // ... rest of main
}
```

- [ ] **Step 2: Commit**

```bash
git add cmd/skillsec/main.go
git commit -m "feat: add version flag to CLI"
```

---

### Task 2: Configure MkDocs

**Files:**
- Create: `mkdocs.yml`
- Create: `docs/index.md`

- [ ] **Step 1: Create `mkdocs.yml`**

```yaml
site_name: SkillSec
site_description: Security Auditor for AI Skills
site_author: AnotherDevBoy
repo_url: https://github.com/albertowar/skillsec

theme:
  name: material
  palette:
    primary: indigo
    accent: indigo
  features:
    - navigation.tabs
    - search.highlight

nav:
  - Home: index.md
  - Usage: usage.md
  - Architecture: architecture.md
  - Results: interpreting-results.md
  - Extending: extending.md

markdown_extensions:
  - admonition
  - pymdownx.highlight:
      anchor_linenums: true
  - pymdownx.inlinehilite
  - pymdownx.snippets
  - pymdownx.superfences
```

- [ ] **Step 2: Create `docs/index.md` as a copy of `README.md`**

```bash
cp README.md docs/index.md
```

- [ ] **Step 3: Commit**

```bash
git add mkdocs.yml docs/index.md
git commit -m "docs: configure mkdocs for documentation site"
```

---

### Task 3: Update Release Workflow

**Files:**
- Modify: `.github/workflows/release.yml`

- [ ] **Step 1: Update Build and Release job**
  - Add version injection to `go build`.
  - Add checksum generation step.
  - Update `softprops/action-gh-release` to include `checksums.txt`.

```yaml
      - name: Build
        env:
          GOOS: ${{ matrix.goos }}
          GOARCH: ${{ matrix.goarch }}
        run: |
          OUTPUT_NAME=skillsec-${{ matrix.goos }}-${{ matrix.goarch }}
          if [ "${{ matrix.goos }}" = "windows" ]; then
            OUTPUT_NAME=skillsec-${{ matrix.goos }}-${{ matrix.goarch }}.exe
          fi
          go build -ldflags "-X main.Version=${{ github.event.inputs.tag_name }}" -o $OUTPUT_NAME ./cmd/skillsec

      - name: Generate Checksums
        run: sha256sum skillsec-* > checksums.txt

      - name: Upload binaries to release
        uses: softprops/action-gh-release@v2
        with:
          tag_name: ${{ github.event.inputs.tag_name }}
          files: |
            skillsec-*
            checksums.txt
```

- [ ] **Step 2: Add Documentation Deployment job**

```yaml
  deploy-docs:
    name: Deploy Documentation
    needs: release
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Python
        uses: actions/setup-python@v5
        with:
          python-version: '3.x'

      - name: Install MkDocs
        run: pip install mkdocs-material

      - name: Build and Deploy
        run: mkdocs gh-deploy --force
```

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/release.yml
git commit -m "ci: enhance release process with versioning, checksums, and docs deploy"
```
