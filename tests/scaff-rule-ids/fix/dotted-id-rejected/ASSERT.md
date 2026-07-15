## Expected

- Exit code is non-zero (product convention: `2` for unknown rule).
- Output reports an unknown rule and mentions `git.ignore`.
- Dotted ID is not accepted as an alias for `git/ignore`.

## Exit Code

- non-zero (`2` preferred)

## Errors

- Unknown rule for `git.ignore`.

```go
import (
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
	if resp.ExitCode == 0 {
		t.Fatalf("exit code = 0, want non-zero for dotted id git.ignore\n%s", resp.Combined)
	}
	out := resp.Combined
	if !strings.Contains(out, "git.ignore") {
		t.Fatalf("expected dotted id git.ignore in output, got:\n%s", out)
	}
	low := strings.ToLower(out)
	if !strings.Contains(low, "unknown") {
		t.Fatalf("expected unknown-rule signal for git.ignore, got:\n%s", out)
	}
}
```
