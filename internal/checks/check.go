package checks

import (
	"context"

	"github.com/albertowar/skillauditai/internal/behavioral"
	"github.com/albertowar/skillauditai/pkg/api"
)

// Check defines the interface for all security checks.
type Check interface {
	ID() string
	Name() string
	Weight() float64
	Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error)
}

// DangerousToolsCheck is a placeholder for the dangerous tools check.
// It will be fully implemented in a later task.
type DangerousToolsCheck struct{}

func (d *DangerousToolsCheck) ID() string      { return "dangerous-tools" }
func (d *DangerousToolsCheck) Name() string    { return "Dangerous Tools Check" }
func (d *DangerousToolsCheck) Weight() float64 { return 1.0 }
func (d *DangerousToolsCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	return api.CheckResult{}, nil
}

// AllChecks returns all registered security checks.
func AllChecks() []Check {
	return []Check{
		&DangerousToolsCheck{},
		// ... more will be added
	}
}
