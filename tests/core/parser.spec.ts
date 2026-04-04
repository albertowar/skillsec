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
    expect(context.tools).toContain('read_file');
    expect(context.systemPrompt).toContain('Act as a helper');
  });
});
