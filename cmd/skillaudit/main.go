package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"github.com/albertowar/skillauditai/internal/engine"
	"github.com/albertowar/skillauditai/internal/behavioral"
	"github.com/albertowar/skillauditai/pkg/api"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func main() {
	format := flag.String("format", "table", "Output format (table|json)")
	apiKey := flag.String("api-key", os.Getenv("SKILLAUDIT_API_KEY"), "LLM API Key")
	model := flag.String("model", "gemini-1.5-pro", "LLM Model Name")
	provider := flag.String("provider", "google", "LLM Provider (google|openai)")
	baseURL := flag.String("base-url", "", "Custom LLM base URL")
	flag.Parse()

	filePath := flag.Arg(0)
	if filePath == "" {
		fmt.Println("Usage: skillaudit <skill.md> [flags]")
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	skillCtx := engine.ParseSkill(string(content))
	skillCtx.Metadata = engine.GetGitMetadata(filePath)
	bService, _ := behavioral.NewService(*provider, *apiKey, *model, *baseURL)
	auditor := engine.NewAuditor(bService)
	
	report, err := auditor.Audit(context.Background(), skillCtx)
	if err != nil {
		fmt.Printf("Audit failed: %v\n", err)
		os.Exit(1)
	}

	if *format == "json" {
		data, _ := json.MarshalIndent(report, "", "  ")
		fmt.Println(string(data))
	} else {
		renderTable(report)
	}
}

func renderTable(report api.AuditReport) {
	color.New(color.Bold).Printf("\nSkillAuditAI Report - Score: %.1f/10\n\n", report.FinalScore)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Level", "Check", "Score", "Justification"})
	table.SetAutoWrapText(true)
	table.SetColWidth(50)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, r := range report.Results {
		levelColor := color.New(color.FgWhite)
		switch r.Level {
		case api.Critical:
			levelColor = color.New(color.FgRed, color.Bold)
		case api.High:
			levelColor = color.New(color.FgRed)
		case api.Medium:
			levelColor = color.New(color.FgYellow)
		case api.Low:
			levelColor = color.New(color.FgGreen)
		case api.Info:
			levelColor = color.New(color.FgBlue)
		}

		table.Append([]string{
			levelColor.Sprint(r.Level),
			color.New(color.Bold).Sprint(r.Name),
			fmt.Sprintf("%.0f/10", r.Score),
			r.Justification,
		})
	}

	table.Render()
	fmt.Println()
}
