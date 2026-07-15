## Expected

- Exit code is `0`.
- Command is accepted as a known rule (not "unknown rule").
- Dry-run must not create release files under `script/github/`.

## Exit Code

- `0`

## Errors

- Must not report unknown rule for `github/release`.

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
	if resp == nil {
		t.Fatal("nil response")
	}
	if resp.ExitCode != 0 {
		t.Fatalf("exit code = %d, want 0 for slash id github/release\n%s", resp.ExitCode, resp.Combined)
	}
	out := strings.ToLower(resp.Combined)
	if strings.Contains(out, "unknown rule") {
		t.Fatalf("slash id github/release must not be unknown:\n%s", resp.Combined)
	}
	// Dry-run must not write scaffold files.
	for _, rel := range []string{
		"script/github/release/main.go",
		"script/github/lib/build_release.go",
	} {
		if _, err := os.Stat(filepath.Join(req.ProjectDir, rel)); err == nil {
			t.Fatalf("dry-run must not create %s", rel)
		}
	}
}
```
