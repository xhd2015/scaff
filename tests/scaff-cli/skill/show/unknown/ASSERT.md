## Expected

- Exit code is non-zero.
- Stderr (or combined output) indicates the topic is unknown/missing.
- Output mentions the requested path `not-a-real-topic`.

## Errors

- Unknown / missing topic path.

## Exit Code

- non-zero

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode == 0 {
		t.Fatalf("exit code = 0, want non-zero for unknown topic\n%s", resp.Combined)
	}
	out := resp.Combined
	if !strings.Contains(out, "not-a-real-topic") {
		t.Fatalf("expected topic path not-a-real-topic in error output, got:\n%s", out)
	}
	lower := strings.ToLower(out)
	// skillcmd: "unknown topic path" or "read skill <path>: ..."
	if !strings.Contains(lower, "unknown") && !strings.Contains(lower, "not found") &&
		!strings.Contains(lower, "no such") && !strings.Contains(lower, "read skill") &&
		!strings.Contains(lower, "missing") {
		t.Fatalf("expected unknown/missing topic error signal, got:\n%s", out)
	}
}
```
