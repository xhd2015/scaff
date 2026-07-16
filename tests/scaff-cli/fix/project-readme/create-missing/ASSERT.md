## Expected

- Exit code is `0`.
- `README.md` is created at the project root.
- Content includes `# myapp` heading.
- Content includes `go install github.com/xhd2015/myapp@latest`.
- No `__NAME__` or `__MODULE__` placeholders remain.

## Side Effects

- `README.md` created at repo root with mode `0644`.

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
	rel := "README.md"
	if !fileExists(t, req, rel) {
		t.Fatal("README.md was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"# myapp",
		"go install github.com/xhd2015/myapp@latest",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected README.md to contain %q, got:\n%s", want, content)
		}
	}
	if strings.Contains(content, "__NAME__") || strings.Contains(content, "__MODULE__") {
		t.Fatalf("expected placeholders substituted in README.md, got:\n%s", content)
	}
}
```