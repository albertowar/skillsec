package provider

type Provider interface {
	ID() string
	Name() string
	SignatureTools() []string
	DangerousTools() []string
}
