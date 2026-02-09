# AI Agent Skills

Open-source skills for patent analysis and principle extraction. Built on the [Agent Skills](https://agentskills.io) standard.

## What Are Agent Skills?

Agent skills are portable instructions that extend what AI coding agents can do. Each skill is a `SKILL.md` file — a Markdown document with YAML frontmatter that any compatible agent can read and execute.

The [Agent Skills standard](https://agentskills.io) was originated by Anthropic and adopted by 27+ tools including Claude Code, Gemini CLI, Codex CLI, Cursor, GitHub Copilot, Goose, and more. Write once, use everywhere.

### Format

```
skill-name/
  SKILL.md           # Required — instructions + YAML frontmatter
  scripts/           # Optional — executable code
  references/        # Optional — detailed docs loaded on demand
  assets/            # Optional — templates, data files
```

The `SKILL.md` file contains:

```yaml
---
name: skill-name
description: What this skill does and when to use it.
---

# Instructions

Step-by-step instructions the agent follows...
```

Agents read the `description` field to decide when to activate the skill. The Markdown body contains the full instructions, loaded on demand to keep context usage efficient.

## Installation

### Claude Code / Gemini CLI / Cursor

Clone the repo and copy any skill into your agent's skills directory:

```bash
git clone https://github.com/Obviously-Not/patent-skills
cp -r patent-skills/code-patent-scanner .claude/skills/
```

The skill becomes available as a `/code-patent-scanner` slash command. Claude Code (and other compatible agents) will also auto-invoke it when it detects a relevant task.

**Project-level** (shared with your team via version control):
```
your-project/.claude/skills/code-patent-scanner/SKILL.md
```

**Personal** (available in all your projects):
```
~/.claude/skills/code-patent-scanner/SKILL.md
```

### Any LLM Agent (Copy/Paste)

Open the `SKILL.md` file on GitHub, copy the contents, and paste directly into your agent's chat. The agent will follow the instructions immediately. This works with any LLM — no special setup required.

### OpenClaw

```bash
clawhub install leegitw/code-patent-scanner
```

Skills install to `./skills/` and OpenClaw loads them automatically.

## Skills Catalog

### Patent Analysis

These skills help you discover and validate patentable patterns in code and concepts.

| Skill | What It Does | Start Here? |
|-------|-------------|-------------|
| **[Code Patent Scanner](https://github.com/Obviously-Not/patent-skills/tree/main/code-patent-scanner)** | Scans your codebase for distinctive patterns. Returns structured scoring and evidence for patent consultation. | Yes — for code |
| **[Code Patent Validator](https://github.com/Obviously-Not/patent-skills/tree/main/code-patent-validator)** | Takes scanner findings and generates search queries to research existing implementations. | Use after Code Patent Scanner |
| **[Patent Scanner](https://github.com/Obviously-Not/patent-skills/tree/main/patent-scanner)** | Describe your concept and discover what makes it distinctive. Structured analysis for patent consultation. | Yes — for concepts |
| **[Patent Validator](https://github.com/Obviously-Not/patent-skills/tree/main/patent-validator)** | Takes concept analysis and generates search queries to research the landscape. | Use after Patent Scanner |

**Recommended workflow:** Scanner → Validator → Research → Attorney

Start with the appropriate Scanner to identify patterns, then use the matching Validator to generate search strategies. Use those strategies for your own research, then bring your findings to a patent attorney.

### Principle-Based Extraction

These skills extract and compare the invariant ideas that survive any rephrasing — useful for understanding what's actually novel in your work.

| Skill | What It Does | When to Use |
|-------|-------------|------------|
| **[PBE Extractor](https://github.com/Obviously-Not/patent-skills/tree/main/pbe-extractor)** | Extracts invariant principles from any text. Finds the ideas that survive rephrasing. | First step — extract before comparing or synthesizing |
| **[Principle Comparator](https://github.com/Obviously-Not/patent-skills/tree/main/principle-comparator)** | Compares two sources to find shared and divergent principles. | After extraction — with 2 sources |
| **[Principle Synthesizer](https://github.com/Obviously-Not/patent-skills/tree/main/principle-synthesizer)** | Synthesizes invariant principles from 3+ sources. Finds the core that survives across all expressions. | After extraction — with 3+ sources |
| **[Essence Distiller](https://github.com/Obviously-Not/patent-skills/tree/main/essence-distiller)** | Finds what actually matters in your content. The ideas that survive any rephrasing. | Extract first, then compare or refine |
| **[Pattern Finder](https://github.com/Obviously-Not/patent-skills/tree/main/pattern-finder)** | Discovers what two sources agree on. Finds the signal in the noise. | After extraction — with 2 sources |

### Advanced

| Skill | What It Does | When to Use |
|-------|-------------|------------|
| **[Core Refinery](https://github.com/Obviously-Not/patent-skills/tree/main/core-refinery)** | Finds the core that runs through everything. The ideas that survive across all your sources. | After 3+ extractions |
| **[Golden Master](https://github.com/Obviously-Not/patent-skills/tree/main/golden-master)** | Tracks source-of-truth relationships between files. Knows when derived content becomes stale. | After synthesis — track derivations |

## Important Disclaimer

**These skills are not legal advice.** They identify technical patterns, not patentability.

What they DO:
- Identify distinctive patterns in code and concepts
- Generate search queries for prior art research
- Suggest research strategies
- Document technical evidence

What they DON'T:
- Search patent databases
- Assess patentability
- Provide legal conclusions
- Replace professional counsel

Always consult a patent attorney for IP guidance.

## Links

- **Skills page:** [obviouslynot.ai/skills](https://obviouslynot.ai/skills)
- **GitHub:** [github.com/Obviously-Not/patent-skills](https://github.com/Obviously-Not/patent-skills)
- **ClawHub:** [clawhub.ai](https://clawhub.ai) (search "leegitw")
- **Agent Skills standard:** [agentskills.io](https://agentskills.io)
- **ObviouslyNot:** [obviouslynot.ai](https://obviouslynot.ai)
