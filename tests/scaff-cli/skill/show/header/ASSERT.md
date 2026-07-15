## Expected

- Exit code is `0`.
- Stdout is the YAML frontmatter block with `---` delimiters.
- Contains `name: scaff`.
- Does not include root body content (e.g. retrieve examples `scaff skill --show`).

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
	if !strings.HasPrefix(out, "---") {
		t.Fatalf("header output must start with ---, got:\n%s", out)
	}
	if !strings.Contains(out, "name: scaff") {
		t.Fatalf("expected name: scaff in header, got:\n%s", out)
	}
	// skillcmd FormatHeaderWithDelimiters: ---\n<header>\n---\n
	if strings.Count(out, "---") < 2 {
		t.Fatalf("expected opening and closing --- delimiters, got:\n%s", out)
	}
	// Body retrieve examples must not appear in header-only mode.
	if strings.Contains(out, "scaff skill --show") {
		t.Fatalf("header-only must omit root body (found scaff skill --show):\n%s", out)
	}
}
```
