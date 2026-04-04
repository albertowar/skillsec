import { describe, it, expect } from 'vitest';
import { parseSkill } from '../src/parser';

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

  it('should extract examples and handle sections robustly', () => {
    const content = `
# Test Skill
## Tools
- \`tool1\`
- tool2

## Instructions
Main instructions here.
Can be multiple lines.

## Examples
User: Hello
Assistant: Hi there!
---
User: How are you?
Assistant: I am good.
`;
    const context = parseSkill(content);
    expect(context.tools).toEqual(['tool1', 'tool2']);
    expect(context.systemPrompt).toContain('Main instructions here.');
    expect(context.systemPrompt).toContain('Can be multiple lines.');
    expect(context.examples).toHaveLength(2);
    expect(context.examples[0]).toContain('User: Hello');
    expect(context.examples[1]).toContain('User: How are you?');
  });
});
