# Spec: Semantic Versioning Release Process

**Date:** 2026-04-06
**Status:** Approved
**Topic:** Automating version calculation based on Git tags and user input (patch, minor, major).

## 1. Purpose
Simplify the release process by automating semantic version increments. This ensures consistency in versioning and removes the need for manual tag management during releases.

## 2. Architecture

### 2.1 Manual Trigger with Version Choice
- **Trigger**: `workflow_dispatch`.
- **Input**: `increment` (Choice: `patch`, `minor`, `major`).

### 2.2 Version Calculation Logic
1.  **Fetch Tags**: `git fetch --tags`.
2.  **Identify Latest**: Get the latest tag matching `v*`.
3.  **Defaulting**: If no tags exist, start from `v0.0.0`.
4.  **Increment**:
    - `patch`: `1.2.3` -> `1.2.4`
    - `minor`: `1.2.3` -> `1.3.0`
    - `major`: `1.2.3` -> `2.0.0`

### 2.3 Git Operations
- **Identity**: Use `github-actions[bot]` for Git commits/tags.
- **Tagging**: Create a new tag locally.
- **Pushing**: Push the new tag to the `origin` repository.

## 3. Workflow Implementation Details

### 3.1 GitHub Action Updates
The `release.yml` will be updated to:
- Use a dedicated step to calculate the next version string.
- Set the calculated version as a workflow output.
- Pass this output to subsequent build and release steps.

### 3.2 Security and Permissions
- **Permissions**: Ensure the `GITHUB_TOKEN` has `contents: write` permission to push tags and create releases.

## 4. Success Criteria
- [ ] Manual release trigger shows a dropdown for version increment.
- [ ] Workflow successfully increments the current version.
- [ ] New tag is pushed to the repository automatically.
- [ ] GitHub Release is created with the new version tag.
