import { execSync } from 'child_process';
import { SkillMetadata } from '@skillauditai/core';

export function getGitMetadata(filePath: string): SkillMetadata | undefined {
  try {
    // Check if in git repo and get info
    const logOutput = execSync(`git log -1 --format="%an|%ae|%at|%G?" -- "${filePath}"`, { 
      encoding: 'utf-8',
      stdio: ['pipe', 'pipe', 'ignore']
    }).trim();

    if (!logOutput) return undefined;

    const [name, email, timestamp, signature] = logOutput.split('|');

    return {
      author: {
        name,
        email,
        isVerified: signature === 'G'
      },
      maintenance: {
        lastUpdated: new Date(parseInt(timestamp) * 1000).toISOString()
      }
    };
  } catch (error) {
    // Not a git repo or file not tracked
    return undefined;
  }
}
