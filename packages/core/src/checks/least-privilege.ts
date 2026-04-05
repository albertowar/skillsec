import { BaseCheck, SkillContext, CheckResult } from '../types';

export class LeastPrivilegeCheck implements BaseCheck {
  id = 'least-privilege';
  name = 'Least Privilege Audit';
  weight = 0.8;

  async run(context: SkillContext): Promise<CheckResult> {
    const dangerous = ['run_shell_command', 'write_file', 'delete_file'];
    const found = context.tools.filter(t => dangerous.includes(t));

    if (found.length === 0) {
      return {
        id: this.id,
        name: this.name,
        score: 10,
        level: 'Low',
        justification: 'No high-risk tools requested.'
      };
    }

    const keywords: Record<string, string[]> = {
      'run_shell_command': ['bash', 'shell', 'command', 'terminal', 'execute', 'cli', 'run'],
      'write_file': ['save', 'write', 'create', 'file', 'output', 'export'],
      'delete_file': ['remove', 'delete', 'unlink', 'cleanup', 'rm']
    };

    const prompt = context.systemPrompt.toLowerCase();
    const unjustified = found.filter(tool => {
      const toolKeywords = keywords[tool] || [];
      return !toolKeywords.some(k => prompt.includes(k));
    });

    if (unjustified.length > 0) {
      return {
        id: this.id,
        name: this.name,
        score: 5,
        level: 'High',
        justification: `High-risk tools requested without clear justification in instructions: ${unjustified.join(', ')}`
      };
    }

    return {
      id: this.id,
      name: this.name,
      score: 10,
      level: 'Low',
      justification: 'Requested tools are justified by the skill instructions.'
    };
  }
}
