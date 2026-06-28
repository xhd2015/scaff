## Expected

- Exit code is `0`.
- `script/bundle/for-linux/main.go` is created with the bundle helper stub content.

## Side Effects

- `script/bundle/for-linux/main.go` created.

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
	rel := "script/bundle/for-linux/main.go"
	if !fileExists(t, req, rel) {
		t.Fatal("script/bundle/for-linux/main.go was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"go run ./script/bundle/for-linux",
		"GOOS=linux",
		"GOARCH=amd64",
		"Bundle ready",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected stub to contain %q, got:\n%s", want, content)
		}
	}
}
```