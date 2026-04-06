package checks

import (
	"context"
	"fmt"

	"github.com/albertowar/skillsec/internal/behavioral"
	"github.com/albertowar/skillsec/pkg/api"
)

type VerifiedAuthorCheck struct{}

func (c *VerifiedAuthorCheck) ID() string      { return "verified-author" }
func (c *VerifiedAuthorCheck) Name() string    { return "Verified Author Audit" }
func (c *VerifiedAuthorCheck) Weight() float64 { return 0.5 }

func (c *VerifiedAuthorCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	if skill.Metadata == nil || skill.Metadata.Author == nil || (skill.Metadata.Author.Name == "" && skill.Metadata.Author.Email == "") {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         0,
			Level:         api.Medium,
			Justification: "No author attribution found.",
		}, nil
	}

	meta := skill.Metadata.Author
	if meta.IsVerified {
		return api.CheckResult{
			ID:            c.ID(),
			Name:          c.Name(),
			Score:         10,
			Level:         api.Low,
			Justification: fmt.Sprintf("Verified author: %s <%s>", meta.Name, meta.Email),
		}, nil
	}

	return api.CheckResult{
		ID:            c.ID(),
		Name:          c.Name(),
		Score:         7,
		Level:         api.Low,
		Justification: fmt.Sprintf("Author identified but not verified: %s <%s>", meta.Name, meta.Email),
	}, nil
}
