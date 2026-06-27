## Expected

- Exit code is `1`.
- Output mentions default lint rules for the nested Go project (`git.ignore`, `github.testing.workflow`).

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
	for _, want := range []string{"git.ignore", "github.testing.workflow"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in output, got:\n%s", want, out)
		}
	}
}
```