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

## Usage

### Running the Server

```
go run cmd/server/main.go
```

This will start a server on port 8080.

### Running the Client

```
go run cmd/client/main.go
```

This sample client will:
1. Send a Ping request to check server availability
2. Execute a sample task with the provided configuration

## API

The service provides the following RPCs:

- `Ping`: A simple health check endpoint
- `ExecuteTask`: Executes a task with the specified configuration

See `proto/remote.proto` for full details on the service definition and message formats.

## Development

### Regenerating Code after Proto Changes

If you make changes to the proto files, regenerate the code with:

```
buf generate
```

This will update both Go and TypeScript code in the `gen` directory.
