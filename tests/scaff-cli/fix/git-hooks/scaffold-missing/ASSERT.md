## Expected

- Exit code is `0`.
- `script/git-hooks/main.go` is created.
- File supports `install`, `pre-commit`, and `pre-push` commands.
- No `pre-commit/go-test/` or similar sub-check directories are created.

## Side Effects

- `script/git-hooks/main.go` created.

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
	if !fileExists(t, req, "script/git-hooks/main.go") {
		t.Fatal("script/git-hooks/main.go was not created")
	}
	content := readProjectFile(t, req, "script/git-hooks/main.go")
	for _, want := range []string{"install", "pre-commit", "pre-push"} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected hook runner to reference %q, got:\n%s", want, content)
		}
	}
	badDirs := []string{
		"pre-commit/go-test",
		"script/git-hooks/pre-commit",
	}
	for _, rel := range badDirs {
		if _, err := os.Stat(filepath.Join(req.ProjectDir, rel)); err == nil {
			t.Fatalf("unexpected sub-check directory %q created", rel)
		}
	}
}
```