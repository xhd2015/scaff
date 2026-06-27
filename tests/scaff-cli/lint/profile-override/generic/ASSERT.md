## Expected

- Exit code is `1`.
- Output mentions `git.ignore`.
- Output does **not** mention Go-only patterns (`bin/`) or Node-only patterns (`node_modules`).

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
	out := resp.Combined
	if !strings.Contains(out, "git.ignore") {
		t.Fatalf("expected git.ignore in output, got:\n%s", out)
	}
	for _, omit := range []string{"node_modules", "bin/"} {
		if strings.Contains(out, omit) {
			t.Fatalf("expected generic profile NOT to mention %q, got:\n%s", omit, out)
		}
	}
}
```