## Expected

- Exit code is `0`.
- `cmd/myapp/main.go` is not created on disk.
- Output previews the create action.

## Side Effects

- No `cmd/myapp/main.go` file.

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
	if fileExists(t, req, "cmd/myapp/main.go") {
		t.Fatal("cmd/myapp/main.go should not be created in dry-run")
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") && !strings.Contains(out, "cmd/myapp/main.go") {
		t.Fatalf("expected dry-run preview, got:\n%s", resp.Combined)
	}
}
```