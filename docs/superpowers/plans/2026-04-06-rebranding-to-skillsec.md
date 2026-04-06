# SkillSec Rebranding Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Perform a full technical and visual rebrand from SkillSec to SkillSec, including Go module updates, binary renaming, and documentation refresh.

**Architecture:** Systematic find-and-replace for identifiers, filesystem move for the entry point, and Go module path migration.

**Tech Stack:** Go, Bash, Markdown.

---

### Task 1: Go Module & Import Migration

**Files:**
- Modify: `go.mod`
- Modify: All `.go` files in `internal/`, `pkg/`, and `cmd/`
- Modify: `go.sum`

- [ ] **Step 1: Update go.mod module path**

Update `go.mod` to use the new module path.
```go
module github.com/albertowar/skillsec
```

- [ ] **Step 2: Mass update Go imports**

Run a command to replace all internal import paths.
Run: `find . -name "*.go" -type f -exec sed -i 's|github.com/albertowar/skillsec|github.com/albertowar/skillsec|g' {} +`

- [ ] **Step 3: Run go mod tidy**

Run: `go mod tidy`
Expected: Success, `go.sum` updated.

- [ ] **Step 4: Verify compilation**

Run: `go build ./...`
Expected: Success.

- [ ] **Step 5: Commit**

```bash
git add go.mod go.sum $(find . -name "*.go")
git commit -m "refactor: rename Go module to skillsec and update imports"
```

### Task 2: Entry Point & Binary Renaming

**Files:**
- Rename: `cmd/skillsec/` -> `cmd/skillsec/`
- Modify: `cmd/skillsec/main.go`

- [ ] **Step 1: Rename the cmd directory**

Run: `mv cmd/skillsec cmd/skillsec`

- [ ] **Step 2: Update main.go branding**

Modify `cmd/skillsec/main.go` to update usage strings and the version flag.
```go
// Replace occurrences of "skillsec" with "skillsec" in fmt.Println and flag definitions
```

- [ ] **Step 3: Verify build with new name**

Run: `go build -o skillsec ./cmd/skillsec`
Expected: A `skillsec` binary is created.

- [ ] **Step 4: Commit**

```bash
git add cmd/skillsec
git commit -m "feat: rename entry point directory to cmd/skillsec"
```

### Task 3: Identity, Environment Variables & UI

**Files:**
- Modify: `cmd/skillsec/main.go`
- Modify: `internal/engine/auditor.go`
- Modify: `docs/usage.md`
- Modify: `README.md`

- [ ] **Step 1: Update environment variable**

Change `SKILLSEC_API_KEY` to `SKILLSEC_API_KEY` in `cmd/skillsec/main.go` and documentation.

- [ ] **Step 2: Update CLI display names**

Replace "SkillSec" with "SkillSec" in `cmd/skillsec/main.go` output.

- [ ] **Step 3: Commit**

```bash
git add cmd/skillsec/main.go docs/usage.md README.md
git commit -m "feat: update env vars and CLI display to SkillSec"
```

### Task 4: Documentation & Site Rebranding

**Files:**
- Modify: `mkdocs.yml`
- Modify: `README.md`
- Modify: `AGENTS.md`
- Modify: `docs/**/*.md`

- [ ] **Step 1: Mass replace in Markdown files**

Run: `find . -name "*.md" -type f -exec sed -i 's/SkillSec/SkillSec/g' {} +`
Run: `find . -name "*.md" -type f -exec sed -i 's/skillsec/skillsec/g' {} +`

- [ ] **Step 2: Update mkdocs.yml**

Update `site_name`, `repo_url`, and any other links.
```yaml
site_name: SkillSec
repo_url: https://github.com/albertowar/skillsec
```

- [ ] **Step 3: Update .gitignore**

Change `/skillsec` to `/skillsec`.

- [ ] **Step 4: Commit**

```bash
git add .
git commit -m "docs: full rebrand of documentation and site config"
```

### Task 5: CI/CD & Final Verification

**Files:**
- Modify: `.github/workflows/release.yml`

- [ ] **Step 1: Update Release Workflow**

Update binary names and build paths in `.github/workflows/release.yml`.
```yaml
- run: go build -o skillsec-${{ matrix.goos }}-${{ matrix.goarch }} ./cmd/skillsec
```

- [ ] **Step 2: Final Build Test**

Run: `go build -o skillsec ./cmd/skillsec`
Run: `./skillsec --help`
Expected: Output shows "skillsec" and "SkillSec".

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/release.yml
git commit -m "ci: update release workflow for skillsec binary"
```
