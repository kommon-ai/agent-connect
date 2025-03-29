import { fastify } from 'fastify';
import { fastifyConnectPlugin } from '@connectrpc/connect-fastify';
import { newNoopAgentFactory } from './agent/factory.js';
import { registerRemoteAgentService } from './service/remote.js';

/**
 * Main function to start the server
 */
async function main() {
  // Create a Fastify instance
  const server = fastify({
    logger: true
  });

  // Register the Connect plugin
  await server.register(fastifyConnectPlugin);
  
  // Create a NoopAgentFactory
  const factory = newNoopAgentFactory();
  
  // Get router from fastify instance
  const router = server.connect;
  if (!router) {
    throw new Error('Failed to get Connect router from Fastify instance');
  }
  
  // Register the RemoteAgentService
  registerRemoteAgentService(router, factory);
  
  // Get port from environment or use default
  const port = parseInt(process.env.PORT || '8080', 10);
  const host = process.env.HOST || '0.0.0.0';
  
  try {
    // Start the server
    await server.listen({ port, host });
    console.log(`Server listening on ${host}:${port}`);
  } catch (err) {
    server.log.error(err);
    process.exit(1);
  }
}

// Start the server
main().catch(err => {
  console.error('Failed to start server:', err);
  process.exit(1);
});