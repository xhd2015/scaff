## Expected

- Exit code is `0`.
- Body includes `name: scaff/github/release` and/or rule marker `github.release`.

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
	hasName := strings.Contains(out, "name: scaff/github/release")
	hasRule := strings.Contains(out, "github.release")
	if !hasName && !hasRule {
		t.Fatalf("expected name: scaff/github/release or rule github.release, got:\n%s", out)
	}
}
```
