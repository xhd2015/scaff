## Expected

- Exit code is `0`.
- `LICENSE` is created at the project root.
- Content includes `MIT License`.
- Content includes `xhd2015` (owner from module).
- Content includes current calendar year and `-present` in copyright line.
- No `__OWNER__` or `__YEAR__` placeholders remain.

## Side Effects

- `LICENSE` created at repo root with mode `0644`.

## Exit Code

- `0`

```go
import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("exit code = %d, want 0\n%s", resp.ExitCode, resp.Combined)
	}
	rel := "LICENSE"
	if !fileExists(t, req, rel) {
		t.Fatal("LICENSE was not created")
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"MIT License",
		"xhd2015",
		fmt.Sprintf("%d-present", time.Now().Year()),
		"Permission is hereby granted",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected LICENSE to contain %q, got:\n%s", want, content)
		}
	}
	if strings.Contains(content, "__OWNER__") || strings.Contains(content, "__YEAR__") {
		t.Fatalf("expected placeholders substituted in LICENSE, got:\n%s", content)
	}
}
```