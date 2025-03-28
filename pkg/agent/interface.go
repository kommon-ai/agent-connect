package agent

import "context"

// AgentEnv defines an interface for agent environment variables
type AgentEnv interface {
	// GetEnv returns a map of environment variables
	GetEnv() map[string]string
	// GetRequiredEnv returns a list of required environment variables
	GetRequiredEnv() []string
	// ValidateRequiredEnv validates that all required environment variables are set
	ValidateRequiredEnv() error
}

// Provider defines an interface for a service that provides AI capabilities
type Provider interface {
	GetModelName() string
	GetAPIKey() string
	GetProviderName() string
	GetEnv() map[string]string
}

type GitHub interface {
	GetAPIToken() string
	GetAPIURL() string
	GetRepo() string
	GetFullRepoURL() string
	GetPRNumber() (int, error)
	GetIssueNumber() (int, error)
	GetBranchName() string
}

// Agent represents a general AI agent interface
type Agent interface {
	Execute(ctx context.Context, input string) (string, error)
	GetSessionID() string
	Clean() error
	GetAgentEndpoint() string
	// New methods
	GetProvider() Provider
	GetEnv() AgentEnv
}

type AgentFactory interface {
	NewAgent(sessionID string) Agent
}
