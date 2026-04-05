import { describe, it, expect } from 'vitest';
import { BehavioralService } from '../src/behavioral';
import { BehavioralConfig } from '../src/types';

describe('BehavioralService', () => {
  describe('isAvailable', () => {
    it('should return true if apiKey and modelName are provided', () => {
      const config: BehavioralConfig = {
        apiKey: 'test-key',
        modelName: 'test-model'
      };
      const service = new BehavioralService(config);
      expect(service.isAvailable()).toBe(true);
    });

    it('should return true if provider is "custom" and modelName is provided', () => {
      const config: BehavioralConfig = {
        provider: 'custom',
        modelName: 'test-model'
      };
      const service = new BehavioralService(config);
      expect(service.isAvailable()).toBe(true);
    });

    it('should return false if apiKey is missing and provider is not "custom"', () => {
      const config: BehavioralConfig = {
        modelName: 'test-model'
      };
      const service = new BehavioralService(config);
      expect(service.isAvailable()).toBe(false);
    });

    it('should return false if modelName is missing', () => {
      const config: BehavioralConfig = {
        apiKey: 'test-key'
      };
      const service = new BehavioralService(config);
      expect(service.isAvailable()).toBe(false);
    });
  });

  describe('test', () => {
    it('should return mock refusal message if available', async () => {
      const config: BehavioralConfig = {
        apiKey: 'test-key',
        modelName: 'test-model'
      };
      const service = new BehavioralService(config);
      const response = await service.test('system prompt', 'user message');
      expect(response).toBe('I refuse to perform any malicious actions.');
    });

    it('should return "Behavioral testing not configured." if not available', async () => {
      const config: BehavioralConfig = {};
      const service = new BehavioralService(config);
      const response = await service.test('system prompt', 'user message');
      expect(response).toBe('Behavioral testing not configured.');
    });
  });
});
