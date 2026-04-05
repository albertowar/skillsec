import { BaseCheck, SkillContext, CheckResult } from '../types';

export class VerifiedAuthorCheck implements BaseCheck {
  id = 'verified-author';
  name = 'Verified Author Audit';
  weight = 0.5;

  async run(context: SkillContext): Promise<CheckResult> {
    const meta = context.metadata?.author;

    if (!meta || (!meta.name && !meta.email)) {
      return {
        id: this.id,
        name: this.name,
        score: 0,
        level: 'Medium',
        justification: 'No author attribution found.'
      };
    }

    if (meta.isVerified) {
      return {
        id: this.id,
        name: this.name,
        score: 10,
        level: 'Low',
        justification: `Verified author: ${meta.name} <${meta.email}>`
      };
    }

    return {
      id: this.id,
      name: this.name,
      score: 7,
      level: 'Low',
      justification: `Author identified but not verified: ${meta.name} <${meta.email}>`
    };
  }
}
