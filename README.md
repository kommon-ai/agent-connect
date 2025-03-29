# agent-connect

A remote agent service using Buf Connect for RPC communication.

## Overview

This project provides a remote agent service that can execute tasks via a Connect RPC interface. It uses Protocol Buffers for service definition and Buf Connect for the RPC implementation.

## Setup

### Prerequisites

- Go 1.19 or newer
- [Buf CLI](https://buf.build/product/cli/)
- Node.js (for TypeScript client usage)

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/takutakahashi/agent-connect.git
   cd agent-connect
   ```

2. Install Go dependencies:
   ```
   go mod tidy
   ```

3. Generate code from Proto definitions:
   ```
   buf generate
   ```

4. For TypeScript client development, set up the Node.js environment:
   ```
   make setup-ts
   ```

## Usage

### Running the Server

```
go run cmd/server/main.go
```

This will start a server on port 8080.

### Running the Go Client

```
go run cmd/client/main.go
```

### Running the TypeScript Client

```
make ts-client
```

This will:
1. Ensure Node.js and npm are installed via mise
2. Install TypeScript dependencies
3. Build the TypeScript code
4. Run the sample TypeScript client

## API

The service provides the following RPCs:

- `Ping`: A simple health check endpoint
- `ExecuteTask`: Executes a task with the specified configuration

See `proto/remote.proto` for full details on the service definition and message formats.

## TypeScript Client Library

The TypeScript client library is automatically generated from the Protocol Buffer definitions using buf.

### Installation and Setup

```bash
make setup-ts
```

### Using the Client in Your TypeScript Project

```typescript
import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { RemoteAgentService } from "agent-connect-client/gen/proto";
import { 
  ExecuteTaskRequestSchema, 
  ProviderInfoSchema 
} from "agent-connect-client/gen/proto";
import { create } from "@bufbuild/protobuf";

// Create a transport
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Create a client
const client = createPromiseClient(RemoteAgentService, transport);

// Create a request
const request = create(ExecuteTaskRequestSchema, {
  sessionId: "example-session",
  provider: create(ProviderInfoSchema, {
    modelName: "gpt-4o",
    apiKey: "your-api-key",
    providerName: "openai"
  }),
  instruction: "Hello from TypeScript!"
});

// Call the service
async function executeTask() {
  try {
    const response = await client.executeTask(request);
    console.log("Response:", response);
  } catch (error) {
    console.error("Error:", error);
  }
}

executeTask();
```

See the `examples/typescript` directory for more detailed examples.

## Development

### Regenerating Code after Proto Changes

If you make changes to the proto files, regenerate the code with:

```
buf generate
```

This will update both Go and TypeScript code in the `gen` directory.
