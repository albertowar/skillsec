package engine

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
	"github.com/albertowar/skillauditai/pkg/api"
	"github.com/albertowar/skillauditai/internal/checks"
	"github.com/albertowar/skillauditai/internal/behavioral"
)

type Auditor struct {
	Checks     []checks.Check
	Behavioral *behavioral.Service
}

func NewAuditor(b *behavioral.Service) *Auditor {
	return &Auditor{
		Checks:     checks.AllChecks(),
		Behavioral: b,
	}
}

func (a *Auditor) Audit(ctx context.Context, skill api.SkillContext) (api.AuditReport, error) {
	var wg sync.WaitGroup
	results := make([]api.CheckResult, len(a.Checks))
	
	for i, c := range a.Checks {
		wg.Add(1)
		go func(idx int, check checks.Check) {
			defer wg.Done()
			res, err := check.Run(ctx, skill, a.Behavioral)
			if err != nil {
				results[idx] = api.CheckResult{
					ID:            check.ID(),
					Name:          check.Name(),
					Score:         0,
					Level:         api.Critical,
					Justification: fmt.Sprintf("Check failed: %v", err),
				}
			} else {
				results[idx] = res
			}
		}(i, c)
	}

	wg.Wait()

	var weightedScore float64
	var totalWeight float64
	for i, r := range results {
		w := a.Checks[i].Weight()
		weightedScore += r.Score * w
		totalWeight += w
	}

	finalScore := 10.0
	if totalWeight > 0 {
		finalScore = weightedScore / totalWeight
	}

	hash := sha256.Sum256([]byte(skill.Raw))

	return api.AuditReport{
		SkillHash:  fmt.Sprintf("%x", hash),
		FinalScore: finalScore,
		Results:    results,
		Timestamp:  time.Now(),
	}, nil
}
