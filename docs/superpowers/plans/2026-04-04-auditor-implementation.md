# Auditor Engine Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Implement the `Auditor` class to orchestrate multiple checks and calculate a weighted final score.

**Architecture:** The `Auditor` class will maintain a list of `BaseCheck` implementations, execute them in parallel, and aggregate their results into an `AuditReport`.

**Tech Stack:** TypeScript, Vitest, Node.js `crypto` module.

---

### Task 1: Create Auditor class and initial test

**Files:**
- Create: `src/core/auditor.ts`
- Test: `tests/core/auditor.spec.ts`

- [ ] **Step 1: Write the failing test**

```typescript
import { describe, it, expect } from 'vitest';
import { Auditor } from '../../src/core/auditor';

describe('Auditor', () => {
  it('should perform an audit and return a report', async () => {
    const auditor = new Auditor();
    const context = {
      raw: 'test skill',
      tools: [],
      systemPrompt: '',
      examples: []
    };
    const report = await auditor.audit(context);
    
    expect(report.skillHash).toBeDefined();
    expect(report.finalScore).toBeGreaterThanOrEqual(0);
    expect(report.results.length).toBeGreaterThan(0);
    expect(report.timestamp).toBeDefined();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `npx vitest tests/core/auditor.spec.ts`
Expected: FAIL (Auditor not found)

- [ ] **Step 3: Implement Auditor class**

```typescript
import { AuditReport, SkillContext, BaseCheck } from './types';
import { DangerousToolsCheck } from './checks/dangerous-tools';
import * as crypto from 'crypto';

export class Auditor {
  private checks: BaseCheck[] = [new DangerousToolsCheck()];

  async audit(context: SkillContext): Promise<AuditReport> {
    const results = await Promise.all(this.checks.map(c => c.run(context)));
    
    const weightedScore = results.reduce((acc, r) => {
      const check = this.checks.find(c => c.id === r.id);
      return acc + (r.score * (check?.weight || 1));
    }, 0);
    
    const totalWeight = this.checks.reduce((acc, c) => acc + c.weight, 0);

    return {
      skillHash: crypto.createHash('sha256').update(context.raw).digest('hex'),
      finalScore: totalWeight > 0 ? weightedScore / totalWeight : 10,
      results,
      timestamp: new Date().toISOString()
    };
  }
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `npx vitest tests/core/auditor.spec.ts`
Expected: PASS

- [ ] **Step 5: Add more comprehensive tests**

```typescript
  it('should calculate weighted score correctly', async () => {
    const auditor = new Auditor();
    const context = {
      tools: ['run_shell_command'],
      raw: 'dangerous skill',
      systemPrompt: '',
      examples: []
    };
    const report = await auditor.audit(context);
    
    // DangerousToolsCheck has weight 1.0 and returns 0 for run_shell_command
    expect(report.finalScore).toBe(0);
  });
```

- [ ] **Step 6: Run tests and verify they pass**

Run: `npx vitest tests/core/auditor.spec.ts`
Expected: PASS

- [ ] **Step 7: Commit**

```bash
git add src/core/auditor.ts tests/core/auditor.spec.ts
git commit -m "feat: implement auditor engine"
```
