import { Context } from './context.js';

/**
 * Provider interface defines methods to access information about an AI service provider
 */
export interface Provider {
  /**
   * Get the model name to use with this provider
   */
  getModelName(): string;
  
  /**
   * Get the API key for authentication with the provider
   */
  getAPIKey(): string;
  
  /**
   * Get the name of the provider (e.g., "openai", "anthropic")
   */
  getProviderName(): string;
  
  /**
   * Get environment variables related to this provider
   */
  getEnv(): Record<string, string>;
}

/**
 * AgentEnv interface defines methods to access agent environment variables
 */
export interface AgentEnv {
  /**
   * Get all environment variables
   */
  getEnv(): Record<string, string>;
  
  /**
   * Get list of required environment variables
   */
  getRequiredEnv(): string[];
  
  /**
   * Validate that all required environment variables are set
   * @returns Error if validation fails, null if successful
   */
  validateRequiredEnv(): Error | null;
}

/**
 * Agent interface defines methods that all agent implementations must provide
 */
export interface Agent {
  /**
   * Execute a task with the provided input
   * @param ctx Context for the execution
   * @param input The instruction or input to process
   * @returns Promise resolving to the output string
   */
  execute(ctx: Context, input: string): Promise<string>;
  
  /**
   * Get the session ID associated with this agent
   */
  getSessionID(): string;
  
  /**
   * Clean up any resources used by this agent
   * @returns Promise resolving when cleanup is complete
   */
  clean(): Promise<void>;
  
  /**
   * Get the endpoint URL for this agent
   */
  getAgentEndpoint(): string;
  
  /**
   * Get the provider used by this agent
   */
  getProvider(): Provider;
  
  /**
   * Get the environment variables for this agent
   */
  getEnv(): AgentEnv;
}