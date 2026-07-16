## Expected

- Exit code is `0`.
- `AGENTS.md` is created at the project root.
- Content includes `# AGENTS.md`, `myapp`, `go run ./script/build`, and `doctest test`.
- No `__NAME__` placeholders remain.

## Side Effects

- `AGENTS.md` created at repo root with mode `0644`.

## Exit Code

- `0`

```go
import (
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
	rel := "AGENTS.md"
	if !fileExists(t, req, rel) {
		t.Fatal("AGENTS.md was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"# AGENTS.md",
		"myapp",
		"go run ./script/build",
		"doctest test",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected AGENTS.md to contain %q, got:\n%s", want, content)
		}
	}
	if strings.Contains(content, "__NAME__") {
		t.Fatalf("expected placeholders substituted in AGENTS.md, got:\n%s", content)
	}
}
```