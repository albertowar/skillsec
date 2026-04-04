import { readFileSync } from 'fs';
import { parseSkill, Auditor } from '@skillauditai/core';
import chalk from 'chalk';
import Table from 'cli-table3';

async function run() {
  try {
    const args = process.argv.slice(2);
    const formatIndex = args.indexOf('--format');
    const format = formatIndex !== -1 ? args[formatIndex + 1] : 'table';
    const filePath = args.find((arg, index) => !arg.startsWith('--') && (index === 0 || args[index-1] !== '--format'));

    if (!filePath) {
      console.error('Usage: skillaudit <path-to-skill.md> [--format <json|table>]');
      process.exit(1);
    }

    const content = readFileSync(filePath, 'utf-8');
    const context = parseSkill(content);
    const auditor = new Auditor();
    const report = await auditor.audit(context);

    if (format === 'json') {
      console.log(JSON.stringify(report, null, 2));
    } else {
      console.log(chalk.bold(`\nSkillAuditAI Report - Score: ${report.finalScore.toFixed(1)}/10\n`));
      
      const table = new Table({
        head: [chalk.bold('Level'), chalk.bold('Check'), chalk.bold('Score'), chalk.bold('Justification')],
        colWidths: [15, 25, 10, 50],
        wordWrap: true
      });

      report.results.forEach(r => {
        const color = r.score > 7 ? chalk.green : r.score > 4 ? chalk.yellow : chalk.red;
        table.push([
          color(r.level),
          chalk.bold(r.name),
          `${r.score}/10`,
          r.justification
        ]);
      });

      console.log(table.toString());
      console.log('\n');
    }
  } catch (error) {
    console.error(chalk.red('\nError during audit:'), error instanceof Error ? error.message : String(error));
    process.exit(1);
  }
}

run();
