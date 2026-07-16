## Expected

- Exit code is `1`.
- Output mentions `git/ignore`.
- Output mentions `github/testing-workflow`.
- Output mentions `project/readme`.
- Output mentions `project/license`.
- Output mentions `tests/doctest`.
- Output does **not** mention opt-in fix rules (`script/generate`, `git/hooks`, `github/release`, `install/via-curl`, `script/github/release-assets`, `project/agents`, `project/layout/cmd`, `script/dev`).

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
	for _, want := range []string{"git/ignore", "github/testing-workflow", "project/readme", "project/license", "tests/doctest"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected output to mention %q, got:\n%s", want, out)
		}
	}
	for _, omit := range []string{"script/generate", "git/hooks", "github/release", "install/via-curl", "script/github/release-assets", "project/agents", "project/layout/cmd", "script/dev"} {
		if strings.Contains(out, omit) {
			t.Fatalf("expected output NOT to mention %q, got:\n%s", omit, out)
		}
	}
}
```
