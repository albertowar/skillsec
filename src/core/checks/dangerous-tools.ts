import { BaseCheck, SkillContext, CheckResult } from '../types';

export class DangerousToolsCheck implements BaseCheck {
  id = 'dangerous-tools';
  name = 'Dangerous Tools Audit';
  weight = 1.0;

  async run(context: SkillContext): Promise<CheckResult> {
    const dangerous = ['run_shell_command', 'write_file', 'delete_file'];
    const found = context.tools.filter(t => dangerous.includes(t));

    if (found.length > 0) {
      return {
        id: this.id,
        name: this.name,
        score: 0,
        level: 'Critical',
        justification: `Skill requests highly dangerous tools: ${found.join(', ')}`
      };
    }

    return {
      id: this.id,
      name: this.name,
      score: 10,
      level: 'Low',
      justification: 'No dangerous tools detected.'
    };
  }
}
