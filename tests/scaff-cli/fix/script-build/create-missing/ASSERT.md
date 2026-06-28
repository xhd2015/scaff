## Expected

- Exit code is `0`.
- `script/build/build.go` is created with the build helper stub content.

## Side Effects

- `script/build/build.go` created.

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
	rel := "script/build/build.go"
	if !fileExists(t, req, rel) {
		t.Fatal("script/build/build.go was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"go run ./script/build",
		"func Handle",
		"go build",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected stub to contain %q, got:\n%s", want, content)
		}
	}
}
```