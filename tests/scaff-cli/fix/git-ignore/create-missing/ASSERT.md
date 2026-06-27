## Expected

- Exit code is `0`.
- `.gitignore` is created.
- File contains universal patterns: `.DS_Store`, `.vscode/`, `*.swp`, `*~`.
- File contains Go patterns: `bin/`, `*.test`, `coverage.out`.
- File does **not** contain `vendor/`.

## Side Effects

- `.gitignore` created at project root.

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
	if content == "" {
		t.Fatal(".gitignore was not created")
	}
	for _, want := range []string{".DS_Store", ".vscode/", "*.swp", "*~", "bin/", "*.test", "coverage.out"} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected .gitignore to contain %q, got:\n%s", want, content)
		}
	}
	if strings.Contains(content, "vendor/") {
		t.Fatalf("expected .gitignore NOT to contain vendor/, got:\n%s", content)
	}
}
```