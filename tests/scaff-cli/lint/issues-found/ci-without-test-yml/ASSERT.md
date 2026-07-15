## Expected

- Exit code is `1`.
- Output mentions `github/testing-workflow`.
- Output references missing `test.yml` (or `.github/workflows/test.yml`).

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
	if !strings.Contains(out, "github/testing-workflow") {
		t.Fatalf("expected github/testing-workflow in output, got:\n%s", out)
	}
	if !strings.Contains(out, "test.yml") {
		t.Fatalf("expected test.yml reference in output, got:\n%s", out)
	}
}
```