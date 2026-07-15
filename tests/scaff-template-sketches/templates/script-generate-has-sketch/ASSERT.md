## Expected

- Exit code is `0`.
- `script/generate/main.go` is created.
- File content includes the substring `Proposed behavior` (sketch header).
- File remains a generator stub (`package main`, path/usage for `script/generate`).

## Side Effects

- `script/generate/main.go` written from the generate template.

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
	rel := "script/generate/main.go"
	if !fileExists(t, req, rel) {
		t.Fatalf("%s was not created", rel)
	}
	content := readProjectFile(t, req, rel)
	if !strings.Contains(content, "package main") {
		t.Fatalf("expected package main in %s, got:\n%s", rel, content)
	}
	if !strings.Contains(content, "script/generate") {
		t.Fatalf("expected usage/path hint for script/generate in %s, got:\n%s", rel, content)
	}
	assertProposedBehaviorSketch(t, content, rel)
}
```
