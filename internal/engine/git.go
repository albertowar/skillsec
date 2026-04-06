package engine

import (
	"os/exec"
	"strings"
	"strconv"
	"time"
	"github.com/albertowar/skillauditai/pkg/api"
)

func GetGitMetadata(filePath string) *api.SkillMetadata {
	cmd := exec.Command("git", "log", "-1", "--format=%an|%ae|%at|%G?", "--", filePath)
	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return nil
	}

	parts := strings.Split(strings.TrimSpace(string(out)), "|")
	if len(parts) < 4 {
		return nil
	}

	timestamp, _ := strconv.ParseInt(parts[2], 10, 64)
	
	return &api.SkillMetadata{
		Author: &struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			IsVerified bool   `json:"isVerified"`
		}{
			Name:       parts[0],
			Email:      parts[1],
			IsVerified: parts[3] == "G",
		},
		Maintenance: &struct {
			LastUpdated string `json:"lastUpdated"`
			Version     string `json:"version"`
		}{
			LastUpdated: time.Unix(timestamp, 0).Format(time.RFC3339),
		},
	}
}
