package main

import (
	"log"
	"net/http"

	"github.com/kommon-ai/agent-connect/pkg/service"
	"github.com/kommon-ai/agent-go/pkg/agent"
)

func main() {
	// RemoteAgentServerの作成
	remoteAgent := service.NewRemoteAgentServer(agent.NewNoopAgentFactory())

	// ハンドラの作成
	mux := http.NewServeMux()

	// RemoteAgentServiceハンドラの登録
	path, handler := remoteAgent.Handler()
	mux.Handle(path, handler)

	// サーバーの起動
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
