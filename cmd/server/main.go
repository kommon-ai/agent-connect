package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/takutakahashi/agent-connect/gen/proto"
	"github.com/takutakahashi/agent-connect/gen/proto/protoconnect"
)

// server はRemoteAgentServiceの実装です
type server struct{}

// ExecuteTask はタスクを実行するメソッドです
func (s *server) ExecuteTask(
	ctx context.Context,
	req *connect.Request[proto.ExecuteTaskRequest],
) (*connect.Response[proto.ExecuteTaskResponse], error) {
	log.Printf("ExecuteTask: %v", req.Msg)
	
	// リクエストからsession_idを取得
	sessionID := req.Msg.SessionId
	
	// レスポンスの作成
	res := &proto.ExecuteTaskResponse{
		SessionId: sessionID,
		Stdout:    fmt.Sprintf("Executed task for session: %s", sessionID),
		Stderr:    "",
		Success:   true,
	}
	
	return connect.NewResponse(res), nil
}

// Ping はサーバーの状態を確認するメソッドです
func (s *server) Ping(
	ctx context.Context,
	req *connect.Request[proto.PingRequest],
) (*connect.Response[proto.PingResponse], error) {
	log.Println("Ping received")
	
	res := &proto.PingResponse{
		Status: "OK",
	}
	
	return connect.NewResponse(res), nil
}

func main() {
	// サーバーの実装
	remoteAgent := &server{}
	
	// ハンドラの作成
	mux := http.NewServeMux()
	
	// RemoteAgentServiceハンドラの登録
	path, handler := protoconnect.NewRemoteAgentServiceHandler(remoteAgent)
	mux.Handle(path, handler)
	
	// サーバーの起動
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}