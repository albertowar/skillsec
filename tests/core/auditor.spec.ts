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
});
