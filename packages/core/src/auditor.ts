import { AuditReport, SkillContext, BaseCheck } from './types';
import { DangerousToolsCheck } from './checks/dangerous-tools';
import * as crypto from 'crypto';

export class Auditor {
  private checks: BaseCheck[] = [new DangerousToolsCheck()];

  async audit(context: SkillContext): Promise<AuditReport> {
    const results = await Promise.all(this.checks.map(c => c.run(context)));
    
    const weightedScore = results.reduce((acc, r) => {
      const check = this.checks.find(c => c.id === r.id);
      return acc + (r.score * (check?.weight || 1));
    }, 0);
    
    const totalWeight = this.checks.reduce((acc, c) => acc + c.weight, 0);

    return {
      skillHash: crypto.createHash('sha256').update(context.raw).digest('hex'),
      finalScore: totalWeight > 0 ? weightedScore / totalWeight : 10,
      results,
      timestamp: new Date().toISOString()
    };
  }
}
