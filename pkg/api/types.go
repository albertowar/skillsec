package api

import "time"

type RiskLevel string

const (
	Critical RiskLevel = "Critical"
	High     RiskLevel = "High"
	Medium   RiskLevel = "Medium"
	Low      RiskLevel = "Low"
	Info     RiskLevel = "Info"
)

type CheckResult struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Score         float64   `json:"score"` // 0-10
	Level         RiskLevel `json:"level"`
	Justification string    `json:"justification"`
}

type AuditReport struct {
	SkillHash  string        `json:"skillHash"`
	FinalScore float64       `json:"finalScore"`
	Results    []CheckResult `json:"results"`
	Timestamp  time.Time     `json:"timestamp"`
}

type SkillMetadata struct {
	Author *struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		IsVerified bool   `json:"isVerified"`
	} `json:"author,omitempty"`
	Maintenance *struct {
		LastUpdated string `json:"lastUpdated"`
		Version     string `json:"version"`
	} `json:"maintenance,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
}

type SkillContext struct {
	Raw          string         `json:"raw"`
	Tools        []string       `json:"tools"`
	Provider     string         `json:"provider"`
	SystemPrompt string         `json:"systemPrompt"`
	Examples     []string       `json:"examples"`
	Metadata     *SkillMetadata `json:"metadata,omitempty"`
}
