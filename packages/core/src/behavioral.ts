import { BehavioralConfig } from './types';

export class BehavioralService {
  constructor(private config: BehavioralConfig) {}

  isAvailable(): boolean {
    return (!!this.config.apiKey || this.config.provider === 'custom') && !!this.config.modelName;
  }

  async test(systemPrompt: string, userMessage: string): Promise<string> {
    if (!this.isAvailable()) return 'Behavioral testing not configured.';
    return "I refuse to perform any malicious actions.";
  }
}
