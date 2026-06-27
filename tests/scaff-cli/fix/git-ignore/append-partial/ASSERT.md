## Expected

- Exit code is `0`.
- `.gitignore` retains existing lines.
- Missing patterns (`.vscode/`, `*.swp`, `*~`) are appended.
- Output indicates lines were appended.

## Side Effects

- `.gitignore` updated in place.

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
	content := readProjectFile(t, req, ".gitignore")
	for _, want := range []string{".DS_Store", "bin/", ".vscode/", "*.swp", "*~"} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected .gitignore to contain %q, got:\n%s", want, content)
		}
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "append") && !strings.Contains(out, "added") {
		t.Fatalf("expected append message, got:\n%s", resp.Combined)
	}
}
```