import { describe, it, expect } from 'vitest';
import { parseFlags } from '../src/index';

describe('CLI Flag Parsing', () => {
  it('should parse --api-key correctly', () => {
    const flags = parseFlags(['--api-key', 'test-key', 'skill.md']);
    expect(flags.config.apiKey).toBe('test-key');
  });

  it('should parse --model correctly', () => {
    const flags = parseFlags(['--model', 'gpt-4', 'skill.md']);
    expect(flags.config.modelName).toBe('gpt-4');
  });

  it('should parse --provider correctly', () => {
    const flags = parseFlags(['--provider', 'openai', 'skill.md']);
    expect(flags.config.provider).toBe('openai');
  });

  it('should parse --base-url correctly', () => {
    const flags = parseFlags(['--base-url', 'https://api.openai.com/v1', 'skill.md']);
    expect(flags.config.baseUrl).toBe('https://api.openai.com/v1');
  });

  it('should parse --format correctly', () => {
    const flags = parseFlags(['--format', 'json', 'skill.md']);
    expect(flags.format).toBe('json');
  });

  it('should parse filePath correctly', () => {
    const flags = parseFlags(['--format', 'json', 'skill.md']);
    expect(flags.filePath).toBe('skill.md');
  });

  it('should return undefined filePath if missing', () => {
    const flags = parseFlags(['--format', 'json']);
    expect(flags.filePath).toBeUndefined();
  });
});
