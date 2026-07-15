## Expected

- Exit code is `0`.
- `script/github/release-assets/main.go` is created.
- Content includes `Proposed behavior` sketch.
- Content documents packable help: `--upload` opt-in and a generic directory flag (`--dir`).
- Not browser-agent-specific (must not hard-require `browser-agent` as the only tool).

## Side Effects

- `script/github/release-assets/main.go` created.

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
	rel := "script/github/release-assets/main.go"
	if !fileExists(t, req, rel) {
		t.Fatal("script/github/release-assets/main.go was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"package main",
		"Proposed behavior",
		"--upload",
		"--dir",
		"script/github/release-assets",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected stub to contain %q, got:\n%s", want, content)
		}
	}
	// Keep the scaffold generic: do not require a browser-agent-only workflow.
	if strings.Contains(strings.ToLower(content), "browser-agent required") {
		t.Fatalf("stub must not hard-require browser-agent, got:\n%s", content)
	}
}
```
