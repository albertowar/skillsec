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
