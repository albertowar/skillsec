package checks

import (
	"context"
	"regexp"
	"strings"

	"github.com/albertowar/skillauditai/internal/behavioral"
	"github.com/albertowar/skillauditai/pkg/api"
)

type SecretScanningCheck struct{}

func (c *SecretScanningCheck) ID() string      { return "secret-scanning" }
func (c *SecretScanningCheck) Name() string    { return "Secret Scanning Audit" }
func (c *SecretScanningCheck) Weight() float64 { return 1.0 }

func (c *SecretScanningCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	patterns := []string{
		`\bsk-[a-zA-Z0-9]{20,}\b`,     // OpenAI
		`\bghp_[a-zA-Z0-9]{36}\b`,     // GitHub
		`\bsq0csp-[a-zA-Z0-9]{32}\b`,  // Stripe (called Square in TS, but sq0csp is Square/Stripe? Actually sq0csp is Square)
		`\bAKIA[0-9A-Z]{16}\b`,        // AWS Access Key
	}

	allContent := strings.Join([]string{
		skill.Raw,
		skill.SystemPrompt,
		strings.Join(skill.Examples, "\n"),
	}, "\n")

	found := false
	for _, p := range patterns {
		re := regexp.MustCompile(p)
		if re.MatchString(allContent) {
			found = true
			break
		}
	}

	if found {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         0,
			Level:         api.Critical,
			Justification: "Potential secrets detected in skill content, instructions, or examples.",
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         10,
		Level:         api.Low,
		Justification: "No secrets detected.",
	}, nil
}
