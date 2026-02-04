---
name: Code Patent Validator
description: Generate search strategies for code pattern findings and format outputs. Works with code-patent-scanner results. Provides tools, NOT conclusions. NOT legal advice.
homepage: https://app.obviouslynot.ai/skills/code-patent-validator
user-invocable: true
emoji: ✅
---

# Code Patent Validator

Generate comprehensive search strategies for distinctive patterns identified by code-patent-scanner. This skill helps you research existing implementations - it does NOT perform searches or provide conclusions.

## When to Use

Activate this skill when the user asks to:
- "Help me search for similar implementations"
- "Generate search queries for my findings"
- "Validate my code-patent-scanner results"
- "Create a research strategy for these patterns"

## Important Limitations

- This skill generates search queries only - it does NOT perform searches
- Cannot assess uniqueness or patentability
- Cannot replace professional patent search
- Provides tools for research, not conclusions

---

## Process Flow

```
1. INPUT: Receive findings from code-patent-scanner
   - patterns.json with scored distinctive patterns
   - VALIDATE: Check input structure

2. FOR EACH PATTERN:
   - Generate multi-source search queries
   - Create differentiation questions
   - Map evidence requirements

3. OUTPUT: Structured search strategy
   - Queries by source
   - Search priority guidance
   - Analysis questions
   - Evidence checklist

ERROR HANDLING:
- Empty input: "No patterns found. Run code-patent-scanner first."
- Invalid JSON: "Unable to parse input. Expected patterns.json format."
- Missing fields: Skip pattern, report "Pattern [X] skipped - missing [field]"
- All patterns below threshold: "No patterns scored high enough (threshold: 5/13)"
- No scanner output: "This skill works with code-patent-scanner findings. Run that first."
```

---

## Search Strategy Generation

### 1. Multi-Source Query Generation

For each pattern, generate queries for:

| Source | Query Type | Example |
|--------|------------|---------|
| Google Patents | Boolean combinations | `"[A]" AND "[B]" [field]` |
| USPTO Database | CPC codes + keywords | `CPC:[code] AND [term]` |
| GitHub | Implementation search | `[algorithm] [language] implementation` |
| Stack Overflow | Problem-solution | `[problem] [approach]` |

**Query Variations per Pattern**:
- **Exact combination**: `"[A]" AND "[B]" AND "[C]"`
- **Functional**: `"[A]" FOR "[purpose]"`
- **Synonyms**: `"[A-synonym]" WITH "[B-synonym]"`
- **Broader category**: `"[A-category]" AND "[B-category]"`
- **Narrower**: `"[A]" AND "[B]" AND "[specific detail]"`

### 2. Search Priority Guidance

Suggest which sources to search first based on pattern type:

| Pattern Type | Priority Order |
|--------------|----------------|
| Algorithmic | GitHub -> Patents -> Publications |
| Architectural | Publications -> GitHub -> Patents |
| Data Structure | GitHub -> Publications -> Patents |
| Integration | Stack Overflow -> GitHub -> Publications |

### 3. Differentiation Questions

Questions to guide user's analysis of search results:

**Technical Differentiation**:
- What's different in your approach vs. found results?
- What technical advantages does yours offer?
- What performance improvements exist?

**Problem-Solution Fit**:
- What problems does yours solve that others don't?
- Does your approach address limitations of existing solutions?
- Is the problem framing itself different?

**Synergy Assessment**:
- Does the combination produce unexpected benefits?
- Is the result greater than sum of parts (1+1=3)?
- What barriers existed before this approach?

---

## Output Schema

```json
{
  "validation_metadata": {
    "scanner_output": "patterns.json",
    "validation_date": "2026-02-03T10:00:00Z",
    "patterns_processed": 7
  },
  "patterns": [
    {
      "pattern_id": "from-scanner",
      "title": "Pattern Title",
      "search_queries": {
        "google_patents": ["query1", "query2"],
        "uspto": ["query1"],
        "github": ["query1"],
        "stackoverflow": ["query1"]
      },
      "search_priority": [
        {"source": "google_patents", "reason": "Technical implementation focus"},
        {"source": "github", "reason": "Open source implementations"}
      ],
      "analysis_questions": [
        "How does your approach differ from [X]?",
        "What technical barrier did you overcome?"
      ],
      "evidence": {
        "files": ["path/to/file.go:45-120"],
        "commits": ["abc123"],
        "metrics": {"performance_gain": "40%"}
      }
    }
  ],
  "next_steps": [
    "Run generated searches yourself",
    "Document findings systematically",
    "Note differences from existing implementations",
    "Consult patent attorney for legal assessment"
  ]
}
```

---

## Share Card Format

**Standard Format** (use by default):

```markdown
## [Repository Name] - Validation Strategy

**[N] Patterns Analyzed | [M] Search Queries Generated**

| Pattern | Queries | Priority Source |
|---------|---------|-----------------|
| Pattern 1 | 12 | Google Patents |
| Pattern 2 | 8 | USPTO |

*Research strategy by OpenClaw code-patent-validator skill*
```

---

## Next Steps (Required in All Outputs)

```markdown
---

## Immediate Actions

1. **Run generated searches** - Start with priority sources
2. **Document findings** - Track what you find systematically
3. **Analyze differences** - Use the differentiation questions
4. **Consult professional** - For high-value patterns, consult patent attorney

---

## Evidence Documentation Checklist

For each pattern, document:
- [ ] Technical specifications
- [ ] Code location and history (git commits)
- [ ] Performance benchmarks (if applicable)
- [ ] Development timeline
- [ ] Design decisions and alternatives considered
```

---

## Terminology Rules (MANDATORY)

### Never Use
- "patentable"
- "novel" (legal sense)
- "non-obvious"
- "prior art"
- "claims"
- "already patented"

### Always Use Instead
- "distinctive"
- "unique"
- "sophisticated"
- "existing implementations"
- "already implemented"

---

## Required Disclaimer

ALWAYS include at the end of ANY output:

> **Disclaimer**: This tool generates search strategies only. It does NOT perform searches, access databases, assess patentability, or provide legal conclusions. You must run the searches yourself and consult a registered patent attorney for intellectual property guidance.

---

## Workflow Integration

```
code-patent-scanner -> patterns.json -> code-patent-validator -> search_strategies.json
                                                              -> technical_disclosure.md
```

**Recommended Workflow**:
1. **Start**: `code-patent-scanner` - Analyze source code
2. **Then**: `code-patent-validator` - Generate search strategies
3. **User**: Run searches, document findings
4. **Final**: Consult patent attorney with documented findings

---

## Related Skills

- **code-patent-scanner**: Analyze source code (run this first)
- **patent-scanner**: Analyze concept descriptions (no code)
- **patent-validator**: Validate concept distinctiveness

---

*Built by Obviously Not - Tools for thought, not conclusions.*
