package agent

import (
	"github.com/kommon-ai/agent-connect/gen/proto"
	"github.com/kommon-ai/agent-go/pkg/agent"
)

type AgentFactory interface {
	NewAgentFactory() func(msg *proto.ExecuteTaskRequest) (agent.Agent, error)
}

type NoopAgentFactory struct{}

func (f *NoopAgentFactory) NewAgentFactory() func(msg *proto.ExecuteTaskRequest) (agent.Agent, error) {
	return func(msg *proto.ExecuteTaskRequest) (agent.Agent, error) {
		return &agent.NoopAgent{}, nil
	}
}

func NewNoopAgentFactory() AgentFactory {
	return &NoopAgentFactory{}
}
