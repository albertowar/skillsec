package checks

import (
	"context"
	"testing"
	"time"

	"github.com/albertowar/skillsec/pkg/api"
)

func TestSecretScanningCheck(t *testing.T) {
	check := &SecretScanningCheck{}
	ctx := context.Background()

	tests := []struct {
		name          string
		raw           string
		expectedScore float64
		expectedLevel api.RiskLevel
	}{
		{
			name:          "No secrets",
			raw:           "This is a safe skill.",
			expectedScore: 10,
			expectedLevel: api.Low,
		},
		{
			name:          "OpenAI secret",
			raw:           "My key is sk-1234567890abcdefghij1234567890",
			expectedScore: 0,
			expectedLevel: api.Critical,
		},
		{
			name:          "GitHub token",
			raw:           "Token: ghp_123456789012345678901234567890123456",
			expectedScore: 0,
			expectedLevel: api.Critical,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := api.SkillContext{Raw: tt.raw}
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

func TestDependencyAuditCheck(t *testing.T) {
	check := &DependencyAuditCheck{}
	ctx := context.Background()

	tests := []struct {
		name          string
		prompt        string
		expectedScore float64
		expectedLevel api.RiskLevel
	}{
		{
			name:          "No dependencies",
			prompt:        "Just some instructions.",
			expectedScore: 10,
			expectedLevel: api.Low,
		},
		{
			name:          "With installer",
			prompt:        "Run pip install requests",
			expectedScore: 4,
			expectedLevel: api.Medium,
		},
		{
			name:          "With import",
			prompt:        "Use import os",
			expectedScore: 7,
			expectedLevel: api.Low,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := api.SkillContext{SystemPrompt: tt.prompt}
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

func TestVerifiedAuthorCheck(t *testing.T) {
	check := &VerifiedAuthorCheck{}
	ctx := context.Background()

	tests := []struct {
		name          string
		metadata      *api.SkillMetadata
		expectedScore float64
		expectedLevel api.RiskLevel
	}{
		{
			name:          "No metadata",
			metadata:      nil,
			expectedScore: 0,
			expectedLevel: api.Medium,
		},
		{
			name: "Verified author",
			metadata: &api.SkillMetadata{
				Author: &struct {
					Name       string `json:"name"`
					Email      string `json:"email"`
					IsVerified bool   `json:"isVerified"`
				}{Name: "Test", Email: "test@example.com", IsVerified: true},
			},
			expectedScore: 10,
			expectedLevel: api.Low,
		},
		{
			name: "Unverified author",
			metadata: &api.SkillMetadata{
				Author: &struct {
					Name       string `json:"name"`
					Email      string `json:"email"`
					IsVerified bool   `json:"isVerified"`
				}{Name: "Test", Email: "test@example.com", IsVerified: false},
			},
			expectedScore: 7,
			expectedLevel: api.Low,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := api.SkillContext{Metadata: tt.metadata}
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

func TestMaintenanceCheck(t *testing.T) {
	check := &MaintenanceCheck{}
	ctx := context.Background()

	now := time.Now()
	recent := now.AddDate(0, 0, -30).Format("2006-01-02")
	old := now.AddDate(0, 0, -120).Format("2006-01-02")
	stale := now.AddDate(-2, 0, 0).Format("2006-01-02")

	tests := []struct {
		name          string
		metadata      *api.SkillMetadata
		expectedScore float64
		expectedLevel api.RiskLevel
	}{
		{
			name:          "No metadata",
			metadata:      nil,
			expectedScore: 0,
			expectedLevel: api.Info,
		},
		{
			name: "Recent update",
			metadata: &api.SkillMetadata{
				Maintenance: &struct {
					LastUpdated string `json:"lastUpdated"`
					Version     string `json:"version"`
				}{LastUpdated: recent},
			},
			expectedScore: 10,
			expectedLevel: api.Low,
		},
		{
			name: "Older update",
			metadata: &api.SkillMetadata{
				Maintenance: &struct {
					LastUpdated string `json:"lastUpdated"`
					Version     string `json:"version"`
				}{LastUpdated: old},
			},
			expectedScore: 5,
			expectedLevel: api.Low,
		},
		{
			name: "Stale update",
			metadata: &api.SkillMetadata{
				Maintenance: &struct {
					LastUpdated string `json:"lastUpdated"`
					Version     string `json:"version"`
				}{LastUpdated: stale},
			},
			expectedScore: 2,
			expectedLevel: api.Medium,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := api.SkillContext{Metadata: tt.metadata}
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
