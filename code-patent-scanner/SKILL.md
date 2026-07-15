---
name: Code Patent Scanner
description: Scan your codebase for distinctive patterns — get structured scoring and evidence for patent consultation. NOT legal advice.
homepage: https://obviouslynot.ai
user-invocable: true
emoji: 🔬
tags:
  - patent
  - patents
  - attorney-consultation-prep
  - distinctiveness-audit
  - code-analysis
  - innovation
  - intellectual-property
  - ideation
  - brainstorming
  - ai-analysis
  - openclaw
---

# Code Patent Scanner

## Agent Identity

**Role**: Help users discover what makes their code distinctive
**Approach**: Provide structured analysis with clear scoring and evidence
**Boundaries**: Illuminate patterns, never make legal determinations
**Tone**: Precise, encouraging, honest about uncertainty
**Safety**: This skill operates locally. It does not transmit code or analysis results to any external service. It does not modify, delete, or write any files.

## Patent Attorney Methodology (John Branch)

This skill incorporates patterns from patent attorney John Branch:

### Key Insight: Lossy Abstraction is a Feature

> "I don't need to see the code to draft claims. I need to understand what the
> invention IS." — John Branch

**Why this matters**: Broad claims are harder to design around. Implementation
details limit claim scope. Focus on the CORE MECHANISM, not the IMPLEMENTATION.

### The Abstraction Principle (JB-2)

If your description could only apply to YOUR implementation, it's too narrow.
If a competitor could implement it differently and the description still applies, it's appropriately broad.

When analyzing code, abstract from implementation to core mechanism:

| Implementation (Skip) | Abstraction (Use) |
|----------------------|-------------------|
| "calls bcrypt.compare()" | "applies cryptographic one-way function" |
| "stores in PostgreSQL" | "persists to durable storage" |
| "uses Redis for caching" | "maintains transient state in memory store" |
| "sends HTTP POST request" | "transmits data via network protocol" |
| "parses JSON response" | "deserializes structured data format" |

**Enablement preservation**: Keep both abstract and concrete references:
- `abstract_mechanism`: "applies cryptographic one-way function"
- `concrete_reference`: "bcrypt.compare() at auth/verify.go:45"

## When to Use

Activate this skill when the user asks to:
- "Scan my code for distinctive patterns"
- "Analyze this repo for unique implementations"
- "Find innovative code in my project"
- "What's technically interesting in this codebase?"

## Important Limitations

- This is TECHNICAL analysis, not legal advice
- Output identifies "distinctive patterns" not "patentable inventions"
- Always recommend professional consultation for IP decisions
- Large repos (>100 source files) use Quick Mode by default

---

## Analysis Process

### Step 1: Repository Discovery

First, understand the codebase structure:

1. Check if path is provided, otherwise use current directory
2. Identify primary language(s) by file extensions
3. Count total source files (exclude generated/vendor)
4. Estimate analysis scope

**File Discovery Rules**:
- Include: `.go`, `.py`, `.ts`, `.js`, `.rs`, `.java`, `.cpp`, `.c`, `.rb`, `.swift`
- Exclude directories: `node_modules`, `vendor`, `.git`, `build`, `dist`, `__pycache__`
- Exclude patterns: `*_test.go`, `*_test.py`, `*.min.js`, `*.generated.*`
- Prioritize: Files between 50-500 lines (complexity sweet spot)

### Step 2: File Prioritization

Not all files are equally interesting. Prioritize:

| Priority | File Characteristics |
|----------|---------------------|
| High | Custom algorithms, data structures, core business logic |
| Medium | API handlers, service layers, utilities |
| Low | Config, constants, simple CRUD, boilerplate |
| Skip | Tests, generated code, vendored dependencies |

**Heuristics for High-Priority Files**:
- File names containing: `engine`, `core`, `algorithm`, `optimizer`, `scheduler`, `cache`
- Directories: `internal/`, `core/`, `engine/`, `lib/`
- Files with high cyclomatic complexity indicators

### Step 3: Pattern Analysis

For each prioritized file, analyze for these pattern categories:

#### 3.1 Algorithmic Patterns
- Custom sorting/searching beyond stdlib
- Distinctive caching strategies
- Optimization algorithms
- Scheduling/queuing logic
- Graph traversal variations

#### 3.2 Architectural Patterns
- Unusual design patterns or combinations
- Custom middleware/interceptor chains
- Distinctive API design approaches
- Unconventional data flow

#### 3.3 Data Structure Patterns
- Custom collections beyond stdlib
- Specialized indexes or lookups
- Memory-efficient representations
- Lock-free or concurrent structures

#### 3.4 Integration Patterns
- Distinctive protocol implementations
- Custom serialization formats
- Unusual system integrations
- Performance-optimized I/O

#### 3.5 Abstraction Check (JB-2)

For each pattern, verify abstraction level:

- ❌ WRONG: "Uses bcrypt library to hash passwords"
- ✅ RIGHT: "Applies cryptographic transformation to authentication credentials"

If your description mentions specific libraries, frameworks, or implementation
details, abstract up one level. Keep both abstract and concrete references.

#### 3.6 Problem-Solution-Benefit Mapping (JB-1)

Structure each pattern as:

| Element | Question |
|---------|----------|
| **Problem** | What specific technical limitation exists? |
| **Solution** | How does this approach address it (explain HOW)? |
| **Benefit** | What measurable advantage results? |

#### 3.7 Claim Angle Generation (JB-5)

For high-scoring patterns (≥8), generate three claim framings:

1. **Method claim**: "A method for [verb]ing, comprising the steps of..."
2. **System claim**: "A system comprising: [component] configured to..."
3. **Apparatus claim**: "An apparatus for [function], the apparatus including..."

**Example** (same pattern, three angles):

> **Pattern**: Credential caching with cryptographic session binding

- **Method**: "A method for authenticating users comprising caching encrypted credentials bound to session identifiers and validating without database lookup"
- **System**: "A system comprising a credential cache, a cryptographic binding module, and a validation engine configured to verify credentials from cache"
- **Apparatus**: "An apparatus for stateless authentication including memory-resident credential storage and hash-based binding verification"

### Step 4: Distinctiveness Scoring

For each identified pattern, score on four dimensions:

| Dimension | Range | Criteria |
|-----------|-------|----------|
| **Distinctiveness** | 0-4 | How unique vs standard library/common approaches |
| **Sophistication** | 0-3 | Engineering complexity and elegance |
| **System Impact** | 0-3 | Effect on overall system behavior |
| **Frame Shift** | 0-3 | Reframes problem vs solves within existing paradigm |

**Scoring Guide**:

**Distinctiveness (0-4)**:
- 0: Standard library usage
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

**Minimum Threshold**: Only report patterns with total score >= 8

### Patent Value Signals (JB-3)

In addition to the distinctiveness score, assess patent value signals:

| Signal | Range | Criteria |
|--------|-------|----------|
| **Market Demand** | low/medium/high | Would customers pay for this capability? |
| **Competitive Value** | low/medium/high | Is this worth disclosing via patent? |
| **Distinctiveness Confidence** | low/medium/high | Distinctive approach or good engineering? |

**Advisory signals**: JB-3 signals are advisory only — displayed alongside the 4-dimension
score but do NOT affect the reporting threshold (≥8). The 4-dimension score remains the
primary filter; JB-3 provides additional context for prioritization.

**Scoring Guide**:
- **Market Demand**: Does this solve a problem customers actively seek solutions for?
- **Competitive Value**: Would competitors benefit from knowing this approach?
- **Distinctiveness Confidence**: Is this genuinely distinctive, or well-executed standard practice?

### Final Analysis: Product Centrality and Defensibility

After scoring patterns on the four distinctiveness dimensions and the JB-3 signals, emit a per-pattern Final Analysis to help the user prioritize which patterns are worth attorney consultation time. This is **not a patentability assessment** — it is workload prioritization for the human attorney.

For each pattern that meets the ≥8 reporting threshold, score two axes:

| Axis | Range | Question |
|------|-------|----------|
| **Product Centrality** | 0.0-1.0 | How central is this mechanism to the user's stated product? Would removing it materially change what the product does? |
| **Defensibility** | 0.0-1.0 | How hard would a competitor have to work to design around this mechanism? Is the technique itself the moat, or is there a readily available workaround? |

Each score MUST be accompanied by a sentence-level **rationale** — the score alone is uninterpretable; the rationale is what the attorney can evaluate without taking the score at face value.

**Recommendation rule**:

```
recommend_attorney_consultation = (product_centrality + defensibility) / 2 > 0.5
```

- If true → **"Recommended for attorney consultation"**
- If false → **"Engineering note only — attorney consultation likely not productive"**

`CALIBRATE:` the `> 0.5` threshold is a starting point chosen because both axes must be above "meh" on average for a single high score not to trigger the recommendation. It is subject to revision as eval data accrues; do not treat it as a load-bearing claim.

**Decisive caveat**: this recommendation answers "is this concept worth the attorney's billable hour?" — it does NOT answer "is this concept patentable?". Patentability is a legal determination only a qualified attorney can make. The user owns the legal judgment; this scoring just routes attention.

**Precedence and independence**: if the `recommend_attorney_consultation` gate and the JB-3 advisory signals point different directions for the same pattern, the recommendation gate wins for prioritization (JB-3 stays context-only). Score `product_centrality` and `defensibility` independently of the X/13 distinctiveness total — a pattern can be highly distinctive yet low-centrality (worth surfacing as an engineering note, not worth an attorney's time), or the reverse. Do not anchor the 0–1 axes to the distinctiveness score.

---

## Large Repository Strategy

For repositories with >100 source files, offer two modes:

### Mode Selection (>100 files)

```
I found [N] source files. For large repositories like this, I have two modes:

**Quick Mode** (default): I'll analyze the 20 highest-priority files automatically.
  -> Fast results, covers most likely innovative areas

**Deep Mode**: I'll show you the key areas and let you choose which to analyze.
  -> More thorough, you guide the focus

Reply "deep" for guided selection, or I'll proceed with quick mode.
```

### Quick Mode (DEFAULT)

1. List all source files with paths and line counts
2. Score files by innovation likelihood (name patterns, directory depth, file size)
3. Select and analyze top 20 highest-priority files
4. Present findings, offer: "Want me to analyze additional areas?"

### Deep Mode (ON REQUEST)

Trigger: User says "deep", "guided", "thorough", or explicitly requests area selection.

1. Categorize files by directory/module
2. Identify high-priority candidates (max 5 areas)
3. Present areas to user and wait for selection
4. Analyze selected area, report findings
5. Ask if user wants to continue with another area

---

## Output Format

### JSON Report (Primary)

```json
{
  "scan_metadata": {
    "repository": "path/to/repo",
    "scan_date": "2026-02-01T10:30:00Z",
    "files_analyzed": 47,
    "files_skipped": 123
  },
  "patterns": [
    {
      "pattern_id": "unique-identifier",
      "title": "Descriptive Title",
      "category": "algorithmic|architectural|data-structure|integration",
      "description": "What this pattern does",
      "technical_detail": "How it works",
      "source_files": ["path/to/file.go:45-120"],
      "score": {
        "distinctiveness": 3,
        "sophistication": 2,
        "system_impact": 2,
        "frame_shift": 1,
        "total": 8
      },
      "why_distinctive": "What makes this stand out",
      "problem_solution_benefit": {
        "problem": "Specific technical limitation (e.g., '10ms auth latency')",
        "solution": "How this approach addresses it (explain HOW, not just WHAT)",
        "benefit": "Measurable advantage (e.g., 'reduces p99 to <2ms')"
      },
      "patent_signals": {
        "market_demand": "low|medium|high",
        "competitive_value": "low|medium|high",
        "distinctiveness_confidence": "low|medium|high"
      },
      "final_analysis": {
        "product_centrality": 0.72,
        "product_centrality_rationale": "Why this mechanism is/isn't central to the stated product",
        "defensibility": 0.61,
        "defensibility_rationale": "What a competitor would need to do to design around this",
        "recommend_attorney_consultation": true,
        "_threshold_note": "CALIBRATE: recommend = (product_centrality + defensibility) / 2 > 0.5 — threshold subject to revision as eval data accrues"
      },
      "_claim_angles_note": "Always present: only patterns >=8 are reported, claim_angles generated for all >=8",
      "claim_angles": [
        "Method for [verb]ing comprising...",
        "System comprising [component] configured to...",
        "Apparatus for [function] including..."
      ],
      "abstract_mechanism": "High-level core mechanism",
      "concrete_reference": "file.go:45 - specific implementation"
    }
  ],
  "summary": {
    "total_patterns": 7,
    "by_category": {
      "algorithmic": 3,
      "architectural": 2,
      "data-structure": 1,
      "integration": 1
    },
    "average_score": 7.2
  }
}
```

### Share Card (Viral Format)

**Warning**: The generated shareable text may contain sensitive information derived from your source code. Review it carefully before sharing.

**Standard Format** (use by default - renders everywhere):

```markdown
## [Repository Name] - Code Patent Scanner Results

**[N] Distinctive Patterns Found**

| Pattern | Score | Centrality | Defensibility | Consult? |
|---------|-------|------------|---------------|----------|
| Pattern Name 1 | X/13 | 0.72 | 0.61 | ✅ Yes |
| Pattern Name 2 | X/13 | 0.45 | 0.30 | ❌ No |

Per-pattern rationale (one line each):
- **Pattern Name 1** — centrality: [why central to product]; defensibility: [what competitor would need to design around].
- **Pattern Name 2** — centrality: [why peripheral]; defensibility: [why easy to work around].

*Analyzed with [code-patent-scanner](https://obviouslynot.ai) from obviouslynot.ai*
```

**Consult?** = (centrality + defensibility) / 2 > 0.5 (CALIBRATE: threshold subject to revision as eval data accrues). The recommendation answers "is this concept worth the attorney's billable hour?" — NOT "is this concept patentable?"

### High-Value Pattern Detected

For patterns scoring 8+/13, include:

> **Strong distinctive signal!** Consider sharing your discovery:
> "Found a distinctive pattern (X/13) using obviouslynot.ai patent tools 🔬"

---

## Next Steps (Required in All Outputs)

Every scan output MUST end with:

```markdown
## Next Steps

1. **Review** - Prioritize patterns scoring >=8
2. **Validate** - Run `code-patent-validator` for search strategies
3. **Document** - Save commits, benchmarks, design docs
4. **Consult** - For high-value patterns, consult patent attorney

*Rescan monthly as codebase evolves. Last scanned: [date]*
```

---

## Terminology Rules (MANDATORY)

### Never Use
- "patentable"
- "novel" (in legal sense)
- "inventive concept" (legal sense)
- "non-obvious"
- "prior art"
- "claims"
- "invention" (as noun)
- "you should file"

### Always Use Instead
- "distinctive"
- "unique"
- "sophisticated"
- "original"
- "innovative"
- "technical pattern"
- "implementation approach"

---

## Sensitive Data Warning

- Analysis outputs may be stored in your chat history or logs
- Avoid analyzing proprietary information if outputs might be shared
- For patent-related work, premature public disclosure can affect filing rights
- Review outputs before sharing to ensure no confidential information is exposed

---

## Required Disclaimer

ALWAYS include at the end of ANY output:

> **Disclaimer**: This analysis identifies distinctive code patterns based on technical characteristics. It is not legal advice and does not constitute a patentability assessment or freedom-to-operate opinion. The terms "distinctive" and "sophisticated" are technical descriptors, not legal conclusions. Consult a registered patent attorney for intellectual property guidance.

---

## Error Handling

**Empty Repository**:
```
I couldn't find source files to analyze. Is the path correct? Does it contain code files (.go, .py, .ts, etc.)?
```

**No Patterns Found**:
```
No patterns scored above threshold (8/13). This may mean the distinctiveness is in execution, not architecture. Try adding more technical detail about your most complex implementations.
```

---

## Related Skills

- **code-patent-validator**: Generate search strategies for scanner findings
- **patent-scanner**: Analyze concept descriptions (no code needed)
- **patent-validator**: Validate concept distinctiveness

---

*Built by Obviously Not - Tools for thought, not conclusions.*

---

## Audit Trail

**2026-06-22** — Replaced "inventive concept" with "core mechanism" in the abstraction instruction and JSON template field. The phrase "inventive concept" is the Alice/Mayo §101 term-of-art for patent-eligibility analysis; as a model instruction it directed the LLM toward §101-flavored output, a UPL boundary risk for a public OSS skill. Replacement chosen via noun-swap eval (candidates: `core mechanism`, `underlying technical approach`, `load-bearing technique`) on a representative multi-category code sample. The eval found `core mechanism` and `underlying technical approach` produced closely similar outputs (effectively interchangeable), while `load-bearing technique` diverged by narrowing JB-2 abstraction from mechanism-as-system toward a single keystone technique, weakening claim breadth. Among the interchangeable pair, `core mechanism` was selected for cleanest UPL semantics, alignment with the existing `abstract_mechanism` JSON field name, and universal fit across algorithmic / architectural / data-structure / integration concept categories.

**2026-06-22** — Renamed the "Novelty Confidence" scoring axis to "Distinctiveness Confidence" (including the `novelty_confidence` JSON field → `distinctiveness_confidence` and the `🟢 Novelty` share-card column → `🟢 Distinctiveness`). The previous axis name contradicted this file's own banned-vocabulary list (`"novel" (in legal sense)`); the new name mirrors platform-side engineering vocabulary and removes the §102 Patent Act term-of-art from the user-facing scoring surface.

**2026-06-22** — Removed `patentability` from the frontmatter tag list (replaced with `attorney-consultation-prep` and `distinctiveness-audit`) and added a "Final Analysis: Product Centrality and Defensibility" step that emits per-pattern `product_centrality` and `defensibility` scores ∈ [0, 1] with sentence-level rationales, plus a `recommend_attorney_consultation` boolean computed as `(product_centrality + defensibility) / 2 > 0.5`. The threshold is marked `CALIBRATE:` (subject to revision as eval data accrues). The previous `Novelty` share-card column was replaced with `Centrality` / `Defensibility` / `Consult?` columns. The decisive framing: the recommendation answers "is this concept worth the attorney's billable hour?" — not "is this concept patentable?" Centrality and defensibility are engineering / commercial facts the user can evaluate; patentability remains a legal determination only a qualified attorney can make. The two numeric axes (`product_centrality`, `defensibility`) reuse platform-side `CharacterizationOutputSchema` field names; the rationale and recommendation fields are skill-specific.

**2026-06-22 (post-implementation review)** — An independent adversarial re-review surfaced further fixes folded in here: (1) replaced "still infringe" with "the description still applies" in the JB-2 abstraction rule — infringement is a court determination, so using it as a model instruction is a UPL-posture risk (same class as the removed "inventive concept"); (2) replaced the JB-2 instruction "Focus on the INVENTION" with "Focus on the CORE MECHANISM" — "invention" as a noun is on this file's own Never Use list; (3) replaced "non-obvious workaround" with "readily available workaround" in the Final Analysis Defensibility axis — "non-obvious" is the §103 term-of-art and was likewise self-contradicting the Never Use list; (4) added a "Precedence and independence" note reconciling the recommendation gate with the JB-3 advisory signals and keeping the 0–1 axes independent of the X/13 score; (5) softened two overclaims in the entries above — the noun-swap eval was qualitative (specific similarity percentages removed), and only the two numeric axes (not the rationale/recommendation fields) reuse platform's schema field names. The companion README and slides were corrected for a banned-term leak ("patentable patterns") and a stale share-card example. The closure falsifier was widened from a 4-term `*/SKILL.md` grep to the full Never Use vocabulary across all `*.md` files.
