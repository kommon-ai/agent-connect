import { ExecuteTaskRequest } from '../../gen/proto/remote_pb.js';
import { Agent } from './interfaces.js';
import { NoopAgent } from './noop.js';

/**
 * AgentFactory interface defines methods for creating agents and handling execution hooks
 */
export interface AgentFactory {
  /**
   * Set a function that runs before task execution
   * @param hook Function to run before task execution
   */
  setBeforeTaskExecutionFunc(hook: (msg: ExecuteTaskRequest) => Promise<void>): void;
  
  /**
   * Set a function that runs after task execution
   * @param hook Function to run after task execution
   */
  setAfterTaskExecutionFunc(hook: (msg: ExecuteTaskRequest) => Promise<void>): void;
  
  /**
   * Get the function that runs before task execution
   */
  getBeforeTaskExecutionFunc(): (msg: ExecuteTaskRequest) => Promise<void>;
  
  /**
   * Get the function that runs after task execution
   */
  getAfterTaskExecutionFunc(): (msg: ExecuteTaskRequest) => Promise<void>;
  
  /**
   * Return a factory function that creates agents based on request parameters
   */
  newAgentFactory(): (msg: ExecuteTaskRequest) => Promise<Agent>;
}

/**
 * NoopAgentFactory is an implementation of AgentFactory that creates NoopAgent instances
 */
export class NoopAgentFactory implements AgentFactory {
  private beforeTaskExecutionFunc: (msg: ExecuteTaskRequest) => Promise<void>;
  private afterTaskExecutionFunc: (msg: ExecuteTaskRequest) => Promise<void>;
  
  constructor() {
    // Default no-op hooks
    this.beforeTaskExecutionFunc = async () => {};
    this.afterTaskExecutionFunc = async () => {};
  }
  
  setBeforeTaskExecutionFunc(hook: (msg: ExecuteTaskRequest) => Promise<void>): void {
    this.beforeTaskExecutionFunc = hook;
  }
  
  setAfterTaskExecutionFunc(hook: (msg: ExecuteTaskRequest) => Promise<void>): void {
    this.afterTaskExecutionFunc = hook;
  }
  
  getBeforeTaskExecutionFunc(): (msg: ExecuteTaskRequest) => Promise<void> {
    return this.beforeTaskExecutionFunc;
  }
  
  getAfterTaskExecutionFunc(): (msg: ExecuteTaskRequest) => Promise<void> {
    return this.afterTaskExecutionFunc;
  }
  
  newAgentFactory(): (msg: ExecuteTaskRequest) => Promise<Agent> {
    return async (msg: ExecuteTaskRequest): Promise<Agent> => {
      // Create a NoopAgent with the session ID from the request
      return new NoopAgent(msg.sessionId);
    };
  }
}

/**
 * Create a new NoopAgentFactory instance
 * @returns A new NoopAgentFactory
 */
export function newNoopAgentFactory(): AgentFactory {
  return new NoopAgentFactory();
}