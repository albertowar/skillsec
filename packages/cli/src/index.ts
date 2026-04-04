import { readFileSync } from 'fs';
import { parseSkill, Auditor } from '@skillauditai/core';
import chalk from 'chalk';

async function run() {
  const filePath = process.argv[2];
  if (!filePath) {
    console.error('Usage: skillaudit <path-to-skill.md>');
    process.exit(1);
  }

  const content = readFileSync(filePath, 'utf-8');
  const context = parseSkill(content);
  const auditor = new Auditor();
  const report = await auditor.audit(context);

  console.log(chalk.bold(`\nSkillAuditAI Report - Score: ${report.finalScore.toFixed(1)}/10\n`));
  report.results.forEach(r => {
    const color = r.score > 7 ? chalk.green : r.score > 4 ? chalk.yellow : chalk.red;
    console.log(`${color(`[${r.level}]`)} ${chalk.bold(r.name)}: ${r.score}/10`);
    console.log(`Justification: ${r.justification}\n`);
  });
}

run().catch(console.error);
