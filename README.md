# Patent Pioneer - OpenClaw Skills Suite

A suite of 4 OpenClaw skills for discovering and validating distinctive technical patterns.

## Skills

| Skill | Slug | Purpose |
|-------|------|---------|
| [Code Patent Scanner](./code-patent-scanner/) | `code-patent-scanner` | Analyze source code for distinctive patterns |
| [Code Patent Validator](./code-patent-validator/) | `code-patent-validator` | Generate search strategies for code findings |
| [Patent Scanner](./patent-scanner/) | `patent-scanner` | Analyze concept descriptions |
| [Patent Validator](./patent-validator/) | `patent-validator` | Generate search strategies for concepts |

## Installation

Install via OpenClaw agent:

```
@agent install code-patent-scanner
```

Or via ClawHub CLI:

```bash
clawhub install code-patent-scanner
```

## Recommended Workflow

### For Code Analysis

```
1. code-patent-scanner → Analyze source code
2. code-patent-validator → Generate search strategies
3. Run searches yourself, document findings
4. Consult patent attorney
```

### For Concept Descriptions

```
1. patent-scanner → Analyze concept description
2. patent-validator → Generate search strategies
3. Run searches yourself, document findings
4. Consult patent attorney
```

## Maintainer Notes

### Sync Requirement

**IMPORTANT**: Skill metadata exists in two places that must stay in sync:

1. **SKILL.md files** (source of truth for ClawHub)
2. **Go structs** in `internal/web/views/skills.go` (web landing pages)

When updating skill descriptions, features, or metadata:
1. Update the SKILL.md file
2. Update the matching struct in `internal/web/views/skills.go`
3. Run `go test ./internal/web/handlers/... -run UPL` to verify no UPL violations
4. Republish to ClawHub

### UPL Compliance Check

Before publishing, verify no banned terms in skills:

```bash
# Check SKILL.md files
for skill in code-patent-scanner code-patent-validator patent-scanner patent-validator; do
  echo "=== $skill ==="
  grep -n -i "invention\|patentable\|non-obvious\|prior art" $skill/SKILL.md | \
    grep -v "Never Use\|### Never\|patentable inventions" || echo "PASS"
done

# Check Go structs
go test ./internal/web/handlers/... -run UPL
```

## Publishing (Maintainers Only)

1. Set up authentication:
   ```bash
   cp .env.example .env
   # Edit .env with your CLAWHUB_TOKEN
   ```

2. Publish all skills:
   ```bash
   source .env
   for skill in code-patent-scanner code-patent-validator patent-scanner patent-validator; do
     clawhub --workdir . --registry https://www.clawhub.ai publish $skill
   done
   ```

## Links

- **Website**: https://app.obviouslynot.ai/skills
- **ClawHub**: https://www.clawhub.ai/skills/code-patent-scanner

## Disclaimer

These skills identify distinctive technical patterns and generate research strategies.
They do NOT provide legal advice, assess patentability, or replace professional counsel.
Always consult a registered patent attorney for intellectual property guidance.

---

*Built by Obviously Not - Tools for thought, not conclusions.*
