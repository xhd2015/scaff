## Expected

- Exit code is `0`.
- Body includes frontmatter `name: scaff/overview` (or equivalent identity markers for the overview product-model topic).
- Does not print the root-only index as a substitute (topic is path-resolved).

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
	if !strings.Contains(out, "name: scaff/overview") {
		t.Fatalf("expected frontmatter name: scaff/overview, got:\n%s", out)
	}
}
```
