package checks

import (
	"context"

	"github.com/albertowar/skillsec/internal/behavioral"
	"github.com/albertowar/skillsec/pkg/api"
)

// Check defines the interface for all security checks.
type Check interface {
	ID() string
	Name() string
	Weight() float64
	Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error)
}

// AllChecks returns all registered security checks.
func AllChecks() []Check {
	return []Check{
		&DangerousToolsCheck{},
		&SecretScanningCheck{},
		&DependencyAuditCheck{},
		&VerifiedAuthorCheck{},
		&MaintenanceCheck{},
		&PromptInjectionCheck{},
		&ExfiltrationCheck{},
		&IndirectInjectionCheck{},
		&LeastPrivilegeCheck{},
		&ToolChainingCheck{},
	}
}
