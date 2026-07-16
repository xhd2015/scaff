## Expected

- Exit code is `0`.
- `LICENSE` is **not** created.
- Output reports would-create for `LICENSE`.

## Side Effects

- No `LICENSE` written.

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
	if fileExists(t, req, "LICENSE") {
		t.Fatal("LICENSE should not be created in dry-run")
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") && !strings.Contains(out, "create") {
		t.Fatalf("expected would-create preview, got:\n%s", resp.Combined)
	}
	if !strings.Contains(out, "license") {
		t.Fatalf("expected LICENSE path hint in preview, got:\n%s", resp.Combined)
	}
}
```