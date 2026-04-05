import { BaseCheck, SkillContext, CheckResult } from '../types';

export class PromptInjectionCheck implements BaseCheck {
  id = 'prompt-injection';
  name = 'Prompt Injection Audit';
  weight = 0.8;

  async run(context: SkillContext): Promise<CheckResult> {
    const prompt = context.systemPrompt.toLowerCase();
    
    const delimiters = ['"""', '---', '<context>', '###', '```'];
    const hasDelimiters = delimiters.some(d => prompt.includes(d));

    const safetyPhrases = [
      'do not reveal',
      'ignore previous',
      'stay in character',
      'never mention',
      'private instructions',
      'strictly follow'
    ];
    const hasSafetyPhrases = safetyPhrases.some(p => prompt.includes(p));

    if (hasDelimiters && hasSafetyPhrases) {
      return {
        id: this.id,
        name: this.name,
        score: 10,
        level: 'Low',
        justification: 'Uses both structural delimiters and explicit safety constraints.'
      };
    } else if (hasDelimiters || hasSafetyPhrases) {
      return {
        id: this.id,
        name: this.name,
        score: 6,
        level: 'Medium',
        justification: `Missing either structural delimiters or safety constraints: ${hasDelimiters ? 'Missing constraints' : 'Missing delimiters'}`
      };
    } else {
      return {
        id: this.id,
        name: this.name,
        score: 2,
        level: 'High',
        justification: 'No structural delimiters or safety constraints detected in instructions.'
      };
    }
  }
}
