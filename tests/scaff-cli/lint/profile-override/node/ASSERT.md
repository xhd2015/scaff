## Expected

- Exit code is `1`.
- Output mentions `git/ignore` and `node_modules`.

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
	if !strings.Contains(out, "git/ignore") {
		t.Fatalf("expected git/ignore in output, got:\n%s", out)
	}
	if !strings.Contains(out, "node_modules") {
		t.Fatalf("expected node_modules in output, got:\n%s", out)
	}
}
```