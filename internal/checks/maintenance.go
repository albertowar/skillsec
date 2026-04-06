package checks

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/albertowar/skillauditai/internal/behavioral"
	"github.com/albertowar/skillauditai/pkg/api"
)

type MaintenanceCheck struct{}

func (c *MaintenanceCheck) ID() string      { return "maintenance" }
func (c *MaintenanceCheck) Name() string    { return "Maintenance Audit" }
func (c *MaintenanceCheck) Weight() float64 { return 0.3 }

func (c *MaintenanceCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	if skill.Metadata == nil || skill.Metadata.Maintenance == nil || skill.Metadata.Maintenance.LastUpdated == "" {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         0,
			Level:         api.Info,
			Justification: "No maintenance history available.",
		}, nil
	}

	lastUpdated, err := time.Parse("2006-01-02", skill.Metadata.Maintenance.LastUpdated)
	if err != nil {
		// Try parsing as RFC3339 if ISO date fails
		lastUpdated, err = time.Parse(time.RFC3339, skill.Metadata.Maintenance.LastUpdated)
		if err != nil {
			return api.CheckResult{
				ID:            c.ID(),
				Name:          c.Name(),
				Score:         0,
				Level:         api.Info,
				Justification: "Invalid lastUpdated format.",
			}, nil
		}
	}

	diff := time.Since(lastUpdated)
	diffDays := int(math.Floor(diff.Hours() / 24))

	if diffDays < 90 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         10,
			Level:         api.Low,
			Justification: fmt.Sprintf("Skill is recently maintained (last update %d days ago).", diffDays),
		}, nil
	} else if diffDays < 365 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         5,
			Level:         api.Low,
			Justification: fmt.Sprintf("Skill has not been updated in %d days.", diffDays),
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         2,
		Level:         api.Medium,
		Justification: "Skill is potentially stale (last update over a year ago).",
	}, nil
}
