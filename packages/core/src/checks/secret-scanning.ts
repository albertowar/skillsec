import { BaseCheck, SkillContext, CheckResult } from '../types';

export class SecretScanningCheck implements BaseCheck {
  id = 'secret-scanning';
  name = 'Secret Scanning Audit';
  weight = 1.0;

  async run(context: SkillContext): Promise<CheckResult> {
    const patterns = [
      /\bsk-[a-zA-Z0-9]{20,}\b/, // OpenAI
      /\bghp_[a-zA-Z0-9]{36}\b/, // GitHub
      /\bsq0csp-[a-zA-Z0-9]{32}\b/, // Stripe
      /\bAKIA[0-9A-Z]{16}\b/ // AWS Access Key
    ];

    const allContent = [
      context.raw,
      context.systemPrompt,
      ...context.examples
    ].join('\n');

    const found = patterns.some(p => p.test(allContent));

    if (found) {
      return {
        id: this.id,
        name: this.name,
        score: 0,
        level: 'Critical',
        justification: 'Potential secrets detected in skill content, instructions, or examples.'
      };
    }

    return {
      id: this.id,
      name: this.name,
      score: 10,
      level: 'Low',
      justification: 'No secrets detected.'
    };
  }
}
