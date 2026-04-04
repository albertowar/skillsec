# SkillAuditAI Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a security auditor for AI skills (SKILL.md) that provides a weighted safety scorecard (0-10) using a modular engine.

**Architecture:** A decoupled TypeScript architecture with a core `Auditor` engine that orchestrates independent `Check` modules, returning a serializable `AuditReport`.

**Tech Stack:** Node.js, TypeScript, Vitest (testing), Zod (validation), Chalk/Table (CLI formatting).

---

### Task 1: Project Initialization & Types

**Files:**
- Create: `package.json`
- Create: `tsconfig.json`
- Create: `src/core/types.ts`

- [ ] **Step 1: Create `package.json`**
```json
{
  "name": "@skillauditai/core",
  "version": "0.1.0",
  "description": "Security auditor for AI skills",
  "main": "dist/index.js",
  "types": "dist/index.d.ts",
  "scripts": {
    "build": "tsc",
    "test": "vitest run",
    "lint": "eslint . --ext .ts"
  },
  "dependencies": {
    "zod": "^3.22.0",
    "chalk": "^4.1.2",
    "cli-table3": "^0.6.3"
  },
  "devDependencies": {
    "typescript": "^5.2.2",
    "vitest": "^1.0.0",
    "@types/node": "^20.8.0"
  }
}
```

- [ ] **Step 2: Create `tsconfig.json`**
```json
{
  "compilerOptions": {
    "target": "ESNext",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "outDir": "./dist",
    "rootDir": "./src",
    "strict": true,
    "declaration": true,
    "esModuleInterop": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "**/*.spec.ts"]
}
```

- [ ] **Step 3: Define core types in `src/core/types.ts`**
```typescript
export type RiskLevel = 'Critical' | 'High' | 'Medium' | 'Low' | 'Info';

export interface CheckResult {
  id: string;
  name: string;
  score: number; // 0-10
  level: RiskLevel;
  justification: string;
}

export interface AuditReport {
  skillHash: string;
  finalScore: number;
  results: CheckResult[];
  timestamp: string;
}

export interface SkillContext {
  raw: string;
  tools: string[];
  systemPrompt: string;
  examples: string[];
}

export interface BaseCheck {
  id: string;
  name: string;
  weight: number;
  run(context: SkillContext): Promise<CheckResult>;
}
```

- [ ] **Step 4: Commit**
```bash
git add package.json tsconfig.json src/core/types.ts
git commit -m "chore: initialize project and core types"
```

---

### Task 2: Skill Parser Implementation

**Files:**
- Create: `src/core/parser.ts`
- Test: `tests/core/parser.spec.ts`

- [ ] **Step 1: Write failing test for parser**
```typescript
import { describe, it, expect } from 'vitest';
import { parseSkill } from '../../src/core/parser';

describe('parseSkill', () => {
  it('should extract tools and system prompt from SKILL.md', () => {
    const content = `
# Test Skill
## Tools
- run_shell_command
- read_file

## Instructions
Act as a helper. Use delimiters.

## Examples
User: hi
`;
    const context = parseSkill(content);
    expect(context.tools).toContain('run_shell_command');
    expect(context.systemPrompt).toContain('Act as a helper');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**
Run: `npx vitest tests/core/parser.spec.ts`

- [ ] **Step 3: Implement `parseSkill`**
```typescript
import { SkillContext } from './types';

export function parseSkill(content: string): SkillContext {
  const tools: string[] = [];
  const toolMatches = content.match(/## Tools\n([\s\S]*?)\n##/);
  if (toolMatches) {
    const lines = toolMatches[1].split('\n');
    lines.forEach(line => {
      const match = line.match(/-\s*`?([\w_]+)`?/);
      if (match) tools.push(match[1]);
    });
  }

  const instructionsMatch = content.match(/## Instructions\n([\s\S]*?)\n##/);
  const systemPrompt = instructionsMatch ? instructionsMatch[1].trim() : '';

  return {
    raw: content,
    tools,
    systemPrompt,
    examples: [] // Placeholder for now
  };
}
```

- [ ] **Step 4: Run test to verify it passes**
Run: `npx vitest tests/core/parser.spec.ts`

- [ ] **Step 5: Commit**
```bash
git add src/core/parser.ts tests/core/parser.spec.ts
git commit -m "feat: implement skill parser"
```

---

### Task 3: Dangerous Tools Check

**Files:**
- Create: `src/core/checks/dangerous-tools.ts`
- Test: `tests/core/checks/dangerous-tools.spec.ts`

- [ ] **Step 1: Write failing test for `dangerous-tools` check**
```typescript
import { describe, it, expect } from 'vitest';
import { DangerousToolsCheck } from '../../../src/core/checks/dangerous-tools';

describe('DangerousToolsCheck', () => {
  it('should score low if run_shell_command is present', async () => {
    const check = new DangerousToolsCheck();
    const result = await check.run({ tools: ['run_shell_command'], raw: '', systemPrompt: '', examples: [] });
    expect(result.score).toBeLessThan(5);
    expect(result.level).toBe('Critical');
  });
});
```

- [ ] **Step 2: Implement `DangerousToolsCheck`**
```typescript
import { BaseCheck, SkillContext, CheckResult } from '../types';

export class DangerousToolsCheck implements BaseCheck {
  id = 'dangerous-tools';
  name = 'Dangerous Tools Audit';
  weight = 1.0;

  async run(context: SkillContext): Promise<CheckResult> {
    const dangerous = ['run_shell_command', 'write_file', 'delete_file'];
    const found = context.tools.filter(t => dangerous.includes(t));

    if (found.length > 0) {
      return {
        id: this.id,
        name: this.name,
        score: 0,
        level: 'Critical',
        justification: `Skill requests highly dangerous tools: ${found.join(', ')}`
      };
    }

    return {
      id: this.id,
      name: this.name,
      score: 10,
      level: 'Low',
      justification: 'No dangerous tools detected.'
    };
  }
}
```

- [ ] **Step 3: Run test and commit**
Run: `npx vitest tests/core/checks/dangerous-tools.spec.ts`
```bash
git add src/core/checks/dangerous-tools.ts tests/core/checks/dangerous-tools.spec.ts
git commit -m "feat: add dangerous-tools check"
```

---

### Task 4: Auditor Engine

**Files:**
- Create: `src/core/auditor.ts`
- Test: `tests/core/auditor.spec.ts`

- [ ] **Step 1: Implement `Auditor` engine**
```typescript
import { AuditReport, SkillContext, BaseCheck } from './types';
import { DangerousToolsCheck } from './checks/dangerous-tools';
import { crypto } from 'crypto';

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
      skillHash: 'hash-placeholder', // Implement actual hash if needed
      finalScore: weightedScore / totalWeight,
      results,
      timestamp: new Date().toISOString()
    };
  }
}
```

- [ ] **Step 2: Commit**
```bash
git add src/core/auditor.ts
git commit -m "feat: implement auditor engine"
```

---

### Task 5: CLI Interface

**Files:**
- Create: `src/index.ts`

- [ ] **Step 1: Implement basic CLI**
```typescript
import { readFileSync } from 'fs';
import { parseSkill } from './core/parser';
import { Auditor } from './core/auditor';
import chalk from 'chalk';

async function run() {
  const filePath = process.argv[2];
  if (!filePath) {
    console.error('Usage: skillaudit <path-to-skill.md>');
    process.exit(1);
  }

  const content = readFileSync(filePath, 'utf-8');
  const context = parseSkill(content);
  const auditor = new Auditor();
  const report = await auditor.audit(context);

  console.log(chalk.bold(`\nSkillAuditAI Report - Score: ${report.finalScore.toFixed(1)}/10\n`));
  report.results.forEach(r => {
    const color = r.score > 7 ? chalk.green : r.score > 4 ? chalk.yellow : chalk.red;
    console.log(`${color(`[${r.level}]`)} ${chalk.bold(r.name)}: ${r.score}/10`);
    console.log(`Justification: ${r.justification}\n`);
  });
}

run().catch(console.error);
```

- [ ] **Step 2: Commit**
```bash
git add src/index.ts
git commit -m "feat: basic CLI interface"
```
