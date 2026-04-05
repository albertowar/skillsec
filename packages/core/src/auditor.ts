import { AuditReport, SkillContext, BaseCheck, CheckResult, BehavioralConfig } from './types';
import { DangerousToolsCheck } from './checks/dangerous-tools';
import { SecretScanningCheck } from './checks/secret-scanning';
import { ToolChainingCheck } from './checks/tool-chaining';
import { ExfiltrationVectorCheck } from './checks/exfiltration';
import { IndirectInjectionCheck } from './checks/indirect-injection';
import { LeastPrivilegeCheck } from './checks/least-privilege';
import { PromptInjectionCheck } from './checks/prompt-injection';
import { DependencyAuditCheck } from './checks/dependency-audit';
import { VerifiedAuthorCheck } from './checks/verified-author';
import { MaintenanceCheck } from './checks/maintenance';
import { BehavioralService } from './behavioral';
import * as crypto from 'crypto';

export class Auditor {
  private behavioralService: BehavioralService;

  constructor(
    private checks: BaseCheck[] = [],
    config: BehavioralConfig = {}
  ) {
    this.behavioralService = new BehavioralService(config);
    if (this.checks.length === 0) {
      this.checks = [
        new DangerousToolsCheck(),
        new SecretScanningCheck(),
        new ToolChainingCheck(),
        new ExfiltrationVectorCheck(),
        new IndirectInjectionCheck(),
        new LeastPrivilegeCheck(),
        new PromptInjectionCheck(),
        new DependencyAuditCheck(),
        new VerifiedAuthorCheck(),
        new MaintenanceCheck()
      ];
    }
  }

  async audit(context: SkillContext): Promise<AuditReport> {
    const outcomes = await Promise.allSettled(
      this.checks.map(c => c.run(context, this.behavioralService))
    );

    const results: CheckResult[] = outcomes.map((outcome, index) => {
      const check = this.checks[index];
      if (outcome.status === 'fulfilled') {
        return outcome.value;
      } else {
        return {
          id: check.id,
          name: check.name,
          score: 0,
          level: 'Critical',
          justification: `Check failed: ${outcome.reason}`
        };
      }
    });

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
