import { readFileSync } from 'fs';
import { parseSkill, Auditor, BehavioralConfig } from '@skillauditai/core';
import chalk from 'chalk';
import Table from 'cli-table3';

export function parseFlags(args: string[]) {
  const formatIndex = args.indexOf('--format');
  const format = formatIndex !== -1 ? args[formatIndex + 1] : 'table';
  
  const apiKeyIndex = args.indexOf('--api-key');
  const apiKey = apiKeyIndex !== -1 ? args[apiKeyIndex + 1] : undefined;
  
  const modelIndex = args.indexOf('--model');
  const modelName = modelIndex !== -1 ? args[modelIndex + 1] : undefined;
  
  const providerIndex = args.indexOf('--provider');
  const provider = providerIndex !== -1 ? args[providerIndex + 1] : undefined;
  
  const baseUrlIndex = args.indexOf('--base-url');
  const baseUrl = baseUrlIndex !== -1 ? args[baseUrlIndex + 1] : undefined;

  const filePath = args.find((arg, index) => 
    !arg.startsWith('--') && 
    (index === 0 || !['--format', '--api-key', '--model', '--provider', '--base-url'].includes(args[index-1]))
  );

  return {
    format,
    filePath,
    config: {
      apiKey,
      modelName,
      provider: provider as any,
      baseUrl
    }
  };
}

async function run() {
  try {
    const { format, filePath, config } = parseFlags(process.argv.slice(2));

    if (!filePath) {
      console.error('Usage: skillaudit <path-to-skill.md> [--format <json|table>] [--api-key <key>] [--model <model>] [--provider <provider>] [--base-url <url>]');
      process.exit(1);
    }

    const content = readFileSync(filePath, 'utf-8');
    const context = parseSkill(content);
    const auditor = new Auditor([], config);
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

if (process.env.NODE_ENV !== 'test') {
  run();
}
