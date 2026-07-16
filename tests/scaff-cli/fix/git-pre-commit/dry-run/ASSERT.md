## Expected

- Exit code is `0`.
- `script/git/pre-commit/main.go` is not created.
- Output previews that the file would be created.

## Side Effects

- No pre-commit helper written on disk.

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
	if fileExists(t, req, "script/git/pre-commit/main.go") {
		t.Fatal("expected script/git/pre-commit/main.go not created on dry-run")
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") {
		t.Fatalf("expected dry-run preview, got:\n%s", resp.Combined)
	}
	if !strings.Contains(resp.Combined, "script/git/pre-commit/main.go") {
		t.Fatalf("expected path in dry-run preview, got:\n%s", resp.Combined)
	}
}
```
