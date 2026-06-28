## Expected

- Exit code is `0`.
- Release scaffold files are **not** created.
- Output reports would-create for release paths (main and/or lib).

## Side Effects

- No files written under `script/github/`.

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
	if fileExists(t, req, "script/github/release/main.go") {
		t.Fatal("release main.go should not be created in dry-run")
	}
	if fileExists(t, req, "script/github/lib/build_release.go") {
		t.Fatal("build_release.go should not be created in dry-run")
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") && !strings.Contains(out, "create") {
		t.Fatalf("expected would-create preview, got:\n%s", resp.Combined)
	}
	if !strings.Contains(out, "script/github") && !strings.Contains(out, "release") {
		t.Fatalf("expected release path hint in preview, got:\n%s", resp.Combined)
	}
}
```