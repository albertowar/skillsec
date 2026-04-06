package provider_test

import (
	"testing"
	"github.com/albertowar/skillsec/internal/provider"
)

type MockProvider struct{}

func (m *MockProvider) ID() string               { return "mock" }
func (m *MockProvider) Name() string             { return "Mock Provider" }
func (m *MockProvider) SignatureTools() []string { return []string{"tool1"} }
func (m *MockProvider) DangerousTools() []string { return []string{"danger1"} }

func TestProviderInterface(t *testing.T) {
	var _ provider.Provider = (*MockProvider)(nil)
}
