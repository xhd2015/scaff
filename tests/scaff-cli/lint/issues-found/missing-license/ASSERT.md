## Expected

- Exit code is `1`.
- Output mentions `project/license`.

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
	if !strings.Contains(resp.Combined, "project/license") {
		t.Fatalf("expected output to mention project/license, got:\n%s", resp.Combined)
	}
}
```