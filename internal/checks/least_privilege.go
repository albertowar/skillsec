package checks

import (
	"context"
	"fmt"
	"strings"

	"github.com/albertowar/skillauditai/internal/behavioral"
	"github.com/albertowar/skillauditai/pkg/api"
)

type LeastPrivilegeCheck struct{}

func (c *LeastPrivilegeCheck) ID() string      { return "least-privilege" }
func (c *LeastPrivilegeCheck) Name() string    { return "Least Privilege Audit" }
func (c *LeastPrivilegeCheck) Weight() float64 { return 0.8 }

func (c *LeastPrivilegeCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	dangerous := []string{"run_shell_command", "write_file", "delete_file"}
	var found []string
	for _, t := range skill.Tools {
		lowerTool := strings.ToLower(t)
		for _, d := range dangerous {
			if lowerTool == d {
				found = append(found, t)
				break
			}
		}
	}

	if len(found) == 0 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         10,
			Level:         api.Low,
			Justification: "No high-risk tools requested.",
		}, nil
	}

	keywords := map[string][]string{
		"run_shell_command": {"bash", "shell", "command", "terminal", "execute", "cli", "run"},
		"write_file":        {"save", "write", "create", "file", "output", "export"},
		"delete_file":       {"remove", "delete", "unlink", "cleanup", "rm"},
	}

	prompt := strings.ToLower(skill.SystemPrompt)
	var unjustified []string

	for _, tool := range found {
		lowerTool := strings.ToLower(tool)
		toolKeywords := keywords[lowerTool]
		hasKeyword := false
		for _, k := range toolKeywords {
			if strings.Contains(prompt, k) {
				hasKeyword = true
				break
			}
		}
		if !hasKeyword {
			unjustified = append(unjustified, tool)
		}
	}

	if len(unjustified) > 0 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         5,
			Level:         api.High,
			Justification: fmt.Sprintf("High-risk tools requested without clear justification in instructions: %s", strings.Join(unjustified, ", ")),
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         10,
		Level:         api.Low,
		Justification: "Requested tools are justified by the skill instructions.",
	}, nil
}
