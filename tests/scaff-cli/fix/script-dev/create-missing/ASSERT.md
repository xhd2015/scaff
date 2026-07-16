## Expected

- Exit code is `0`.
- `script/dev/main.go` is created with the dev wrapper stub content.

## Side Effects

- `script/dev/main.go` created.

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
	rel := "script/dev/main.go"
	if !fileExists(t, req, rel) {
		t.Fatal("script/dev/main.go was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"go run ./script/dev",
		"func Handle",
		`"run", ".", "--dev"`,
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected dev stub to contain %q, got:\n%s", want, content)
		}
	}
}
```