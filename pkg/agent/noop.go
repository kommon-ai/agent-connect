package agent

import "context"

// NoopProvider implements a no-operation Provider
type NoopProvider struct {
	ModelName    string
	APIKey       string
	ProviderName string
	Env          map[string]string
}

func (p *NoopProvider) GetModelName() string    { return p.ModelName }
func (p *NoopProvider) GetAPIKey() string       { return p.APIKey }
func (p *NoopProvider) GetProviderName() string { return p.ProviderName }
func (p *NoopProvider) GetEnv() map[string]string {
	return p.Env
}

type NoopEnv struct {
	Env          map[string]string
	RequiredEnv  []string
	ValidateFunc func() error
}

func (e *NoopEnv) GetEnv() map[string]string {
	return e.Env
}

func (e *NoopEnv) GetRequiredEnv() []string {
	return e.RequiredEnv
}

func (e *NoopEnv) ValidateRequiredEnv() error {
	return e.ValidateFunc()
}

// NoopAgent implements a no-operation Agent
type NoopAgent struct {
	Provider Provider
	Env      AgentEnv
}

func (a *NoopAgent) Execute(ctx context.Context, input string) (string, error) {
	return "hello, input: " + input, nil
}

func (a *NoopAgent) GetSessionID() string {
	return ""
}

func (a *NoopAgent) Clean() error {
	return nil
}

func (a *NoopAgent) GetProvider() Provider {
	return a.Provider
}

func (a *NoopAgent) GetEnv() AgentEnv {
	return a.Env
}

func (a *NoopAgent) GetAgentEndpoint() string {
	return ""
}

type NoopAgentFactory struct{}

func (f *NoopAgentFactory) NewAgent(sessionID string) Agent {
	return &NoopAgent{
		Provider: &NoopProvider{},
		Env:      &NoopEnv{},
	}
}

func NewNoopAgentFactory() AgentFactory {
	return &NoopAgentFactory{}
}
