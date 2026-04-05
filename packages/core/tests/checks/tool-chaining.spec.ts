import { describe, it, expect } from 'vitest';
import { ToolChainingCheck } from '../../src/checks/tool-chaining';
import { SkillContext } from '../../src/types';

describe('ToolChainingCheck', () => {
  const check = new ToolChainingCheck();

  it('should flag Critical risk for Source + External Sink without HITL', async () => {
    const context: SkillContext = {
      tools: ['read_file', 'api_post'],
      systemPrompt: 'Read sensitive files and send them to the API.',
      raw: '',
      examples: []
    };
    const result = await check.run(context);
    expect(result.level).toBe('Critical');
    expect(result.score).toBe(0);
  });

  it('should flag High risk for Source + External Sink with HITL', async () => {
    const context: SkillContext = {
      tools: ['read_file', 'api_post'],
      systemPrompt: 'Read the file and ask the user for approval before sending to API.',
      raw: '',
      examples: []
    };
    const result = await check.run(context);
    expect(result.level).toBe('High');
  });

  it('should flag High risk for Source + Internal Sink without HITL', async () => {
    const context: SkillContext = {
      tools: ['list_directory', 'write_file'],
      systemPrompt: 'List files and write results to local log.',
      raw: '',
      examples: []
    };
    const result = await check.run(context);
    expect(result.level).toBe('High');
  });

  it('should flag Medium risk for Source + Internal Sink with HITL', async () => {
    const context: SkillContext = {
      tools: ['get_history', 'append_to_log'],
      systemPrompt: 'Get shell history and confirm with user before appending to logs.',
      raw: '',
      examples: []
    };
    const result = await check.run(context);
    expect(result.level).toBe('Medium');
  });

  it('should flag Low risk if no dangerous chain is detected', async () => {
    const context: SkillContext = {
      tools: ['read_file'],
      systemPrompt: 'Just read the file and display its content.',
      raw: '',
      examples: []
    };
    const result = await check.run(context);
    expect(result.level).toBe('Low');
    expect(result.score).toBe(10);
  });
});
