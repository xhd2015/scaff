## Expected

- Exit code is non-zero (unknown command; typically `2`).
- Output indicates unknown command `install`.
- Does not succeed as a skill install (no successful install plan for skill `scaff`).

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
		t.Fatalf("exit code = 0, want non-zero for top-level install alias\n%s", resp.Combined)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "unknown") || !strings.Contains(out, "install") {
		t.Fatalf("expected unknown command install, got:\n%s", resp.Combined)
	}
	// Must not look like a successful skillcmd dry-run install.
	if strings.Contains(resp.Combined, "[dry-run]") && strings.Contains(resp.Combined, "SKILL.md") {
		t.Fatalf("top-level install must not run skill install, got:\n%s", resp.Combined)
	}
}
```
