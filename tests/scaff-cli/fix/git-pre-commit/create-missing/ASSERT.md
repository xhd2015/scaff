## Expected

- Exit code is `0`.
- `script/git/pre-commit/main.go` is created.
- File includes install comment for git-hooks and ensure/git add behavior.

## Side Effects

- `script/git/pre-commit/main.go` created.

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
	if !fileExists(t, req, "script/git/pre-commit/main.go") {
		t.Fatal("script/git/pre-commit/main.go was not created")
	}
	content := readProjectFile(t, req, "script/git/pre-commit/main.go")
	for _, want := range []string{
		"package main",
		"go run ./script/git/pre-commit",
		"git-hooks pre-commit add 'script.git.pre-commit' go run ./script/git/pre-commit",
		"var ensure",
		"git add",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected pre-commit stub to contain %q, got:\n%s", want, content)
		}
	}
}
```
