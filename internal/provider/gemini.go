package provider

type GeminiProvider struct{}

func (p *GeminiProvider) ID() string   { return "google" }
func (p *GeminiProvider) Name() string { return "Gemini" }
func (p *GeminiProvider) SignatureTools() []string {
	return []string{"run_shell_command", "write_file", "google_search"}
}
func (p *GeminiProvider) DangerousTools() []string {
	return []string{"run_shell_command", "write_file", "delete_file"}
}
