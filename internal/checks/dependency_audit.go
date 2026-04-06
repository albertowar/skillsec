package checks

import (
	"context"
	"regexp"
	"strings"

	"github.com/albertowar/skillsec/internal/behavioral"
	"github.com/albertowar/skillsec/pkg/api"
)

type DependencyAuditCheck struct{}

func (c *DependencyAuditCheck) ID() string      { return "dependency-audit" }
func (c *DependencyAuditCheck) Name() string    { return "Dependency Audit" }
func (c *DependencyAuditCheck) Weight() float64 { return 0.5 }

var dependencyPatterns = []*regexp.Regexp{
	regexp.MustCompile(`import\s+['"]?[\w-]+['"]?`),
	regexp.MustCompile(`require\s*\(['"]?[\w-]+['"]?\)`),
	regexp.MustCompile(`\[.*?\]\(.*?\.md\)`), // Skill links
}

func (c *DependencyAuditCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	allContent := strings.Join([]string{
		skill.SystemPrompt,
		strings.Join(skill.Examples, "\n"),
	}, "\n")
	lowerContent := strings.ToLower(allContent)

	installers := []string{"pip install", "npm install", "cargo add", "gem install", "apt install"}
	var foundInstallers []string
	for _, i := range installers {
		if strings.Contains(lowerContent, i) {
			foundInstallers = append(foundInstallers, i)
		}
	}

	if len(foundInstallers) > 0 {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         4,
			Level:         api.Medium,
			Justification: "Skill contains external library installers: " + strings.Join(foundInstallers, ", "),
		}, nil
	}

	foundImports := false
	for _, re := range dependencyPatterns {
		if re.MatchString(allContent) {
			foundImports = true
			break
		}
	}

	if foundImports {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         7,
			Level:         api.Low,
			Justification: "Skill references external skills or libraries but lacks explicit installers.",
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         10,
		Level:         api.Low,
		Justification: "No external dependencies detected.",
	}, nil
}
