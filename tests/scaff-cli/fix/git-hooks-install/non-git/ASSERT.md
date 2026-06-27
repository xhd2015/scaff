## Expected

- Exit code is `1`.
- Output indicates the directory is not a git repository.

## Errors

- Fix fails because `.git` is missing.

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
	if !strings.Contains(out, "git") {
		t.Fatalf("expected git-related error, got:\n%s", resp.Combined)
	}
}
```