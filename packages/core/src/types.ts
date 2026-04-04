export type RiskLevel = 'Critical' | 'High' | 'Medium' | 'Low' | 'Info';

export interface CheckResult {
  id: string;
  name: string;
  score: number; // 0-10
  level: RiskLevel;
  justification: string;
}

export interface AuditReport {
  skillHash: string;
  finalScore: number;
  results: CheckResult[];
  timestamp: string;
}

export interface SkillContext {
  raw: string;
  tools: string[];
  systemPrompt: string;
  examples: string[];
}

export interface BaseCheck {
  id: string;
  name: string;
  weight: number;
  run(context: SkillContext): Promise<CheckResult>;
}
