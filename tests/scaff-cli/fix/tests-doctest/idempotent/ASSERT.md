## Expected

- Exit code is `0`.
- Existing `tests/myapp-cli/DOCTEST.md` is not overwritten.
- Output indicates nothing to do (or already exists).

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
	content := readProjectFile(t, req, "tests/myapp-cli/DOCTEST.md")
	if !strings.Contains(content, "CUSTOM_DOCTEST") {
		t.Fatalf("expected custom DOCTEST preserved, got:\n%s", content)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "nothing") && !strings.Contains(out, "exists") && !strings.Contains(out, "no-op") && !strings.Contains(out, "present") && !strings.Contains(out, "already") {
		t.Fatalf("expected no-op message, got:\n%s", resp.Combined)
	}
}
```