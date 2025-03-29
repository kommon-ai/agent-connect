import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { RemoteAgentService } from "../../gen/proto/remote_connect";
import { 
  ExecuteTaskRequestSchema, 
  ProviderInfoSchema,
  PingRequestSchema 
} from "../../gen/proto/remote_pb";
import { create } from "@bufbuild/protobuf";

// Create a transport for the connection
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080", // Adjust to your server URL
});

// Create a client using the transport
const client = createPromiseClient(RemoteAgentService, transport);

// Example: Ping the server
async function pingExample() {
  try {
    const request = create(PingRequestSchema, {});
    const response = await client.ping(request);
    console.log("Ping response:", response);
  } catch (error) {
    console.error("Error during ping:", error);
  }
}

// Example: Execute a task
async function executeTaskExample() {
  try {
    const provider = create(ProviderInfoSchema, {
      modelName: "gpt-4o",
      apiKey: "your-api-key",  // Replace with actual API key
      providerName: "openai",
      env: {}
    });
    
    const request = create(ExecuteTaskRequestSchema, {
      sessionId: "example-session",
      provider: provider,
      instruction: "Hello from TypeScript client!"
    });
    
    const response = await client.executeTask(request);
    console.log("Execute task response:", response);
  } catch (error) {
    console.error("Error during task execution:", error);
  }
}

// Run the examples
async function main() {
  console.log("Running ping example:");
  await pingExample();
  
  console.log("\nRunning execute task example:");
  await executeTaskExample();
}

main().catch(console.error);