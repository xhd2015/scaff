## Expected

- Exit code is `0`.
- `.gitignore` remains the original partial content (no `.vscode/` added).
- Output previews lines that would be appended.

## Side Effects

- `.gitignore` file unchanged on disk.

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
	content := readProjectFile(t, req, ".gitignore")
	if strings.Contains(content, ".vscode/") {
		t.Fatalf("expected .gitignore unchanged (no .vscode/), got:\n%s", content)
	}
	out := strings.ToLower(resp.Combined)
	if !strings.Contains(out, "dry") && !strings.Contains(out, "would") && !strings.Contains(out, ".vscode/") {
		t.Fatalf("expected dry-run preview, got:\n%s", resp.Combined)
	}
}
```