package provider

type OpenAIProvider struct{}

func (p *OpenAIProvider) ID() string   { return "openai" }
func (p *OpenAIProvider) Name() string { return "ChatGPT" }
func (p *OpenAIProvider) SignatureTools() []string {
	return []string{"code_interpreter", "dalle", "browser"}
}
func (p *OpenAIProvider) DangerousTools() []string {
	return []string{"code_interpreter"}
}
