package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintRulesIncludesLintAndFix(t *testing.T) {
	var buf bytes.Buffer
	PrintRules(&buf)
	out := buf.String()
	for _, id := range []string{"git/ignore", "github/testing-workflow", "script/generate", "git/hooks"} {
		if !strings.Contains(out, id) {
			t.Fatalf("expected rules output to contain %q:\n%s", id, out)
		}
	}
	if !strings.Contains(out, "lint:") || !strings.Contains(out, "fix:") {
		t.Fatalf("expected lint and fix sections:\n%s", out)
	}
}

func TestPrintRulesJSON(t *testing.T) {
	var buf bytes.Buffer
	if err := PrintRulesJSON(&buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, `"lint"`) || !strings.Contains(out, `"fix"`) {
		t.Fatalf("expected JSON lint/fix keys:\n%s", out)
	}
	if !strings.Contains(out, `"git/ignore"`) {
		t.Fatalf("expected git/ignore in JSON:\n%s", out)
	}
}