## Expected

- Exit code is `0`.
- Created `.gitignore` does **not** contain `vendor/`.

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
	content := readProjectFile(t, req, ".gitignore")
	if strings.Contains(content, "vendor/") {
		t.Fatalf("expected .gitignore NOT to contain vendor/, got:\n%s", content)
	}
}
```