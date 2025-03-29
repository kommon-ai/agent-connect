import { Agent, AgentEnv, Provider } from './interfaces.js';
import { Context } from './context.js';

/**
 * NoopProvider is a Provider implementation that does nothing
 */
export class NoopProvider implements Provider {
  private modelName: string = '';
  private apiKey: string = '';
  private providerName: string = '';
  private env: Record<string, string> = {};
  
  constructor(
    modelName: string = '',
    apiKey: string = '',
    providerName: string = '',
    env: Record<string, string> = {}
  ) {
    this.modelName = modelName;
    this.apiKey = apiKey;
    this.providerName = providerName;
    this.env = env;
  }
  
  getModelName(): string { 
    return this.modelName;
  }
  
  getAPIKey(): string { 
    return this.apiKey;
  }
  
  getProviderName(): string { 
    return this.providerName;
  }
  
  getEnv(): Record<string, string> { 
    return this.env;
  }
}

/**
 * NoopAgentEnv is an AgentEnv implementation that does nothing
 */
export class NoopAgentEnv implements AgentEnv {
  private env: Record<string, string> = {};
  private requiredEnv: string[] = [];
  
  constructor(env: Record<string, string> = {}, requiredEnv: string[] = []) {
    this.env = env;
    this.requiredEnv = requiredEnv;
  }
  
  getEnv(): Record<string, string> { 
    return this.env;
  }
  
  getRequiredEnv(): string[] { 
    return this.requiredEnv;
  }
  
  validateRequiredEnv(): Error | null { 
    // Simple validation - check that all required environment variables exist
    for (const key of this.requiredEnv) {
      if (!(key in this.env)) {
        return new Error(`Required environment variable missing: ${key}`);
      }
    }
    return null;
  }
}

/**
 * NoopAgent is an Agent implementation that does nothing substantial
 * It just returns a simple response without actual AI processing
 */
export class NoopAgent implements Agent {
  private sessionId: string = '';
  private provider: Provider;
  private agentEnv: AgentEnv;
  
  constructor(
    sessionId: string = '', 
    provider: Provider = new NoopProvider(),
    agentEnv: AgentEnv = new NoopAgentEnv()
  ) {
    this.sessionId = sessionId;
    this.provider = provider;
    this.agentEnv = agentEnv;
  }
  
  async execute(ctx: Context, input: string): Promise<string> {
    if (ctx.isCancelled()) {
      throw new Error('Context cancelled before execution started');
    }
    
    return `hello, input: ${input}`;
  }
  
  getSessionID(): string {
    return this.sessionId;
  }
  
  async clean(): Promise<void> {
    // Nothing to clean up
  }
  
  getAgentEndpoint(): string {
    return '';
  }
  
  getProvider(): Provider {
    return this.provider;
  }
  
  getEnv(): AgentEnv {
    return this.agentEnv;
  }
}