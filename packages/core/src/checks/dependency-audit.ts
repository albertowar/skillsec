import { BaseCheck, SkillContext, CheckResult } from '../types';

export class DependencyAuditCheck implements BaseCheck {
  id = 'dependency-audit';
  name = 'Dependency Audit';
  weight = 0.5;

  async run(context: SkillContext): Promise<CheckResult> {
    const allContent = [
      context.systemPrompt,
      ...context.examples
    ].join('\n').toLowerCase();

    const installers = ['pip install', 'npm install', 'cargo add', 'gem install', 'apt install'];
    const hasInstallers = installers.filter(i => allContent.includes(i));

    const imports = [
      /import\s+['"]?[\w-]+['"]?/,
      /require\s*\(['"]?[\w-]+['"]?\)/,
      /\[.*?\]\(.*?\.md\)/ // Skill links
    ];
    const foundImports = imports.filter(r => r.test(allContent));

    if (hasInstallers.length > 0) {
      return {
        id: this.id,
        name: this.name,
        score: 4,
        level: 'Medium',
        justification: `Skill contains external library installers: ${hasInstallers.join(', ')}`
      };
    } else if (foundImports.length > 0) {
      return {
        id: this.id,
        name: this.name,
        score: 7,
        level: 'Low',
        justification: 'Skill references external skills or libraries but lacks explicit installers.'
      };
    }

    return {
      id: this.id,
      name: this.name,
      score: 10,
      level: 'Low',
      justification: 'No external dependencies detected.'
    };
  }
}
