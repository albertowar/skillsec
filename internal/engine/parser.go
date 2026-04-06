package engine

import (
	"regexp"
	"strings"

	"github.com/albertowar/skillsec/internal/provider"
	"github.com/albertowar/skillsec/pkg/api"
)

func ParseSkill(content string) api.SkillContext {
	getSection := func(name string) string {
		// Use a simpler regex that RE2 supports
		re := regexp.MustCompile(`(?i)## ` + name + `\s+([\s\S]*?)(\n## |$)`)
		match := re.FindStringSubmatch(content)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		}
		return ""
	}

	toolsSection := getSection("Tools")
	var tools []string
	if toolsSection != "" {
		lines := strings.Split(toolsSection, "\n")
		re := regexp.MustCompile(`-\s*` + "`" + `?([\w_]+)` + "`" + `?`)
		for _, line := range lines {
			match := re.FindStringSubmatch(line)
			if len(match) > 1 {
				tools = append(tools, match[1])
			}
		}
	}

	examplesSection := getSection("Examples")
	var examples []string
	if examplesSection != "" {
		parts := strings.Split(examplesSection, "\n---\n")
		for _, p := range parts {
			if trimmed := strings.TrimSpace(p); trimmed != "" {
				examples = append(examples, trimmed)
			}
		}
	}

	return api.SkillContext{
		Raw:          content,
		Tools:        tools,
		Provider:     provider.Detect(tools),
		SystemPrompt: getSection("Instructions"),
		Examples:     examples,
	}
}
