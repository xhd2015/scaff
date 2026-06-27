## Expected

- Exit code is `0`.
- Output indicates all rules pass (e.g. "all good").

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
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "all good") && !strings.Contains(out, "passing") {
		t.Fatalf("expected success message, got:\n%s", resp.Combined)
	}
}
```