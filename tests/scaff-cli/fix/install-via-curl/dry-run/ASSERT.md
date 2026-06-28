## Expected

- Exit code is `0`.
- `install-via-curl.sh` is **not** created.
- Output reports would-create for the installer path.

## Side Effects

- No installer script written.

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
	if fileExists(t, req, "install-via-curl.sh") {
		t.Fatal("install-via-curl.sh should not be created in dry-run")
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") && !strings.Contains(out, "create") {
		t.Fatalf("expected would-create preview, got:\n%s", resp.Combined)
	}
	if !strings.Contains(out, "install-via-curl") {
		t.Fatalf("expected installer path in preview, got:\n%s", resp.Combined)
	}
}
```