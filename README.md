# Obviously Not - OpenClaw Skills Suite

A suite of 11 OpenClaw skills for pattern discovery and knowledge distillation.

## Skills

### Patent Analysis Suite (4 skills)

| Skill | Slug | Purpose |
|-------|------|---------|
| [Code Patent Scanner](./code-patent-scanner/) | `code-patent-scanner` | Analyze source code for distinctive patterns |
| [Code Patent Validator](./code-patent-validator/) | `code-patent-validator` | Generate search strategies for code findings |
| [Patent Scanner](./patent-scanner/) | `patent-scanner` | Analyze concept descriptions |
| [Patent Validator](./patent-validator/) | `patent-validator` | Generate search strategies for concepts |

### Principle-Based Distillation Suite (6 skills)

| Skill | Slug | Voice | Purpose |
|-------|------|-------|---------|
| [PBE Extractor](./pbe-extractor/) | `pbe-extractor` | Technical | Extract invariant principles from text |
| [Essence Distiller](./essence-distiller/) | `essence-distiller` | Conversational | Find what actually matters |
| [Principle Comparator](./principle-comparator/) | `principle-comparator` | Technical | Compare two sources for shared patterns |
| [Pattern Finder](./pattern-finder/) | `pattern-finder` | Conversational | Discover what two sources agree on |
| [Principle Synthesizer](./principle-synthesizer/) | `principle-synthesizer` | Technical | Synthesize 3+ sources for Golden Masters |
| [Core Refinery](./core-refinery/) | `core-refinery` | Conversational | Find the core across all sources |

### Utility Skills (1 skill)

| Skill | Slug | Purpose |
|-------|------|---------|
| [Golden Master](./golden-master/) | `golden-master` | Track source/derived file relationships and staleness |

## Installation

Install via OpenClaw agent:

```
@agent install pbe-extractor
```

Or via ClawHub CLI:

```bash
clawhub install pbe-extractor
```

## Recommended Workflows

### Patent Analysis Workflow

```
1. code-patent-scanner → Analyze source code
2. code-patent-validator → Generate search strategies
3. Run searches yourself, document findings
4. Consult patent attorney
```

### PBD: Methodology Extraction Pipeline

Extract invariant principles from multiple methodology documents:

```
# Gather 3+ methodology documents
doc1.md → pbe-extractor → extraction1.json
doc2.md → pbe-extractor → extraction2.json
doc3.md → pbe-extractor → extraction3.json

# Synthesize to find invariants
[all extractions] → principle-synthesizer → invariant_principles.json

# Result: Golden Master candidates (N≥3 principles)
# Use as canonical source for new documentation
```

### PBD: Cross-Project Pattern Validation

Validate whether a principle is truly invariant across projects:

```
# Extract from Project A
project_a_docs → essence-distiller → extraction_a.json

# Extract from Project B
project_b_docs → essence-distiller → extraction_b.json

# Compare to find shared patterns
[extraction_a, extraction_b] → pattern-finder → comparison.json

# Shared patterns (N=2) warrant investigation
# Domain-specific patterns stay at N=1
```

### PBD: Incremental N-Count Building

Build confidence in principles over time:

```
Session 1:
  source1 → pbe-extractor → 5 principles (N=1 each)

Session 2:
  source2 → pbe-extractor → 4 principles (N=1 each)
  [source1, source2] → principle-comparator → 3 shared (N=2)

Session 3:
  source3 → pbe-extractor → 6 principles (N=1 each)
  [all 3 sources] → principle-synthesizer → 2 Golden Masters (N=3)
```

**N-Count State Management**: The skills generate observations with N-counts, but **you are responsible for tracking state across sessions**. Use `source_hash` values to link extractions over time. The skills don't maintain persistent state — each invocation is independent.

### Integration: PBD + Golden Master

After identifying Golden Master candidates:

```
# PBD identifies invariants
[sources] → principle-synthesizer → golden_master_candidates

# Track relationships with Golden Master skill
candidates → golden-master:establish → tracked_relationships

# Later: detect when source documents change
golden-master:validate → staleness_report
```

### Golden Master Pattern Details

The golden-master skill tracks source/derived file relationships using checksums.

**Metadata Format** (in-file comments):

```markdown
<!-- Source file declares its derived works -->
<!-- golden-master:source checksum:a1b2c3d4 derived:[README.md,QUICKSTART.md] -->

<!-- Derived file stores source checksum at derivation time -->
<!-- golden-master:derived source:ARCHITECTURE.md source_checksum:a1b2c3d4 derived_at:2026-02-04 -->
```

**Workflow**:
1. `analyze` - Discover potential relationships (content overlap detection)
2. `establish` - Create tracking metadata for confirmed relationships
3. `validate` - Check for staleness (source checksum changed since derivation)
4. `refresh` - Update metadata after manually syncing derived content

**Problem Solved**: When README derives from ARCHITECTURE.md, how do you know when ARCHITECTURE.md changed and README is stale? Golden Master tracks this automatically via checksums.

## Agent Identity Evolution

Skills use **Agent Identity** patterns to define AI persona. The pattern evolved over skill development:

### 4-Field Pattern (Patent Skills, v1.0.0)

```markdown
**Role**: What the agent does
**Approach**: How the agent works
**Boundaries**: What the agent won't do
**Tone**: How the agent communicates
```

Used by: `code-patent-scanner`, `code-patent-validator`, `patent-scanner`, `patent-validator`

### 6-Field Pattern (PBD + Utility Skills, v1.2.0+)

```markdown
**Role**: What the agent does
**Understands**: Theory of Mind — what the agent knows about user context
**Approach**: How the agent works
**Boundaries**: What the agent won't do
**Tone**: How the agent communicates
**Opening Pattern**: Example first-response to calibrate expectations
```

Used by: All PBD skills and `golden-master`

**Why the evolution?** The `Understands` field helps the AI acknowledge user pain points upfront. The `Opening Pattern` calibrates response tone from the first interaction. Both improve perceived empathy and consistency.

**For new skills**: Use the 6-field pattern.

---

## Multi-Voice Strategy

Each PBD operation has two voices:

| Operation | Technical Voice | Conversational Voice |
|-----------|-----------------|---------------------|
| Extraction | `pbe-extractor` | `essence-distiller` |
| Comparison | `principle-comparator` | `pattern-finder` |
| Synthesis | `principle-synthesizer` | `core-refinery` |

**Same methodology, different tone.** Use whichever fits your context:
- Technical: Precise language, structured output, detailed metrics
- Conversational: Warm language, discovery-focused, simplified output

**Mixing voices is fine**: You can extract with `essence-distiller` (conversational) and synthesize with `principle-synthesizer` (technical) based on your context. The output schemas are compatible — only tone and metric detail differ.

## Maintainer Notes

### Sync Requirement

**IMPORTANT**: Skill metadata exists in two places that must stay in sync:

1. **SKILL.md files** (source of truth for ClawHub)
2. **Go structs** in `internal/web/views/skills.go` (web landing pages)

When updating skill descriptions, features, or metadata:
1. Update the SKILL.md file
2. Update the matching struct in `internal/web/views/skills.go`
3. Run `go test ./internal/web/handlers/... -run Skills` to verify
4. Republish to ClawHub

### UPL Compliance Check (Patent Skills Only)

Before publishing patent skills, verify no banned terms:

```bash
# Check SKILL.md files
for skill in code-patent-scanner code-patent-validator patent-scanner patent-validator; do
  echo "=== $skill ==="
  grep -n -i "invention\|patentable\|non-obvious\|prior art" skills/$skill/SKILL.md | \
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

   # Patent skills
   for skill in code-patent-scanner code-patent-validator patent-scanner patent-validator; do
     clawhub --registry https://clawhub.ai publish $skill
   done

   # PBD skills - Technical voice
   for skill in pbe-extractor principle-comparator principle-synthesizer; do
     clawhub --registry https://clawhub.ai publish $skill
   done

   # PBD skills - Conversational voice
   for skill in essence-distiller pattern-finder core-refinery; do
     clawhub --registry https://clawhub.ai publish $skill
   done

   # Utility skills
   clawhub --registry https://clawhub.ai publish golden-master
   ```

## Links

- **Website**: https://github.com/Obviously-Not/patent-skills
- **ClawHub**: https://clawhub.ai/skills/pbe-extractor

## Disclaimers

### Patent Skills

These skills identify distinctive technical patterns and generate research strategies.
They do NOT provide legal advice, assess patentability, or replace professional counsel.
Always consult a registered patent attorney for intellectual property guidance.

### PBD Skills

These skills extract patterns from content, not verified truth. Principles are observations
that need validation (N≥2) and human judgment. Use for analysis, not as authority.

---

*Built by Obviously Not — Tools for thought, not conclusions.*
