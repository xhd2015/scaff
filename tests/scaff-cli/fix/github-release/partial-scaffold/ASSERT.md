## Expected

- Exit code is `0`.
- Existing `main.go` is preserved (still contains `CUSTOM_RELEASE_MAIN`).
- Missing `build_release.go` is created with `myapp` substitution.

## Side Effects

- Only the lib file is added; release main is not overwritten.

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
	main := readProjectFile(t, req, "script/github/release/main.go")
	if !strings.Contains(main, "CUSTOM_RELEASE_MAIN") {
		t.Fatalf("expected custom main preserved, got:\n%s", main)
	}
	lib := readProjectFile(t, req, "script/github/lib/build_release.go")
	if lib == "" {
		t.Fatal("build_release.go was not created")
	}
	if !strings.Contains(lib, "myapp") {
		t.Fatalf("expected lib to substitute myapp, got:\n%s", lib)
	}
}
```