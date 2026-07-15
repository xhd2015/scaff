## Expected

- Exit code is `0`.
- Body includes `name: scaff/install-via-curl` and/or rule marker `install/via-curl`.

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
	hasName := strings.Contains(out, "name: scaff/install-via-curl")
	hasRule := strings.Contains(out, "install/via-curl")
	if !hasName && !hasRule {
		t.Fatalf("expected name: scaff/install-via-curl or rule install/via-curl, got:\n%s", out)
	}
}
```
