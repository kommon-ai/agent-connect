import type { ConnectRouter } from '@connectrpc/connect';
import type { 
  ExecuteTaskRequest, 
  ExecuteTaskResponse, 
  ExecuteTaskResponseSchema,
  PingRequest, 
  PingResponse,
  PingResponseSchema
} from '@gen/remote_pb.js';
import { RemoteAgentService } from '@gen/remote_pb.js';
import { AgentFactory } from '../agent/factory.js';
import { Context } from '../agent/context.js';

/**
 * RemoteAgentServer implements the RemoteAgentService
 */
export class RemoteAgentServer {
  private factory: AgentFactory;

  /**
   * Create a new RemoteAgentServer
   * @param factory The agent factory to use for creating agents
   */
  constructor(factory: AgentFactory) {
    this.factory = factory;
  }

  /**
   * Execute a task asynchronously
   * @param request The task execution request
   * @returns Response indicating the task was enqueued
   */
  async executeTask(request: ExecuteTaskRequest): Promise<ExecuteTaskResponse> {
    try {
      const agentFactory = this.factory.newAgentFactory();
      const agent = await agentFactory(request);
      
      // Execute the task in the background
      setTimeout(async () => {
        try {
          console.log(`Starting execution for session: ${request.sessionId}`);
          
          // Call the before hook
          const beforeHook = this.factory.getBeforeTaskExecutionFunc();
          if (beforeHook) {
            await beforeHook(request);
          }
          
          // Execute the task
          const ctx = new Context();
          await agent.execute(ctx, request.instruction);
          
          // Call the after hook
          const afterHook = this.factory.getAfterTaskExecutionFunc();
          if (afterHook) {
            await afterHook(request);
          }
          
          console.log(`Completed execution for session: ${request.sessionId}`);
        } catch (error) {
          console.error('Error in background execution:', error);
        }
      }, 0);

      // Return immediate success response
      return {
        sessionId: request.sessionId,
        stdout: 'Task enqueued successfully',
        stderr: '',
        success: true
      } as ExecuteTaskResponse;
    } catch (error) {
      console.error('Error setting up task execution:', error);
      const errorMessage = error instanceof Error ? error.message : String(error);
      
      return {
        sessionId: request.sessionId,
        stdout: '',
        stderr: `Error: ${errorMessage}`,
        success: false
      } as ExecuteTaskResponse;
    }
  }

  /**
   * Simple ping method to check if the server is running
   * @param request Empty ping request
   * @returns Response with OK status
   */
  async ping(request: PingRequest): Promise<PingResponse> {
    console.log('Ping received');
    return {
      status: 'OK'
    } as PingResponse;
  }
}

/**
 * Register the RemoteAgentService with a ConnectRouter
 * @param router The ConnectRouter to register with
 * @param factory The AgentFactory to use
 * @returns The path where the service is registered
 */
export function registerRemoteAgentService(router: ConnectRouter, factory: AgentFactory): string {
  const server = new RemoteAgentServer(factory);
  
  router.service(RemoteAgentService, {
    executeTask: server.executeTask.bind(server),
    ping: server.ping.bind(server)
  });
  
  return '/remote.RemoteAgentService';
}