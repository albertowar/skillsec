package checks

import (
	"context"
	"fmt"
	"strings"
	"github.com/albertowar/skillauditai/pkg/api"
	"github.com/albertowar/skillauditai/internal/behavioral"
)

type DangerousToolsCheck struct{}

func (c *DangerousToolsCheck) ID() string      { return "dangerous-tools" }
func (c *DangerousToolsCheck) Name() string    { return "Dangerous Tools Audit" }
func (c *DangerousToolsCheck) Weight() float64 { return 1.0 }

func (c *DangerousToolsCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	dangerous := []string{"run_shell_command", "write_file", "delete_file"}
	var found []string
	
	for _, t := range skill.Tools {
		for _, d := range dangerous {
			if t == d {
				found = append(found, t)
			}
		}
	}

	if len(found) > 0 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         0,
			Level:         api.Critical,
			Justification: fmt.Sprintf("Skill requests highly dangerous tools: %s", strings.Join(found, ", ")),
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         10,
		Level:         api.Low,
		Justification: "No dangerous tools detected.",
	}, nil
}
