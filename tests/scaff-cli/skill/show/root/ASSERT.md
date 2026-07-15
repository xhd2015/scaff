## Expected

- Exit code is `0`.
- Stdout includes YAML frontmatter `name: scaff` (root skill identity).
- Body describes the multi-topic index / skill surface.
- Body includes retrieve examples using `scaff skill --show`.
- Root body does **not** document install plumbing flags `--cursor` or `--global`.

## Side Effects

- None (show is read-only).

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
	if !strings.Contains(out, "name: scaff") {
		t.Fatalf("expected frontmatter name: scaff in stdout, got:\n%s", out)
	}
	if !strings.Contains(out, "scaff skill --show") {
		t.Fatalf("expected retrieve example 'scaff skill --show' in root body, got:\n%s", out)
	}
	// Multi-topic index markers (any of these signals index intent).
	lower := strings.ToLower(out)
	if !strings.Contains(lower, "topic") && !strings.Contains(lower, "multi-topic") && !strings.Contains(lower, "topics") {
		t.Fatalf("expected root body to describe multi-topic/index surface, got:\n%s", out)
	}
	// Install target flags belong in skill --install --help, not root SKILL.md body.
	if strings.Contains(out, "--cursor") {
		t.Fatalf("root skill body must not document --cursor install plumbing, got:\n%s", out)
	}
	if strings.Contains(out, "--global") {
		t.Fatalf("root skill body must not document --global install plumbing, got:\n%s", out)
	}
}
```
