package checks

import (
	"context"
	"strings"

	"github.com/albertowar/skillsec/internal/behavioral"
	"github.com/albertowar/skillsec/pkg/api"
)

type PromptInjectionCheck struct{}

func (c *PromptInjectionCheck) ID() string      { return "prompt-injection" }
func (c *PromptInjectionCheck) Name() string    { return "Prompt Injection Audit" }
func (c *PromptInjectionCheck) Weight() float64 { return 0.8 }

func (c *PromptInjectionCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	prompt := strings.ToLower(skill.SystemPrompt)

	delimiters := []string{"\"\"\"", "---", "<context>", "###", "```"}
	hasDelimiters := false
	for _, d := range delimiters {
		if strings.Contains(prompt, d) {
			hasDelimiters = true
			break
		}
	}

	safetyPhrases := []string{
		"do not reveal",
		"ignore previous",
		"stay in character",
		"never mention",
		"private instructions",
		"strictly follow",
	}
	hasSafetyPhrases := false
	for _, p := range safetyPhrases {
		if strings.Contains(prompt, p) {
			hasSafetyPhrases = true
			break
		}
	}

	if hasDelimiters && hasSafetyPhrases {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         10,
			Level:         api.Low,
			Justification: "Uses both structural delimiters and explicit safety constraints.",
		}, nil
	} else if hasDelimiters || hasSafetyPhrases {
		justification := "Missing either structural delimiters or safety constraints: "
		if hasDelimiters {
			justification += "Missing constraints"
		} else {
			justification += "Missing delimiters"
		}
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         6,
			Level:         api.Medium,
			Justification: justification,
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         2,
		Level:         api.High,
		Justification: "No structural delimiters or safety constraints detected in instructions.",
	}, nil
}
