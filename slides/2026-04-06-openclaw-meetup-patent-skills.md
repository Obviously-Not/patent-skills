# Patent Skills for AI Agents: Discover What Makes Your Code Distinctive

## OpenClaw Meetup | April 2026

**Slides:** [View on Gamma](https://gamma.app/docs/Patent-Skills-for-AI-Agents-o0hwxd7kyc1pj6u)

---

# The Problem

Every developer writes code they think is unique

But how do you know if your approach is *actually* distinctive?

**Traditional path:**
- Hire a patent attorney ($500+/hour)
- Hope they understand your code
- Wait weeks for analysis

**What if AI agents could help you discover distinctive patterns first?**

---

# Our Story

My brother and I are developers in Alaska. Two-person team. No VC. No legal department.

We kept building things we thought were novel — but had no way to know for sure. The traditional path ($500/hour attorneys, weeks of waiting) wasn't realistic for us.

So we built these skills for ourselves first. Used them to identify distinctive patterns in our own work. Filed several provisional patents. **It worked.**

Then we realized: every small team, every solo inventor, every garage startup faces this same asymmetry.

**Big corporations** have patent portfolios, in-house counsel, and prior art databases.

**The rest of us** have... nothing. Until now.

---

# Introducing: Patent Skills

Four open-source skills for discovering and validating distinctive technical patterns

**github.com/Obviously-Not/patent-skills** | Built on the Agent Skills standard

| Skill | Downloads | Stars | What It Does |
|-------|-----------|-------|-------------|
| **code-patent-scanner** | 2.5k | ⭐13 | Scan codebases for distinctive patterns |
| **code-patent-validator** | 2.1k | ⭐14 | Generate search strategies for code findings |
| **patent-scanner** | 2.5k | ⭐14 | Analyze concept descriptions |
| **patent-validator** | 2.1k | ⭐14 | Generate search strategies for concepts |

**Workflow:** Scanner → Validator → Research → Attorney

---

# Why Skills, Not Apps?

**Skills are portable instructions** that any AI agent can read and execute

```
code-patent-scanner/
  SKILL.md           ← Instructions + metadata
  scripts/           ← Optional automation
  references/        ← Detailed docs
```

**Privacy by design:**
- Runs on YOUR model, YOUR machine
- No code uploaded to third-party servers
- Pre-disclosure protection (patent bar date matters!)
- You control what leaves your laptop

**Write once, use everywhere:**
- Claude Code, OpenClaw, Cursor, Gemini CLI, GitHub Copilot
- Or just paste the SKILL.md into any LLM chat!

---

# Deep Dive: code-patent-scanner

**clawhub install leegitw/code-patent-scanner**

⭐ 13 stars | 📊 2.5k downloads | MIT-0 license

✅ Security scanned: VirusTotal + OpenClaw (Benign, High Confidence)

**What it does:**
- Scans your codebase for technically distinctive patterns
- Provides structured 4-dimension scoring (0-13 scale)
- Generates problem-solution-benefit mappings
- Suggests claim angles for high-scoring patterns

**What it doesn't do:**
- Search patent databases
- Make legal determinations
- Replace professional counsel

---

# The Scoring Framework

Four dimensions, each scored independently:

| Dimension | Range | What It Measures |
|-----------|-------|------------------|
| **Distinctiveness** | 0-4 | How unique vs standard approaches |
| **Sophistication** | 0-3 | Engineering complexity and elegance |
| **System Impact** | 0-3 | Effect on overall system behavior |
| **Frame Shift** | 0-3 | Reframes the problem itself |

**Threshold:** Only patterns scoring ≥8/13 are reported

---

# The Abstraction Principle

> "I don't need to see the code to draft claims. I need to understand what the invention IS."

**Key insight: Lossy abstraction is a feature, not a bug**

| Skip This | Use This Instead |
|-----------|-----------------|
| "uses bcrypt.compare()" | "applies cryptographic one-way function" |
| "stores in PostgreSQL" | "persists to durable storage" |
| "sends HTTP POST" | "transmits data via network protocol" |

Broader descriptions = harder to design around

---

# Live Demo Time

Let's scan OpenClaw's own security code

**Target:** `openclaw/openclaw` → `src/security/`

**Why this target?**
- 25 source files (perfect size)
- Real production security code
- Should have interesting patterns
- Meta: scanning OpenClaw at an OpenClaw meetup 🦞

---

# Demo: Live via Telegram + OpenClaw

```
🔬 Code Patent Scanner — openclaw/src/security

Repository: github.com/openclaw/openclaw
Directory: src/security
Files analyzed: 14 non-test source files
Scan date: 2026-04-06
```

Installed the skill via Telegram, pointed it at OpenClaw's own repo 🦞

---

# Demo Result: 12/13 🔥

| Pattern | Score | Category |
|---------|-------|----------|
| Cryptographically-Bound External Content Wrapping with Unicode Homoglyph Normalization | **12/13** 🔥 | architectural |

**Highest possible score is 13** — this pattern hit 12!

---

# Full Scan Results: 4 Patterns Found

| Pattern | Score | Signals |
|---------|-------|---------|
| Cryptographic Boundary + Unicode Normalization | **12/13** 🔥 | 🟢 Market 🟢 Competitive 🟢 Novelty |
| Static ReDoS Detection (Frame Analysis) | **10/13** | 🟡 Market 🟢 Competitive 🟡 Novelty |
| Separated Command Gate / DM Policy | **9/13** | 🟡 Market 🟢 Competitive 🟡 Novelty |
| Triple-Key Scan Cache | **8/13** | ⚪ Market 🟡 Competitive 🟡 Novelty |

All 4 patterns scored ≥8/13 (the reporting threshold)

The scanner surfaces patterns across categories: **architectural**, **algorithmic**, **data structure**

---

# The 12/13 Pattern: Prompt Injection Defense

**Cryptographically-Bound External Content Wrapping with Unicode Homoglyph Normalization**

**The Problem:**
Standard content wrapping uses static boundary strings that attackers can embed in external content to escape the untrusted zone and inject instructions.

**The Solution:**
A prompt injection defense system that generates per-message cryptographic boundary markers (randomBytes) AND applies multi-layer Unicode normalization — covering fullwidth ASCII, CJK angle brackets, mathematical brackets, invisible zero-width characters, and soft hyphens.

**Why 12/13?**
- Distinctiveness: 4/4 — genuinely unique combination
- Sophistication: 3/3 — elegant multi-layer defense
- System Impact: 3/3 — foundational to LLM security
- Frame Shift: 2/3 — redefines content boundary problem

**File:** `external-content.ts`

---

# The Output: Share Card Format

```markdown
## openclaw/openclaw - Code Patent Scanner Results

**🔥 High-Value Pattern Found**

| Pattern | Score | Signals |
|---------|-------|---------|
| Cryptographically-Bound Content Wrapping | 12/13 🔥 | 🟢🟢🟢 |

*Analyzed with code-patent-scanner from obviouslynot.ai*
```

A 12/13 score means: **take this to a patent attorney**

---

# What Happens Next?

**Step 1: Review**
Prioritize patterns scoring ≥8

**Step 2: Validate**
Run `code-patent-validator` to generate search strategies

**Step 3: Research**
Use those strategies to find existing implementations

**Step 4: Consult**
Bring findings to a patent attorney with evidence documented

---

# Important Disclaimer

These skills identify **distinctive technical patterns**, not patentability

**What they DO:**
✅ Identify unique combinations of techniques
✅ Generate search queries for research
✅ Document technical evidence
✅ Suggest claim angles

**What they DON'T:**
❌ Search patent databases
❌ Assess legal patentability
❌ Replace professional counsel

Always consult a registered patent attorney for IP guidance

---

# Try It Yourself

**Clone from GitHub (open source):**
```bash
git clone https://github.com/Obviously-Not/patent-skills
cp -r patent-skills/code-patent-scanner ~/.claude/skills/
```

**Or install via ClawHub:**
```bash
clawhub install leegitw/code-patent-scanner
```

**Or just paste SKILL.md into any LLM chat!**

---

# Patent Skills: Quick Reference

| Start Here | Then Use |
|------------|----------|
| **code-patent-scanner** (for code) | code-patent-validator |
| **patent-scanner** (for concepts) | patent-validator |

**Links:**
- GitHub: github.com/Obviously-Not/patent-skills
- ClawHub: clawhub.ai (search @leegitw)
- Website: obviouslynot.ai/skills

---

# Q&A

**Lee Brown** | @leegitw
**Lucas Brown** | @lucasgeeksinthewood

🦞 Obviously Not — Tools for thought, not conclusions

---

# An Invitation

What if every developer could discover what makes their work distinctive?

For too long, IP protection has been a game only well-funded companies could play. The garage inventor, the two-person startup, the solo developer — they create genuinely novel work, but can't afford to find out what they have.

**That asymmetry isn't inevitable. It's a problem we can solve together.**

The tools are open source. The methodology is documented. The barrier is gone.

I'm looking for collaborators:
- **Developers** — scan your repos, file issues, submit PRs
- **Patent attorneys** — help refine the scoring methodology
- **AI builders** — integrate these skills into your agents

**Fork it. Break it. Make it better.**

github.com/Obviously-Not/patent-skills

*The best ideas survive when they're shared.*

---

# Appendix: Full Skill Portfolio

**@leegitw on ClawHub** — 45+ skills across 7 categories

**Soul & Identity:**
- neon-soul (1.8k ⭐13) — Automated soul synthesis for AI agents
- consciousness-soul-identity (1.1k ⭐15) — Discover identity through memory

**Principle-Based Decomposition (PBD):**
- essence-distiller (2.4k ⭐6) — Find what actually matters
- pbe-extractor (1.6k ⭐6) — Extract invariant principles
- core-refinery (1.5k) — Find the core across sources
- pattern-finder (1.5k ⭐3) — Discover agreement between sources
- golden-master (1.3k ⭐4) — Track source-of-truth relationships
- principle-comparator (1.1k) — Compare two sources
- principle-synthesizer (1.1k) — Synthesize from 3+ sources

**Agent Safety & Governance:**
- failure-memory (780) — Turn failures into prevention patterns
- context-verifier (714 ⭐8) — Verify file integrity before editing
- safety-checks (606 ⭐7) — Model pinning and runtime validation
- review-orchestrator (323 ⭐7) — Multi-perspective code review
- agentic-governance (285 ⭐5) — Constraint lifecycle management
- constraint-engine (252 ⭐5) — Learn from consequences
- workflow-tools (224 ⭐7) — Loop detection and parallel decisions

**Creative Skills:**
- visual-concept (643) — Technical insights → visual guides
- side-quests (369 ⭐3) — Song + visual + TED talk synthesis
- ted-talk (368 ⭐3) — Technical insights → TED-style talks
- insight-song (331) — Technical insights → Suno-ready songs
- ethics-guardrails (312 ⭐2) — Publish ethical guardrails
- song-remix (305) — Twin Remix method for songs

**Animal House Pets (19 skills):**
- Virtual pets for AI agents — 73+ species, real hunger, permanent death
- Model-specific: gemini-pet, copilot-pet, llama-pet, qwen-pet
- Buddy variants: anthropic-buddy, dev-pet, terminal-tamagotchi
- Species: thornfox, bramblebear, gustowl, cloudferret, and more

---

## Usage Notes

**For gamma.app:**
1. Select "Paste" when creating a new presentation
2. Paste all the content above (excluding this Usage Notes section)
3. Choose "Presentation" format
4. Pick a theme (suggest: modern/minimal to match technical content)
5. Gamma will create ~20 cards from the headers

**Sources:**
- [ClawHub Skill Registry](https://clawhub.ai/)
- [Agent Skills Standard](https://agentskills.io/home)
- [Anthropic Agent Skills Announcement](https://venturebeat.com/technology/anthropic-launches-enterprise-agent-skills-and-opens-the-standard)
- [OpenClaw Skills Ecosystem](https://github.com/VoltAgent/awesome-openclaw-skills)
