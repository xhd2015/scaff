## Expected

- Exit code is `0`.
- `.git/hooks/pre-commit` contains `# scaff hooks` marker.
- `.git/hooks/pre-push` contains `# scaff hooks` marker.
- Hook scripts invoke `go run ./script/git-hooks`.

## Side Effects

- Git hook files patched under `.git/hooks/`.

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
	for _, hook := range []string{".git/hooks/pre-commit", ".git/hooks/pre-push"} {
		content := readProjectFile(t, req, hook)
		if content == "" {
			t.Fatalf("%s was not created or patched", hook)
		}
		if !strings.Contains(content, "# scaff hooks") {
			t.Fatalf("expected # scaff hooks marker in %s, got:\n%s", hook, content)
		}
		if !strings.Contains(content, "script/git-hooks") {
			t.Fatalf("expected script/git-hooks invocation in %s, got:\n%s", hook, content)
		}
	}
}
```