# Interpreting Audit Results

SkillSec provides a Safety Score from 0 to 10. A higher score indicates a safer skill.

## Scoring Logic

The final score is a weighted average of all active checks. Weights are assigned based on the severity of the potential vulnerability:

- **Critical (1.0)**: Direct threats to system integrity.
- **High (0.8)**: Serious misconfigurations.
- **Medium (0.5)**: Significant risks requiring review.
- **Low (0.3)**: Minor best-practice violations.

## Risk Level Definitions

- **🔴 Critical**: Direct Remote Code Execution (RCE) or confirmed data exfiltration path.
- **🟠 High**: Usage of dangerous tools without any safety delimiters or human-in-the-loop (HITL) constraints.
- **🟡 Medium**: Potentially dangerous patterns that should be manually audited.
- **🟢 Low**: Minor issues like missing delimiters but having safety phrases.
- **⚪ Info**: Lack of metadata (author, maintenance date).

## Audit Catalog

| Check ID | Description |
|----------|-------------|
| `dangerous-tools` | Detects high-risk tools like `run_shell_command`. |
| `tool-chaining` | Detects \"Read -> Move\" patterns (e.g., File Read to API Post). |
| `prompt-injection` | Checks for structural delimiters and safety phrases. |
| `exfiltration` | Scans for network sinks and tests robustness via LLM. |
| `indirect-injection` | Audits trust boundaries for external content readers. |
| `secret-scanning` | Searches for hardcoded API keys or secrets. |
