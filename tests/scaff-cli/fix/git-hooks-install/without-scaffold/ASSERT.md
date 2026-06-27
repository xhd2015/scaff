## Expected

- Exit code is `1`.
- Output includes hint to run `scaff fix git.hooks` (or `fix git.hooks`).

## Errors

- Fix fails because hook scaffold is missing.

## Exit Code

- `1`

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 1 {
		t.Fatalf("exit code = %d, want 1\n%s", resp.ExitCode, resp.Combined)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "git.hooks") {
		t.Fatalf("expected hint mentioning git.hooks, got:\n%s", resp.Combined)
	}
}
```