---
name: Patent Scanner
description: Describe your concept in plain language and identify potentially distinctive aspects with structured scoring. Works with any concept type - software, hardware, processes, methods. NOT legal advice.
homepage: https://app.obviouslynot.ai/skills/patent-scanner
user-invocable: true
emoji: 🔍
---

# Patent Scanner

Analyze concept descriptions to identify potentially distinctive aspects using the recombination framework with structured scoring. Works with any concept type - no codebase required.

## When to Use

Activate this skill when the user asks to:
- "Analyze my concept"
- "What's distinctive about this?"
- "Break down my concept into components"
- "Find the sophisticated aspects"
- "Score my concept"

## Important Limitations

- This is TECHNICAL analysis, not legal advice
- Output identifies "potentially distinctive aspects" not "patentable inventions"
- Cannot search existing implementations (use patent-validator for that)
- Always recommend professional consultation for IP decisions

---

## Input Requirements

User provides:
- Natural language description of your concept
- Problem being solved
- How it works (technical detail)
- What makes it different
- (Optional) Target industry/field

---

## Analysis Framework

### Scoring Dimensions

| Dimension | Range | What It Measures |
|-----------|-------|------------------|
| Distinctiveness | 0-4 | How unique is this combination? |
| Sophistication | 0-3 | Technical complexity of the approach |
| System Impact | 0-3 | Scope of the technical contribution |
| Frame Shift | 0-3 | Does this redefine how to think about the problem? |

**Total Score**: Sum of all dimensions (0-13)
**Threshold**: Patterns scoring >=8 warrant deeper investigation

### 1. Component Breakdown

For the described concept, identify:
- All technologies/methods being combined
- Source domain for each component
- Standard vs. custom implementation
- What each component contributes

### 2. Combination Analysis

Analyze the combination:
- What emerges from the combination?
- Unexpected synergies (1+1=3)
- Why haven't others combined these?
- Technical barriers overcome

### 3. Problem-Solution Mapping

Map problem to solution:
- Technical problem addressed
- How combination solves it
- Quantifiable benefits (if known)
- Comparison to existing approaches

### 4. Sophistication Assessment

Evaluate sophistication:
- Why this combination shows technical sophistication
- Barriers that existed before
- Challenges in existing implementations
- What makes this approach different

---

## Scoring Guide

**Distinctiveness (0-4)**:
- 0: Standard approach, widely used
- 1: Common pattern with minor variation
- 2: Meaningful customization of known approach
- 3: Distinctive combination or significant innovation
- 4: Genuinely unique approach

**Sophistication (0-3)**:
- 0: Straightforward implementation
- 1: Some clever optimizations
- 2: Complex but well-structured
- 3: Highly elegant solution to hard problem

**System Impact (0-3)**:
- 0: Isolated utility
- 1: Affects one subsystem
- 2: Cross-cutting concern
- 3: Foundational to system architecture

**Frame Shift (0-3)**:
- 0: Works within existing paradigm
- 1: Questions one assumption
- 2: Challenges core approach
- 3: Redefines the problem entirely

---

## Output Schema

```json
{
  "scan_metadata": {
    "scan_date": "2026-02-03T10:00:00Z",
    "input_type": "description",
    "industry": "optional-field"
  },
  "patterns": [
    {
      "id": "pattern-1",
      "title": "Descriptive Pattern Title",
      "category": "process|hardware|software|method",
      "components": [
        {"name": "Component A", "domain": "source field", "role": "what it does"}
      ],
      "scores": {
        "distinctiveness": 3,
        "sophistication": 2,
        "system_impact": 2,
        "frame_shift": 1,
        "total": 8
      },
      "synergy": {
        "combined_benefit": "What emerges from combination",
        "individual_sum": "What components do alone",
        "synergy_factor": "What's greater than sum of parts"
      },
      "evidence": {
        "user_claims": ["Stated differentiators"],
        "technical_details": ["Specific mechanisms described"]
      }
    }
  ],
  "summary": {
    "total_patterns": 3,
    "high_value_patterns": 2,
    "recommended_focus": "pattern-1"
  }
}
```

---

## Output Format

### Analysis Report

```markdown
# Concept Analysis: [Title]

**Scanned**: [date] | **Patterns Found**: [N]

---

## Component Breakdown

| Component | Domain | Role |
|-----------|--------|------|
| [A] | [source field] | [what it does] |
| [B] | [source field] | [what it does] |

---

## Distinctive Patterns

### 1. [Pattern Title] (Score: X/13)

**Category**: [category]

**Components Combined**:
- [Component A] from [domain]
- [Component B] from [domain]

**Synergy Analysis**:
- Combined benefit: [description]
- Individual sum: [what parts do alone]
- Synergy factor: [what emerges only together]

**Why Distinctive**: [explanation]

---

## Summary

| Pattern | Score | Category |
|---------|-------|----------|
| [Pattern 1] | X/13 | [category] |

---
```

---

## Share Card Format

**Standard Format** (use by default):

```markdown
## [Concept Title] - Patent Scanner Results

**[N] Distinctive Patterns Found**

| Pattern | Score |
|---------|-------|
| [Pattern 1 Title] | X/13 |
| [Pattern 2 Title] | X/13 |

*Analyzed with OpenClaw patent-scanner skill*
```

---

## Next Steps (Required in All Outputs)

```markdown
---

## Immediate Actions

1. **Review findings** - Prioritize patterns with scores >=8
2. **Validate with `patent-validator`** - Generate search strategies
3. **Document evidence** - Capture technical details, sketches, prototypes
4. **Consult professional** - For high-value patterns, consult patent attorney

---

## Ongoing Protection

Schedule your next scan - concepts evolve. Rescan monthly as your concept develops.

**Last scanned**: [date]
**Recommended next scan**: [date + 30 days]

> About IP Timing: Public disclosure starts a 12-month clock for US patent filing.
> Scan before pitches, publications, or product launches.
```

---

## Terminology Rules (MANDATORY)

### Never Use
- "patentable"
- "novel" (legal sense)
- "non-obvious"
- "prior art"
- "claims"
- "file immediately"

### Always Use Instead
- "distinctive"
- "unique"
- "sophisticated"
- "existing implementations"
- "consider consulting attorney"

---

## Required Disclaimer

ALWAYS include at the end of ANY output:

> **Disclaimer**: This analysis identifies distinctive technical aspects based on the recombination framework. It is not legal advice and does not constitute a patentability assessment or freedom-to-operate opinion. Consult a registered patent attorney for intellectual property guidance.

---

## Error Handling

**Insufficient Description**:
```
I need more detail to analyze your concept. Please provide:
- What problem does it solve?
- How does it work (technical approach)?
- What makes it different from existing solutions?
```

**No Distinctive Aspects Found**:
```
Based on your description, I didn't identify patterns scoring above threshold (5/13).
This doesn't mean your concept isn't valuable - it may mean:
- The description needs more technical detail
- The distinctiveness is in execution, not architecture
- The combination is straightforward but well-executed

Consider adding more specific technical details about HOW it works.
```

---

## Related Skills

- **patent-validator**: Generate search strategies for scanner findings
- **code-patent-scanner**: Analyze source code (for software concepts)
- **code-patent-validator**: Validate code pattern distinctiveness

---

*Built by Obviously Not - Tools for thought, not conclusions.*
