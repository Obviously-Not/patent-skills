package tests

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// UPL terms that should never appear in skill descriptions
var uplTerms = []string{
	"patentable",
	"non-obvious",
	"prior art",
	"file a patent",
	"patent claims",
}

// TestSkillDescriptions_NoUPLTerms verifies skills don't use legal terminology
func TestSkillDescriptions_NoUPLTerms(t *testing.T) {
	// TR-1: Use ../ prefix since tests run from tests/ subdirectory
	skills, err := filepath.Glob("../*/SKILL.md")
	if err != nil {
		t.Fatalf("Failed to glob skills: %v", err)
	}

	for _, skillPath := range skills {
		content, err := os.ReadFile(skillPath)
		if err != nil {
			t.Errorf("Failed to read %s: %v", skillPath, err)
			continue
		}

		contentLower := strings.ToLower(string(content))
		for _, term := range uplTerms {
			if strings.Contains(contentLower, term) {
				// Allow terms in "Never Use" sections
				if !strings.Contains(contentLower, "never use") {
					t.Errorf("%s contains UPL term: %q", skillPath, term)
				}
			}
		}
	}
}

// TR-5: PBD skills require 6-field Agent Identity pattern
var pbdSkills = map[string]bool{
	"pbe-extractor": true, "essence-distiller": true,
	"principle-comparator": true, "pattern-finder": true,
	"principle-synthesizer": true, "core-refinery": true,
	"golden-master": true,
}

// TestSkillMD_AgentIdentityPattern verifies Agent Identity section exists
// TR-5: PBD skills get 6-field check, patent skills get 4-field check
func TestSkillMD_AgentIdentityPattern(t *testing.T) {
	skills, err := filepath.Glob("../*/SKILL.md")
	if err != nil {
		t.Fatalf("Failed to glob skills: %v", err)
	}

	baseFields := []string{"Role", "Approach", "Boundaries", "Tone"}
	pbdFields := []string{"Role", "Understands", "Approach", "Boundaries", "Tone", "Opening Pattern"}

	for _, skillPath := range skills {
		skillName := filepath.Base(filepath.Dir(skillPath))
		t.Run(skillName, func(t *testing.T) {
			content, err := os.ReadFile(skillPath)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", skillPath, err)
			}
			contentStr := string(content)

			if !strings.Contains(contentStr, "## Agent Identity") {
				t.Errorf("%s missing Agent Identity section", skillPath)
				return
			}

			// TR-5: Use 6-field pattern for PBD skills
			requiredFields := baseFields
			if pbdSkills[skillName] {
				requiredFields = pbdFields
			}

			for _, field := range requiredFields {
				pattern := regexp.MustCompile(`\*\*` + field + `\*\*:`)
				if !pattern.MatchString(contentStr) {
					t.Errorf("%s missing Agent Identity field: %s", skillPath, field)
				}
			}
		})
	}
}

// TestSkillMD_RequiredSections verifies required sections exist
// TR-6: Include "## Required Disclaimer" in required sections
func TestSkillMD_RequiredSections(t *testing.T) {
	skills, err := filepath.Glob("../*/SKILL.md")
	if err != nil {
		t.Fatalf("Failed to glob skills: %v", err)
	}

	// TR-6: Added "## Required Disclaimer"
	requiredSections := []string{
		"## Agent Identity",
		"## When to Use",
		"## Required Disclaimer",
	}

	for _, skillPath := range skills {
		skillName := filepath.Base(filepath.Dir(skillPath))
		t.Run(skillName, func(t *testing.T) {
			content, err := os.ReadFile(skillPath)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", skillPath, err)
			}

			for _, section := range requiredSections {
				if !strings.Contains(string(content), section) {
					t.Errorf("%s missing section: %s", skillPath, section)
				}
			}
		})
	}
}

// TR-3: Voice pairs should have matching methodology sections
func TestSkillMD_VoicePairConsistency(t *testing.T) {
	voicePairs := []struct {
		technical      string
		conversational string
	}{
		{"pbe-extractor", "essence-distiller"},
		{"principle-comparator", "pattern-finder"},
		{"principle-synthesizer", "core-refinery"},
	}

	for _, pair := range voicePairs {
		t.Run(pair.technical+"_"+pair.conversational, func(t *testing.T) {
			convContent, err := os.ReadFile("../" + pair.conversational + "/SKILL.md")
			if err != nil {
				t.Skipf("SKILL.md not found for %s", pair.conversational)
				return
			}

			// Conversational skills should document voice differences
			if !strings.Contains(string(convContent), "Voice Differences") {
				t.Errorf("conversational skill %s should document voice differences from %s",
					pair.conversational, pair.technical)
			}
		})
	}
}

// TR-4: Verify checksum normalization for golden-master skill
func TestGoldenMaster_ChecksumNormalization(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "CRLF to LF conversion",
			input:    "line1\r\nline2\r\n",
			expected: "line1\nline2\n",
		},
		{
			name:     "trailing whitespace trimmed",
			input:    "line1   \nline2\t\n",
			expected: "line1\nline2\n",
		},
		{
			name:     "golden-master metadata stripped",
			input:    "# Title\n<!-- golden-master:source checksum:abc123 -->\nContent here",
			expected: "# Title\n\nContent here",
		},
		{
			name:     "golden-master derived metadata stripped",
			input:    "<!-- golden-master:derived source:file.md source_checksum:abc123 derived_at:2026-02-04 -->\nContent",
			expected: "\nContent",
		},
		{
			name:     "combined normalization",
			input:    "# Doc\r\n<!-- golden-master:source checksum:xyz789 -->  \r\nBody text   \r\n",
			expected: "# Doc\n\nBody text\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizeForChecksum(tc.input)
			if result != tc.expected {
				t.Errorf("normalizeForChecksum() mismatch\ninput:    %q\nexpected: %q\ngot:      %q",
					tc.input, tc.expected, result)
			}
		})
	}
}

// normalizeForChecksum applies normalization rules from golden-master/SKILL.md
func normalizeForChecksum(content string) string {
	// Step 1: CRLF -> LF
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")

	// Step 2: Strip golden-master metadata comments
	for {
		start := strings.Index(content, "<!--")
		if start == -1 {
			break
		}
		afterStart := content[start:]
		checkLen := len(afterStart)
		if checkLen > 50 {
			checkLen = 50
		}
		if !strings.Contains(afterStart[:checkLen], "golden-master:") {
			end := strings.Index(afterStart[4:], "-->")
			if end == -1 {
				break
			}
			content = content[:start] + content[start+4+end+3:]
			continue
		}
		end := strings.Index(afterStart, "-->")
		if end == -1 {
			break
		}
		content = content[:start] + content[start+end+3:]
	}

	// Step 3: Trim trailing whitespace from each line
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	return strings.Join(lines, "\n")
}
