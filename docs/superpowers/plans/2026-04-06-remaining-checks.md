# SkillAuditAI Remaining Checks Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Implement the remaining security checks from the TypeScript codebase in the new Go architecture.

**Architecture:** Each check is a separate file in `internal/checks/`, implementing the `Check` interface and registered in `internal/checks/check.go`.

**Tech Stack:** Go 1.21+, `langchaingo` (for behavioral checks).

---

### Task 1: Implement Static Checks (Secret Scanning, Dependency Audit, Verified Author, Maintenance)

**Files:**
- Create: `internal/checks/secret_scanning.go`
- Create: `internal/checks/dependency_audit.go`
- Create: `internal/checks/verified_author.go`
- Create: `internal/checks/maintenance.go`
- Modify: `internal/checks/check.go`

- [ ] **Step 1: Implement SecretScanningCheck** (Regex-based scanning for keys/tokens in the raw content)
- [ ] **Step 2: Implement DependencyAuditCheck** (Audit listed dependencies for known risks)
- [ ] **Step 3: Implement VerifiedAuthorCheck** (Check if author metadata exists and is verified)
- [ ] **Step 4: Implement MaintenanceCheck** (Check last updated date and versioning)
- [ ] **Step 5: Register all static checks in `AllChecks()`**
- [ ] **Step 6: Commit**

---

### Task 2: Implement Behavioral Checks (Prompt Injection, Exfiltration, Indirect Injection, Least Privilege, Tool Chaining)

**Files:**
- Create: `internal/checks/prompt_injection.go`
- Create: `internal/checks/exfiltration.go`
- Create: `internal/checks/indirect_injection.go`
- Create: `internal/checks/least_privilege.go`
- Create: `internal/checks/tool_chaining.go`
- Modify: `internal/checks/check.go`

- [ ] **Step 1: Implement PromptInjectionCheck** (Use `behavioral.Service` to probe for injection vulnerabilities)
- [ ] **Step 2: Implement ExfiltrationCheck** (Check for data exfiltration patterns in system prompt and behavior)
- [ ] **Step 3: Implement IndirectInjectionCheck** (Probe for vulnerabilities to indirect injection)
- [ ] **Step 4: Implement LeastPrivilegeCheck** (Verify tool usage adheres to least privilege principles)
- [ ] **Step 5: Implement ToolChainingCheck** (Analyze risks from chaining multiple tools)
- [ ] **Step 6: Register all behavioral checks in `AllChecks()`**
- [ ] **Step 7: Commit**
