package checks

import (
	"context"
	"strings"

	"github.com/albertowar/skillsec/internal/behavioral"
	"github.com/albertowar/skillsec/pkg/api"
)

type ToolChainingCheck struct{}

func (c *ToolChainingCheck) ID() string      { return "tool-chaining" }
func (c *ToolChainingCheck) Name() string    { return "Tool Chaining Risk Audit" }
func (c *ToolChainingCheck) Weight() float64 { return 0.9 }

func (c *ToolChainingCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	sources := []string{"read_file", "list_directory", "get_history", "web_search"}
	externalSinks := []string{"send_email", "api_post", "webhook_call", "http_request"}
	internalSinks := []string{"write_file", "append_to_log"}

	hasSource := false
	hasExternalSink := false
	hasInternalSink := false

	for _, t := range skill.Tools {
		lowerTool := strings.ToLower(t)
		for _, s := range sources {
			if lowerTool == s {
				hasSource = true
				break
			}
		}
		for _, s := range externalSinks {
			if lowerTool == s {
				hasExternalSink = true
				break
			}
		}
		for _, s := range internalSinks {
			if lowerTool == s {
				hasInternalSink = true
				break
			}
		}
	}

	prompt := strings.ToLower(skill.SystemPrompt)
	hitlKeywords := []string{"ask the user", "confirm with user", "require approval"}
	hasHITL := false
	for _, k := range hitlKeywords {
		if strings.Contains(prompt, k) {
			hasHITL = true
			break
		}
	}

	if hasSource && hasExternalSink {
		if !hasHITL {
			return api.CheckResult{
				ID:            c.ID(),
				Name:          c.Name(),
				Score:         0,
				Level:         api.Critical,
				Justification: "Dangerous tool chain detected: Source + External Sink present without Human-in-the-loop (HITL) instructions.",
			}, nil
		}
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         3,
			Level:         api.High,
			Justification: "Risk detected: Source + External Sink present. HITL instructions are present but risk remains high due to potential for bypass.",
		}, nil
	}

	if hasSource && hasInternalSink {
		if !hasHITL {
			return api.CheckResult{
				ID:            c.ID(),
				Name:          c.Name(),
				Score:         3,
				Level:         api.High,
				Justification: "Risk detected: Source + Internal Sink present without HITL instructions.",
			}, nil
		}
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         7,
			Level:         api.Medium,
			Justification: "Source + Internal Sink present with HITL instructions.",
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         10,
		Level:         api.Low,
		Justification: "No dangerous tool chains detected.",
	}, nil
}
