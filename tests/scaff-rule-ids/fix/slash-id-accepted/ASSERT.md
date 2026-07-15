## Expected

- Exit code is `0`.
- Command is accepted as a known rule (not "unknown rule").
- Dry-run may print a would-change preview; files need not change.

## Exit Code

- `0`

## Errors

- Must not report unknown rule for `git/ignore`.

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
	if resp.ExitCode != 0 {
		t.Fatalf("exit code = %d, want 0 for slash id git/ignore\n%s", resp.ExitCode, resp.Combined)
	}
	out := strings.ToLower(resp.Combined)
	if strings.Contains(out, "unknown rule") {
		t.Fatalf("slash id git/ignore must not be unknown:\n%s", resp.Combined)
	}
}
```
