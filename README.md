# SkillSec

> **Security Auditor for AI Skills**  
> Evaluates AI skills against the [agentskills.io](https://agentskills.io) specification to provide a unified "Safety Score" (0-10).

## 🏗 Architecture

SkillSec is built in Go for high performance and easy distribution as a single binary. It uses concurrent check execution and leverages `langchaingo` for behavioral analysis.

- **Auditor Engine**: Orchestrates security checks and calculates weighted scores.
- **Security Checks**: Modular implementations for static and behavioral analysis.
- **Behavioral Service**: LLM-agnostic testing layer powered by `langchaingo`.

### The Audit Pipeline
1. **Parse**: Converts `SKILL.md` into a structured context.
2. **Context**: Enriches the audit with Git metadata (author, maintenance).
3. **Check**: Executes a registry of security checks concurrently.
4. **Score**: Aggregates check results into a weighted scorecard.
5. **Report**: Generates output in Table or JSON formats.

## 🛠 Getting Started

### Prerequisites
- Go 1.21+

### Installation
```bash
go build -o skillsec ./cmd/skillsec
```

### Basic Usage
```bash
./skillsec path/to/skill.md --api-key YOUR_KEY --provider google
```

## 📖 Documentation

- [CLI Usage Guide](docs/usage.md)
- [Interpreting Results](docs/interpreting-results.md)
- [Architecture Overview](docs/architecture.md)
- [Extending the Auditor](docs/extending.md)

## 🧩 Extending SkillSec

To add a new security check, implement the `Check` interface in `internal/checks`:

1. Define your check in `internal/checks/your_check.go`.
2. Register the check in the `AllChecks()` function in `internal/checks/check.go`.

## 🧪 Development

- **Test**: `go test ./...`
- **Build**: `go build -o skillsec ./cmd/skillsec`

## 📄 License

This project is licensed under the **Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)** license.

- **Non-Commercial**: You may not use this material for commercial purposes.
- **Attribution**: You must give appropriate credit, provide a link to the license, and indicate if changes were made.

See the [LICENSE](LICENSE) file for the full text.
