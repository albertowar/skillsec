# SkillAuditAI

> **Security Auditor for AI Skills**  
> Evaluates AI skills against the [agentskills.io](https://agentskills.io) specification to provide a unified "Safety Score" (0-10).

## 🏗 Architecture

SkillAuditAI is built as a TypeScript monorepo to ensure the auditing engine remains portable and "edge-ready."

- **`@skillauditai/core`**: The stateless engine. Handles parsing, check orchestration, and scoring.
- **`@skillauditai/cli`**: The user-facing tool for file I/O and terminal rendering.

### The Audit Pipeline
1. **Parse**: Converts `SKILL.md` into a structured `SkillAST`.
2. **Context**: (Optional) Enriches the audit with external metadata.
3. **Check**: Executes a registry of `BaseCheck` implementations.
4. **Score**: Aggregates check results into a weighted scorecard.
5. **Report**: Generates output in Table, JSON, or SARIF formats.

## 🛠 Getting Started

### Prerequisites
- Node.js 18+
- npm

### Installation
```bash
npm install
npm run build
```

## 📖 Documentation

- [CLI Usage Guide](docs/usage.md)
- [Interpreting Results](docs/interpreting-results.md)
- [Architecture Overview](docs/architecture.md)
- [Extending the Auditor](docs/extending.md)

## 🧩 Extending SkillAuditAI

To add a new security check, implement the `BaseCheck` interface in `@skillauditai/core`:

1. Define your check in `packages/core/src/checks/`.
2. Assign a weight: `Critical (1.0)`, `High (0.8)`, `Medium (0.5)`, or `Low (0.3)`.
3. Register the check in the `Auditor` engine.

## 🧪 Development

- **Test**: `npm test` (Uses Vitest)
- **Build**: `npm run build`
- **Lint**: `npm run lint`
