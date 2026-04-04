import { SkillContext } from './types';

export function parseSkill(content: string): SkillContext {
  const getSection = (name: string) => {
    const regex = new RegExp(`## ${name}\\s+([\\s\\S]*?)(?=##|$)`, 'i');
    const match = content.match(regex);
    return match ? match[1].trim() : '';
  };

  const toolsSection = getSection('Tools');
  const tools: string[] = [];
  if (toolsSection) {
    const lines = toolsSection.split('\n');
    lines.forEach(line => {
      const match = line.match(/-\s*`?([\w_]+)`?/);
      if (match) tools.push(match[1]);
    });
  }

  const systemPrompt = getSection('Instructions');
  
  const examplesSection = getSection('Examples');
  const examples = examplesSection 
    ? examplesSection.split('\n---\n').map(e => e.trim()).filter(Boolean)
    : [];

  return {
    raw: content,
    tools,
    systemPrompt,
    examples
  };
}
