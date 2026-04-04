import { describe, it, expect } from 'vitest';
import { Auditor } from '../src/auditor';

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

  it('should calculate weighted score correctly when all checks fail', async () => {
    const auditor = new Auditor();
    const context = {
      tools: ['run_shell_command'],
      raw: 'dangerous skill with secret sk-12345678901234567890',
      systemPrompt: '',
      examples: []
    };
    const report = await auditor.audit(context);
    
    // Both DangerousToolsCheck and SecretScanningCheck fail
    expect(report.finalScore).toBe(0);
  });
});
