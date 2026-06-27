## Expected

- Exit code is `0`.
- `script/generate/main.go` is created with the no-op stub content.

## Side Effects

- `script/generate/main.go` created.

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
	if !fileExists(t, req, "script/generate/main.go") {
		t.Fatal("script/generate/main.go was not created")
	}
	content := readProjectFile(t, req, "script/generate/main.go")
	for _, want := range []string{"package main", "go run ./script/generate", "add generators as subpackages"} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected stub to contain %q, got:\n%s", want, content)
		}
	}
}
```