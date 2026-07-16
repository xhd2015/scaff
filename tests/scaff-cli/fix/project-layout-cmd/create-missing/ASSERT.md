## Expected

- Exit code is `0`.
- `cmd/myapp/main.go` is created with the cmd entry template.
- Content includes `package main`, `func run`, and stderr format uses `myapp` (not `__NAME__`).

## Side Effects

- `cmd/myapp/main.go` created.

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
	rel := "cmd/myapp/main.go"
	if !fileExists(t, req, rel) {
		t.Fatal("cmd/myapp/main.go was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"package main",
		"func run",
		`fmt.Fprintf(os.Stderr, "myapp: %v\n", err)`,
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected cmd main to contain %q, got:\n%s", want, content)
		}
	}
	if strings.Contains(content, "__NAME__") {
		t.Fatalf("expected __NAME__ substituted in cmd main, got:\n%s", content)
	}
}
```