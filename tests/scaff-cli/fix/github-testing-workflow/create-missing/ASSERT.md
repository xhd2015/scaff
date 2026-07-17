## Expected

- Exit code is `0`.
- `.github/workflows/test.yml` is created.
- Workflow runs `go test -v ./...` before `doctest test -v --label-all ./...`.
- Workflow uses a `golang:` container image derived from `go.mod`.

## Side Effects

- `test.yml` created at `.github/workflows/test.yml`.

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
	if !fileExists(t, req, ".github/workflows/test.yml") {
		t.Fatal("test.yml was not created")
	}
	content := readProjectFile(t, req, ".github/workflows/test.yml")
	for _, want := range []string{"go test -v ./...", "doctest test -v --label-all ./...", "golang:"} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected workflow to contain %q, got:\n%s", want, content)
		}
	}
	goTest := strings.Index(content, "go test -v ./...")
	doctest := strings.Index(content, "doctest test -v --label-all ./...")
	if goTest < 0 || doctest < 0 || goTest > doctest {
		t.Fatalf("expected go test before doctest, got:\n%s", content)
	}
}
```
