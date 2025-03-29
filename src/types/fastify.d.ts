import { ConnectRouter } from '@connectrpc/connect';
import { FastifyInstance } from 'fastify';

declare module 'fastify' {
  interface FastifyInstance {
    connect: ConnectRouter;
  }
}