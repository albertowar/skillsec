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
