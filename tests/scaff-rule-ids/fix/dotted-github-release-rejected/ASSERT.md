## Expected

- Exit code is non-zero (product convention: `2` for unknown rule).
- Output reports an unknown rule and mentions `github.release`.
- Dotted ID is not accepted as an alias for `github/release`.

## Exit Code

- non-zero (`2` preferred)

## Errors

- Unknown rule for `github.release`.

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
		t.Fatalf("exit code = 0, want non-zero for dotted id github.release\n%s", resp.Combined)
	}
	out := resp.Combined
	if !strings.Contains(out, "github.release") {
		t.Fatalf("expected dotted id github.release in output, got:\n%s", out)
	}
	low := strings.ToLower(out)
	if !strings.Contains(low, "unknown") {
		t.Fatalf("expected unknown-rule signal for github.release, got:\n%s", out)
	}
}
```
