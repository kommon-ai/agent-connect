package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/takutakahashi/agent-connect/gen/proto"
	"github.com/takutakahashi/agent-connect/gen/proto/protoconnect"
)

func main() {
	// クライアントの作成
	client := protoconnect.NewRemoteAgentServiceClient(
		http.DefaultClient,
		"http://localhost:8080", // サーバーのアドレス
	)

	// タイムアウト付きのコンテキスト作成
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pingリクエスト
	pingResp, err := client.Ping(ctx, connect.NewRequest(&proto.PingRequest{}))
	if err != nil {
		log.Fatalf("Ping error: %v", err)
	}
	log.Printf("Ping response: %v", pingResp.Msg.Status)

	// ExecuteTaskリクエスト
	executeReq := &proto.ExecuteTaskRequest{
		SessionId: "test-session-1",
		Provider: &proto.ProviderInfo{
			ModelName:    "gpt-4o",
			ApiKey:       "dummy-key",
			ProviderName: "openai",
		},
		Github: &proto.GitHubInfo{
			ApiToken:    "github-token",
			Repo:        "takutakahashi/agent-connect",
			BranchName:  "main",
		},
		Instruction: "Hello, execute this task!",
	}

	executeResp, err := client.ExecuteTask(ctx, connect.NewRequest(executeReq))
	if err != nil {
		log.Fatalf("ExecuteTask error: %v", err)
	}

	log.Printf("ExecuteTask response: Success=%v", executeResp.Msg.Success)
	log.Printf("ExecuteTask stdout: %s", executeResp.Msg.Stdout)
	if executeResp.Msg.Stderr != "" {
		log.Printf("ExecuteTask stderr: %s", executeResp.Msg.Stderr)
	}
}