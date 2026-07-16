## Expected

- Exit code is `2`.
- Output mentions unknown rule `unknown.rule`.
- Output lists available rules including the full fix catalog (18 slash IDs), among them the new project/tests/script rules.

## Exit Code

- `2`

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 2 {
		t.Fatalf("exit code = %d, want 2\n%s", resp.ExitCode, resp.Combined)
	}
	out := resp.Combined
	if !strings.Contains(out, "unknown.rule") {
		t.Fatalf("expected unknown.rule in output, got:\n%s", out)
	}
	for _, rule := range []string{
		"git/ignore",
		"github/testing-workflow",
		"project/readme",
		"project/license",
		"tests/doctest",
		"project/agents",
		"project/layout/cmd",
		"script/generate",
		"script/install",
		"script/build",
		"script/dev",
		"script/bundle/for-linux",
		"git/hooks",
		"git/hooks/install",
		"git/pre-commit",
		"github/release",
		"install/via-curl",
		"script/github/release-assets",
	} {
		if !strings.Contains(out, rule) {
			t.Fatalf("expected available rule %q in output, got:\n%s", rule, out)
		}
	}
}
```