import { describe, it, expect } from 'vitest';
import { SecretScanningCheck } from '../../src/checks/secret-scanning';

describe('SecretScanningCheck', () => {
  const check = new SecretScanningCheck();

  it('should flag an OpenAI API key', async () => {
    const context = {
      raw: 'Some content with sk-1234567890abcdef1234567890abcdef',
      tools: [],
      systemPrompt: 'Keep your sk-1234567890abcdef1234567890abcdef secret.',
      examples: ['User: Use sk-1234567890abcdef1234567890abcdef']
    };
    const result = await check.run(context);
    expect(result.score).toBe(0);
    expect(result.level).toBe('Critical');
    expect(result.justification).toContain('Potential secrets detected');
  });

  it('should flag a GitHub personal access token', async () => {
    const context = {
      raw: 'ghp_abc123DEF456ghi789JKL012mno345PQR678',
      tools: [],
      systemPrompt: '',
      examples: []
    };
    const result = await check.run(context);
    expect(result.score).toBe(0);
    expect(result.level).toBe('Critical');
  });

  it('should pass if no secrets are present', async () => {
    const context = {
      raw: 'Just some safe content.',
      tools: [],
      systemPrompt: 'Act as a helper.',
      examples: ['User: hi']
    };
    const result = await check.run(context);
    expect(result.score).toBe(10);
    expect(result.level).toBe('Low');
    expect(result.justification).toBe('No secrets detected.');
  });
});
