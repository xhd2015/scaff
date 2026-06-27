## Expected

- Exit code is `0`.
- `.github/workflows/test.yml` is created.
- `.github/workflows/ci.yml` remains unchanged.

## Side Effects

- `test.yml` created; `ci.yml` preserved.

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
	if !fileExists(t, req, ".github/workflows/test.yml") {
		t.Fatal("test.yml was not created")
	}
	ci := readProjectFile(t, req, ".github/workflows/ci.yml")
	if !strings.Contains(ci, "echo ci") {
		t.Fatalf("expected ci.yml preserved, got:\n%s", ci)
	}
}
```