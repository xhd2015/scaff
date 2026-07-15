## Expected

- Exit code is `0`.
- `script/github/release-assets/main.go` is not created on disk.
- Output previews the create action.

## Side Effects

- No `script/github/release-assets/main.go` file.

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
	if fileExists(t, req, "script/github/release-assets/main.go") {
		t.Fatal("script/github/release-assets/main.go should not be created in dry-run")
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") && !strings.Contains(out, "script/github/release-assets") {
		t.Fatalf("expected dry-run preview, got:\n%s", resp.Combined)
	}
}
```
