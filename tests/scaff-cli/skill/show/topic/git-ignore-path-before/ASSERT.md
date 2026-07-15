## Expected

- Exit code is `0`.
- Same identity markers as flag-before order: `name: scaff/git/ignore` and/or `git.ignore`.

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
	hasName := strings.Contains(out, "name: scaff/git/ignore")
	hasRule := strings.Contains(out, "git.ignore")
	if !hasName && !hasRule {
		t.Fatalf("expected name: scaff/git/ignore or rule git.ignore, got:\n%s", out)
	}
}
```
