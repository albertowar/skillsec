package provider

type AnthropicProvider struct{}

func (p *AnthropicProvider) ID() string   { return "anthropic" }
func (p *AnthropicProvider) Name() string { return "Claude" }
func (p *AnthropicProvider) SignatureTools() []string {
	return []string{"computer", "bash", "text_editor"}
}
func (p *AnthropicProvider) DangerousTools() []string {
	return []string{"computer", "bash", "text_editor"}
}
