package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/kommon-ai/agent-connect/gen/proto"
	"github.com/kommon-ai/agent-connect/gen/proto/protoconnect"
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

	// GitHubInfo付きのExecuteTaskリクエスト
	executeReqWithGitHub := &proto.ExecuteTaskRequest{
		SessionId: "test-session-1",
		Provider: &proto.ProviderInfo{
			ModelName:    "gpt-4o",
			ApiKey:       "dummy-key",
			ProviderName: "openai",
		},
		Github: &proto.GitHubInfo{
			ApiToken:    "github-token",
			Repo:        "kommon-ai/agent-connect",
			BranchName:  "main",
		},
		Instruction: "Hello with GitHub info!",
	}

	executeRespWithGitHub, err := client.ExecuteTask(ctx, connect.NewRequest(executeReqWithGitHub))
	if err != nil {
		log.Fatalf("ExecuteTask with GitHub error: %v", err)
	}

	log.Printf("ExecuteTask with GitHub response: Success=%v", executeRespWithGitHub.Msg.Success)
	log.Printf("ExecuteTask with GitHub stdout: %s", executeRespWithGitHub.Msg.Stdout)
	if executeRespWithGitHub.Msg.Stderr != "" {
		log.Printf("ExecuteTask with GitHub stderr: %s", executeRespWithGitHub.Msg.Stderr)
	}

	// GitHubInfoなしのExecuteTaskリクエスト
	executeReqWithoutGitHub := &proto.ExecuteTaskRequest{
		SessionId: "test-session-2",
		Provider: &proto.ProviderInfo{
			ModelName:    "gpt-4o",
			ApiKey:       "dummy-key",
			ProviderName: "openai",
		},
		Instruction: "Hello without GitHub info!",
	}

	executeRespWithoutGitHub, err := client.ExecuteTask(ctx, connect.NewRequest(executeReqWithoutGitHub))
	if err != nil {
		log.Fatalf("ExecuteTask without GitHub error: %v", err)
	}

	log.Printf("ExecuteTask without GitHub response: Success=%v", executeRespWithoutGitHub.Msg.Success)
	log.Printf("ExecuteTask without GitHub stdout: %s", executeRespWithoutGitHub.Msg.Stdout)
	if executeRespWithoutGitHub.Msg.Stderr != "" {
		log.Printf("ExecuteTask without GitHub stderr: %s", executeRespWithoutGitHub.Msg.Stderr)
	}
}