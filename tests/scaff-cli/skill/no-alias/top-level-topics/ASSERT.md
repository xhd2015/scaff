## Expected

- Exit code is non-zero (unknown command; typically `2`).
- Output indicates unknown command `topics`.
- Does not print the skill topic inventory as a successful topics command.

## Errors

- Unknown top-level command.

## Exit Code

- non-zero

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode == 0 {
		t.Fatalf("exit code = 0, want non-zero for top-level topics alias\n%s", resp.Combined)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "unknown") || !strings.Contains(out, "topics") {
		t.Fatalf("expected unknown command topics, got:\n%s", resp.Combined)
	}
}
```
