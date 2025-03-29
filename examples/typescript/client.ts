import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { RemoteAgentService } from "../../gen/proto/remote_connect.js";
import { 
  ExecuteTaskRequest, 
  ProviderInfo,
  PingRequest 
} from "../../gen/proto/remote_pb.js";

// Create a transport for the connection
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080", // Adjust to your server URL
});

// Create a client using the transport
const client = createClient(RemoteAgentService, transport);

// Example: Ping the server
async function pingExample() {
  try {
    const request = new PingRequest();
    const response = await client.ping(request);
    console.log("Ping response:", response);
  } catch (error) {
    console.error("Error during ping:", error);
  }
}

// Example: Execute a task
async function executeTaskExample() {
  try {
    const provider = new ProviderInfo({
      modelName: "gpt-4o",
      apiKey: "your-api-key",  // Replace with actual API key
      providerName: "openai",
      env: {}
    });
    
    const request = new ExecuteTaskRequest({
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