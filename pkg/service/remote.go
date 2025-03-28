package service

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/kommon-ai/agent-connect/gen/proto"
	"github.com/kommon-ai/agent-connect/gen/proto/protoconnect"
	"github.com/kommon-ai/agent-connect/pkg/agent"
)

// RemoteAgentServer はRemoteAgentServiceの実装です
type RemoteAgentServer struct {
	factory agent.AgentFactory
}

// NewRemoteAgentServer は新しいRemoteAgentServerを作成します
func NewRemoteAgentServer(factory agent.AgentFactory) *RemoteAgentServer {
	return &RemoteAgentServer{factory: factory}
}

// ExecuteTask はタスクを実行するメソッドです
func (s *RemoteAgentServer) ExecuteTask(
	ctx context.Context,
	req *connect.Request[proto.ExecuteTaskRequest],
) (*connect.Response[proto.ExecuteTaskResponse], error) {
	slog.Info("ExecuteTask", "request", req.Msg)

	agent, err := s.factory.NewAgentFactory()(req.Msg)
	if err != nil {
		return nil, err
	}

	longCtx := context.Background()
	go func() {
		if _, err := agent.Execute(longCtx, req.Msg.Instruction); err != nil {
			slog.Error("Error executing task", "error", err)
		}
	}()

	resp := &proto.ExecuteTaskResponse{
		Stdout: "Task enqueued successfully",
		Stderr: "",
	}

	return connect.NewResponse(resp), nil
}

// Ping はサーバーの状態を確認するメソッドです
func (s *RemoteAgentServer) Ping(
	ctx context.Context,
	req *connect.Request[proto.PingRequest],
) (*connect.Response[proto.PingResponse], error) {
	slog.Info("Ping received")

	res := &proto.PingResponse{
		Status: "OK",
	}

	return connect.NewResponse(res), nil
}

// Handler はRemoteAgentServiceのHTTPハンドラを返します
func (s *RemoteAgentServer) Handler() (string, http.Handler) {
	return protoconnect.NewRemoteAgentServiceHandler(s)
}
