## Expected

- Exit code is `0`.
- Stdout includes `[dry-run]` plan lines.
- Plan mentions `SKILL.md` under the target dir.
- Plan mentions nested topic files such as `overview/TOPIC.md` and `git/ignore/TOPIC.md`.
- No skill files are written on disk under the target dir.

## Side Effects

- Target directory remains without installed skill files (dry-run only).

## Exit Code

- `0`

```go
import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("exit code = %d, want 0\n%s", resp.ExitCode, resp.Combined)
	}
	out := resp.Stdout
	if !strings.Contains(out, "[dry-run]") {
		t.Fatalf("expected [dry-run] in install plan, got:\n%s", out)
	}
	if !strings.Contains(out, "SKILL.md") {
		t.Fatalf("expected SKILL.md in dry-run plan, got:\n%s", out)
	}
	// Nested TOPIC.md paths from inventory (skillcmd collectTreeSkillFiles).
	hasOverview := strings.Contains(out, "overview"+string(filepath.Separator)+"TOPIC.md") ||
		strings.Contains(out, "overview/TOPIC.md")
	hasGitIgnore := strings.Contains(out, "git"+string(filepath.Separator)+"ignore"+string(filepath.Separator)+"TOPIC.md") ||
		strings.Contains(out, "git/ignore/TOPIC.md")
	if !hasOverview || !hasGitIgnore {
		t.Fatalf("expected nested TOPIC.md paths (overview, git/ignore) in dry-run plan, got:\n%s", out)
	}
	// Dry-run must not materialize skill files.
	skillMD := filepath.Join(req.ProjectDir, "SKILL.md")
	if _, statErr := os.Stat(skillMD); !os.IsNotExist(statErr) {
		t.Fatalf("dry-run must not create %s", skillMD)
	}
	overview := filepath.Join(req.ProjectDir, "overview", "TOPIC.md")
	if _, statErr := os.Stat(overview); !os.IsNotExist(statErr) {
		t.Fatalf("dry-run must not create %s", overview)
	}
}
```
