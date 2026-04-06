# Semantic Versioning Release Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Automate semantic version calculation and tagging in the GitHub Release workflow.

**Architecture:** 
- `workflow_dispatch` with `increment` input (patch, minor, major).
- Version calculation logic using Git tags.
- Automatic tagging and pushing back to `origin`.
- Multi-architecture build and release using the new tag.

**Tech Stack:** GitHub Actions, Git, Go.

---

### Task 1: Update Workflow Inputs and Version Calculation

**Files:**
- Modify: `.github/workflows/release.yml`

- [ ] **Step 1: Update inputs and add version calculation job**

```yaml
on:
  workflow_dispatch:
    inputs:
      increment:
        description: 'Version increment'
        required: true
        default: 'patch'
        type: choice
        options:
          - patch
          - minor
          - major

jobs:
  checks:
    # ... existing checks job ...

  calculate-version:
    name: Calculate Next Version
    needs: checks
    runs-on: ubuntu-latest
    outputs:
      next_version: ${{ steps.bump.outputs.next_version }}
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Get latest tag and bump
        id: bump
        run: |
          # Get latest tag, default to v0.0.0 if none found
          LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
          VERSION=${LATEST_TAG#v}
          
          # Parse version components
          IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION"
          
          INCREMENT="${{ github.event.inputs.increment }}"
          
          if [ "$INCREMENT" == "major" ]; then
            MAJOR=$((MAJOR + 1))
            MINOR=0
            PATCH=0
          elif [ "$INCREMENT" == "minor" ]; then
            MINOR=$((MINOR + 1))
            PATCH=0
          else
            PATCH=$((PATCH + 1))
          fi
          
          NEXT_VERSION="v$MAJOR.$MINOR.$PATCH"
          echo "next_version=$NEXT_VERSION" >> $GITHUB_OUTPUT
          echo "Calculated next version: $NEXT_VERSION"
```

- [ ] **Step 2: Commit**

```bash
git add .github/workflows/release.yml
git commit -m "ci: add semantic versioning input and calculation job"
```

---

### Task 2: Implement Tagging and Release Integration

**Files:**
- Modify: `.github/workflows/release.yml`

- [ ] **Step 1: Update release job to tag and build with new version**

Modify the `release` job to depend on `calculate-version`, create the tag, and use the version for building and release creation.

```yaml
  release:
    name: Build and Release
    needs: calculate-version
    runs-on: ubuntu-latest
    env:
      NEXT_VERSION: ${{ needs.calculate-version.outputs.next_version }}
    # ... strategy section remains the same ...
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Tag version
        if: matrix.goos == 'linux' && matrix.goarch == 'amd64' # Only run once
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"
          git tag $NEXT_VERSION
          git push origin $NEXT_VERSION

      - name: Wait for tag propagation
        run: sleep 5 # Small buffer for eventual consistency

      - name: Build
        env:
          GOOS: ${{ matrix.goos }}
          GOARCH: ${{ matrix.goarch }}
        run: |
          OUTPUT_NAME=skillaudit-${{ matrix.goos }}-${{ matrix.goarch }}
          if [ "${{ matrix.goos }}" = "windows" ]; then
            OUTPUT_NAME=skillaudit-${{ matrix.goos }}-${{ matrix.goarch }}.exe
          fi
          go build -ldflags "-X main.Version=$NEXT_VERSION" -o $OUTPUT_NAME ./cmd/skillaudit

      - name: Generate Checksums
        run: sha256sum skillaudit-* > checksums.txt

      - name: Upload binaries to release
        uses: softprops/action-gh-release@v2
        with:
          tag_name: ${{ env.NEXT_VERSION }}
          files: |
            skillaudit-*
            checksums.txt
```

- [ ] **Step 2: Commit**

```bash
git add .github/workflows/release.yml
git commit -m "ci: automate tagging and integrate version into build/release"
```
