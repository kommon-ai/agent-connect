package agent

import (
	remote "github.com/kommon-ai/agent-connect/gen/proto"
	"github.com/kommon-ai/agent-go/pkg/agent"
)

type AgentFactory interface {
	SetBeforeTaskExecutionFunc(func(msg *remote.ExecuteTaskRequest) error) error
	SetAfterTaskExecutionFunc(func(msg *remote.ExecuteTaskRequest) error) error
	GetBeforeTaskExecutionFunc() func(msg *remote.ExecuteTaskRequest) error
	GetAfterTaskExecutionFunc() func(msg *remote.ExecuteTaskRequest) error
	NewAgentFactory() func(msg *remote.ExecuteTaskRequest) (agent.Agent, error)
}

type NoopAgentFactory struct {
	beforeTaskExecutionFunc func(msg *remote.ExecuteTaskRequest) error
	afterTaskExecutionFunc  func(msg *remote.ExecuteTaskRequest) error
}

func (f *NoopAgentFactory) NewAgentFactory() func(msg *remote.ExecuteTaskRequest) (agent.Agent, error) {
	return func(msg *remote.ExecuteTaskRequest) (agent.Agent, error) {
		return &agent.NoopAgent{}, nil
	}
}

func (f *NoopAgentFactory) SetBeforeTaskExecutionFunc(hook func(msg *remote.ExecuteTaskRequest) error) error {
	f.beforeTaskExecutionFunc = hook
	return nil
}

func (f *NoopAgentFactory) SetAfterTaskExecutionFunc(hook func(msg *remote.ExecuteTaskRequest) error) error {
	f.afterTaskExecutionFunc = hook
	return nil
}

func (f *NoopAgentFactory) GetBeforeTaskExecutionFunc() func(msg *remote.ExecuteTaskRequest) error {
	return f.beforeTaskExecutionFunc
}

func (f *NoopAgentFactory) GetAfterTaskExecutionFunc() func(msg *remote.ExecuteTaskRequest) error {
	return f.afterTaskExecutionFunc
}

func NewNoopAgentFactory() AgentFactory {
	return &NoopAgentFactory{}
}
