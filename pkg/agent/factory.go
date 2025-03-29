package agent

import (
	"github.com/kommon-ai/agent-connect/gen/proto"
	"github.com/kommon-ai/agent-go/pkg/agent"
)

type AgentFactory interface {
	SetBeforeTaskExecutionFunc(func(msg *proto.ExecuteTaskRequest) error) error
	SetAfterTaskExecutionFunc(func(msg *proto.ExecuteTaskRequest) error) error
	GetBeforeTaskExecutionFunc() func(msg *proto.ExecuteTaskRequest) error
	GetAfterTaskExecutionFunc() func(msg *proto.ExecuteTaskRequest) error
	NewAgentFactory() func(msg *proto.ExecuteTaskRequest) (agent.Agent, error)
}

type NoopAgentFactory struct {
	beforeTaskExecutionFunc func(msg *proto.ExecuteTaskRequest) error
	afterTaskExecutionFunc  func(msg *proto.ExecuteTaskRequest) error
}

func (f *NoopAgentFactory) NewAgentFactory() func(msg *proto.ExecuteTaskRequest) (agent.Agent, error) {
	return func(msg *proto.ExecuteTaskRequest) (agent.Agent, error) {
		return &agent.NoopAgent{}, nil
	}
}

func (f *NoopAgentFactory) SetBeforeTaskExecutionFunc(hook func(msg *proto.ExecuteTaskRequest) error) error {
	f.beforeTaskExecutionFunc = hook
	return nil
}

func (f *NoopAgentFactory) SetAfterTaskExecutionFunc(hook func(msg *proto.ExecuteTaskRequest) error) error {
	f.afterTaskExecutionFunc = hook
	return nil
}

func (f *NoopAgentFactory) GetBeforeTaskExecutionFunc() func(msg *proto.ExecuteTaskRequest) error {
	return f.beforeTaskExecutionFunc
}

func (f *NoopAgentFactory) GetAfterTaskExecutionFunc() func(msg *proto.ExecuteTaskRequest) error {
	return f.afterTaskExecutionFunc
}

func NewNoopAgentFactory() AgentFactory {
	return &NoopAgentFactory{}
}
