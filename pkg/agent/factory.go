package agent

import (
	"github.com/kommon-ai/agent-connect/gen/proto"
	"github.com/kommon-ai/agent-go/pkg/agent"
)

type AgentFactory interface {
	NewAgentFactory() func(msg *proto.ExecuteTaskRequest) agent.Agent
}

type NoopAgentFactory struct{}

func (f *NoopAgentFactory) NewAgentFactory() func(msg *proto.ExecuteTaskRequest) agent.Agent {
	return func(msg *proto.ExecuteTaskRequest) agent.Agent {
		return &agent.NoopAgent{}
	}
}

func NewNoopAgentFactory() AgentFactory {
	return &NoopAgentFactory{}
}
