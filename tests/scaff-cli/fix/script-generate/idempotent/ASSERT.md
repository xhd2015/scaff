## Expected

- Exit code is `0`.
- Existing `script/generate/main.go` is unchanged.
- Output indicates nothing to do.

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
	content := readProjectFile(t, req, "script/generate/main.go")
	if !strings.Contains(content, "add generators as subpackages") {
		t.Fatalf("expected original stub preserved, got:\n%s", content)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "nothing") && !strings.Contains(out, "exists") && !strings.Contains(out, "no-op") && !strings.Contains(out, "present") {
		t.Fatalf("expected no-op message, got:\n%s", resp.Combined)
	}
}
```