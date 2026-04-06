package checks

import (
	"context"
	"testing"
	"github.com/albertowar/skillsec/pkg/api"
)

func TestDangerousToolsCheck(t *testing.T) {
	check := &DangerousToolsCheck{}
	ctx := context.Background()

	tests := []struct {
		name          string
		tools         []string
		expectedScore float64
		expectedLevel api.RiskLevel
	}{
		{
			name:          "No dangerous tools",
			tools:         []string{"read_file", "list_dir"},
			expectedScore: 10,
			expectedLevel: api.Low,
		},
		{
			name:          "Single dangerous tool",
			tools:         []string{"run_shell_command", "read_file"},
			expectedScore: 0,
			expectedLevel: api.Critical,
		},
		{
			name:          "Case insensitive check",
			tools:         []string{"Run_Shell_Command"},
			expectedScore: 0,
			expectedLevel: api.Critical,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := api.SkillContext{Tools: tt.tools}
			res, err := check.Run(ctx, skill, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res.Score != tt.expectedScore {
				t.Errorf("expected score %v, got %v", tt.expectedScore, res.Score)
			}
			if res.Level != tt.expectedLevel {
				t.Errorf("expected level %v, got %v", tt.expectedLevel, res.Level)
			}
		})
	}
}

func TestDangerousToolsCheck_Run_OpenAI(t *testing.T) {
	check := &DangerousToolsCheck{}
	skill := api.SkillContext{
		Provider: "openai",
		Tools:    []string{"code_interpreter"},
	}
	res, _ := check.Run(context.Background(), skill, nil)
	if res.Level != api.Critical {
		t.Errorf("expected Critical, got %v", res.Level)
	}
}

func TestDangerousToolsCheck_Run_Anthropic(t *testing.T) {
	check := &DangerousToolsCheck{}
	skill := api.SkillContext{
		Provider: "anthropic",
		Tools:    []string{"computer"},
	}
	res, _ := check.Run(context.Background(), skill, nil)
	if res.Level != api.Critical {
		t.Errorf("expected Critical, got %v", res.Level)
	}
}
