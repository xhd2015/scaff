## Expected

- Exit code is `0`.
- Body includes frontmatter `name: scaff/github/upload` (docs-only upload topic).

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
	if !strings.Contains(out, "name: scaff/github/upload") {
		t.Fatalf("expected name: scaff/github/upload, got:\n%s", out)
	}
}
```
