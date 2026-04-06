package provider

type GenericProvider struct{}

func (p *GenericProvider) ID() string               { return "generic" }
func (p *GenericProvider) Name() string             { return "Generic" }
func (p *GenericProvider) SignatureTools() []string { return nil }
func (p *GenericProvider) DangerousTools() []string { return nil }
