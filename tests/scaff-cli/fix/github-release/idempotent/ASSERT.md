## Expected

- Exit code is `0`.
- Existing release files are not overwritten.
- Output indicates nothing to do (or already exists).

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
	lib := readProjectFile(t, req, "script/github/lib/build_release.go")
	if !strings.Contains(main, "CUSTOM_RELEASE_MAIN") {
		t.Fatalf("expected custom main preserved, got:\n%s", main)
	}
	if !strings.Contains(lib, "CUSTOM_RELEASE_LIB") {
		t.Fatalf("expected custom lib preserved, got:\n%s", lib)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "nothing") && !strings.Contains(out, "exists") && !strings.Contains(out, "no-op") && !strings.Contains(out, "present") && !strings.Contains(out, "already") {
		t.Fatalf("expected no-op message, got:\n%s", resp.Combined)
	}
}
```