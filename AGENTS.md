# Agent Instructions: SkillAuditAI

This file provides architectural mandates and type references for AI agents contributing to this repository.

## ⚖️ Core Mandates

1. **Stateless Core**: Everything in `packages/core` must be stateless and "edge-ready." No direct filesystem or network I/O. Use dependency injection or the `Context` object.
2. **Type Safety**: No `any` types. No @ts-ignore. Use Zod for runtime validation where boundaries are crossed.
3. **Spec Alignment**: All parsing and auditing logic must strictly follow the [agentskills.io](https://agentskills.io) specification.

## 🔑 Key Symbols & Types

- `SkillAST`: The unified AST generated from `SKILL.md`.
- `BaseCheck`: Abstract class for all audit checks. Must implement `run(ast: SkillAST): Promise<CheckResult>`.
- `AuditReport`: The final output structure containing the scorecard and findings.

## 📊 Scoring Weights

Consistency in scoring is critical. Use these weights for all `BaseCheck` implementations:
- **Critical (1.0)**: Direct security vulnerabilities (e.g., hardcoded secrets, dangerous shell tools).
- **High (0.8)**: Potential for misuse or prompt injection risks.
- **Medium (0.5)**: Reputation, metadata, or maintenance issues.
- **Low (0.3)**: Stylistic or minor best-practice deviations.

## 🔄 Workflow Requirements

- **TDD**: Every bug fix requires a reproduction test case. Every new check requires a `.spec.ts` file.
- **Parser Logic**: The Markdown parser must target sections: `Tools`, `Instructions`, and `Examples`.
- **Validation**: Run `npm test` before proposing any implementation.
