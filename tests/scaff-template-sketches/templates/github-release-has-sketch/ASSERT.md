## Expected

- Exit code is `0`.
- `script/github/release/main.go` is created.
- File content includes the substring `Proposed behavior` (sketch header).
- Release entrypoint still identifies the release tool path
  (`script/github/release` or equivalent usage).

## Side Effects

- Release scaffold under `script/github/` (main required; lib may also be created).

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
	if resp == nil {
		t.Fatal("nil response")
	}
	if resp.ExitCode != 0 {
		t.Fatalf("exit code = %d, want 0\n%s", resp.ExitCode, resp.Combined)
	}
	rel := "script/github/release/main.go"
	if !fileExists(t, req, rel) {
		t.Fatalf("%s was not created", rel)
	}
	content := readProjectFile(t, req, rel)
	if !strings.Contains(content, "package main") {
		t.Fatalf("expected package main in %s, got:\n%s", rel, content)
	}
	if !strings.Contains(content, "script/github/release") {
		t.Fatalf("expected release path usage in %s, got:\n%s", rel, content)
	}
	assertProposedBehaviorSketch(t, content, rel)
}
```
