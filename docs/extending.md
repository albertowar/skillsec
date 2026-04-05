# Extending SkillAuditAI

## Adding a New Check

1. Create a new file in `packages/core/src/checks/`.
2. Implement the `BaseCheck` interface:

```typescript
import { BaseCheck, SkillContext, CheckResult } from '../types';

export class MyNewCheck implements BaseCheck {
  id = 'my-check';
  name = 'My Security Audit';
  weight = 0.5;

  async run(context: SkillContext): Promise<CheckResult> {
    // Audit logic here
  }
}
```

3. Register the check in `packages/core/src/auditor.ts`.
4. Add tests in `packages/core/tests/checks/`.
