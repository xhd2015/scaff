## Expected

- Exit code is `0`.
- Mentions `--show`, `--install`, and `--list`.
- Includes an available-topics section (e.g. `Available topics:`) or concrete topic paths from TreeFS.

## Side Effects

- None.

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
	out := resp.Stdout
	if out == "" {
		out = resp.Combined
	}
	for _, flag := range []string{"--show", "--install", "--list"} {
		if !strings.Contains(out, flag) {
			t.Fatalf("skill help missing %q, got:\n%s", flag, out)
		}
	}
	// skillcmd FormatTopicIndex or equivalent topic listing.
	hasIndex := strings.Contains(out, "Available topics:")
	hasPaths := strings.Contains(out, "git/ignore") || strings.Contains(out, "overview") ||
		strings.Contains(out, "install-via-curl")
	if !hasIndex && !hasPaths {
		t.Fatalf("expected Available topics index or topic paths in skill help, got:\n%s", out)
	}
}
```
