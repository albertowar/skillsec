package provider

import "strings"

var registry = []Provider{
	&GeminiProvider{},
	&OpenAIProvider{},
	&AnthropicProvider{},
}

func Get(id string) Provider {
	for _, p := range registry {
		if p.ID() == id {
			return p
		}
	}
	return &GenericProvider{}
}

func Detect(tools []string) string {
	for _, t := range tools {
		lowerT := strings.ToLower(t)
		for _, p := range registry {
			for _, sig := range p.SignatureTools() {
				if lowerT == sig {
					return p.ID()
				}
			}
		}
	}
	return "generic"
}
