import { BaseCheck, SkillContext, CheckResult } from '../types';

export class MaintenanceCheck implements BaseCheck {
  id = 'maintenance';
  name = 'Maintenance Audit';
  weight = 0.3;

  async run(context: SkillContext): Promise<CheckResult> {
    const meta = context.metadata?.maintenance;

    if (!meta || !meta.lastUpdated) {
      return {
        id: this.id,
        name: this.name,
        score: 0,
        level: 'Info',
        justification: 'No maintenance history available.'
      };
    }

    const lastUpdate = new Date(meta.lastUpdated);
    const now = new Date();
    const diffMs = now.getTime() - lastUpdate.getTime();
    const diffDays = diffMs / (1000 * 60 * 60 * 24);

    if (diffDays < 90) { // < 3 months
      return {
        id: this.id,
        name: this.name,
        score: 10,
        level: 'Low',
        justification: `Skill is recently maintained (last update ${Math.round(diffDays)} days ago).`
      };
    } else if (diffDays < 365) { // < 1 year
      return {
        id: this.id,
        name: this.name,
        score: 5,
        level: 'Low',
        justification: `Skill has not been updated in ${Math.round(diffDays)} days.`
      };
    } else {
      return {
        id: this.id,
        name: this.name,
        score: 2,
        level: 'Medium',
        justification: `Skill is potentially stale (last update over a year ago).`
      };
    }
  }
}
