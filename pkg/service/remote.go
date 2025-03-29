package service

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	remote "github.com/kommon-ai/agent-connect/gen/proto"
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
	req *connect.Request[remote.ExecuteTaskRequest],
) (*connect.Response[remote.ExecuteTaskResponse], error) {
	agent, err := s.factory.NewAgentFactory()(req.Msg)
	if err != nil {
		return nil, err
	}

	longCtx := context.Background()
	go func() {
		if err := s.factory.GetBeforeTaskExecutionFunc()(req.Msg); err != nil {
			log.Printf("Error executing before hook: %v", err)
		}
		if _, err := agent.Execute(longCtx, req.Msg.Instruction); err != nil {
			log.Printf("Error executing task: %v", err)
		}
		if err := s.factory.GetAfterTaskExecutionFunc()(req.Msg); err != nil {
			log.Printf("Error executing after hook: %v", err)
		}
	}()

	resp := &remote.ExecuteTaskResponse{
		Stdout: "Task enqueued successfully",
		Stderr: "",
	}

	return connect.NewResponse(resp), nil
}

// Ping はサーバーの状態を確認するメソッドです
func (s *RemoteAgentServer) Ping(
	ctx context.Context,
	req *connect.Request[remote.PingRequest],
) (*connect.Response[remote.PingResponse], error) {
	log.Println("Ping received")

	res := &remote.PingResponse{
		Status: "OK",
	}

	return connect.NewResponse(res), nil
}

// Handler はRemoteAgentServiceのHTTPハンドラを返します
func (s *RemoteAgentServer) Handler() (string, http.Handler) {
	return protoconnect.NewRemoteAgentServiceHandler(s)
}
