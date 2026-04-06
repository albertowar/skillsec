# Extending SkillSec

## Adding a New Check

1. Create a new file in `internal/checks/`.
2. Implement the `Check` interface:

```go
package checks

import (
	"context"
	"github.com/albertowar/skillsec/pkg/api"
	"github.com/albertowar/skillsec/internal/behavioral"
)

type MyNewCheck struct{}

func (c *MyNewCheck) ID() string      { return "my-check" }
func (c *MyNewCheck) Name() string    { return "My Security Audit" }
func (c *MyNewCheck) Weight() float64 { return 0.5 }

func (c *MyNewCheck) Run(ctx context.Context, skill api.SkillContext, b *behavioral.Service) (api.CheckResult, error) {
	// Audit logic here
    return api.CheckResult{
        ID:            c.ID(),
        Name:          c.Name(),
        Score:         10,
        Level:         api.Low,
        Justification: "No issues found.",
    }, nil
}
```

3. Register the check in `internal/checks/check.go` within the `AllChecks()` function.
4. Add tests in a corresponding `_test.go` file in the same directory.
