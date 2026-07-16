## Expected

- Exit code is `0`.
- `tests/myapp-cli/DOCTEST.md` is created.
- Content includes `## Version`, `0.0.2`, `# DSN`, `doctest test`, `type Request`, and `func Run`.
- `tests/myapp-cli/SETUP.md` is created with `# Scenario`.
- No `__NAME__` placeholders remain.

## Side Effects

- `tests/myapp-cli/DOCTEST.md` and `tests/myapp-cli/SETUP.md` created at repo root.

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
	doctestRel := "tests/myapp-cli/DOCTEST.md"
	setupRel := "tests/myapp-cli/SETUP.md"
	if !fileExists(t, req, doctestRel) {
		t.Fatal("tests/myapp-cli/DOCTEST.md was not created")
	}
	if !fileExists(t, req, setupRel) {
		t.Fatal("tests/myapp-cli/SETUP.md was not created")
	}
	doctest := readProjectFile(t, req, doctestRel)
	for _, want := range []string{
		"## Version",
		"0.0.2",
		"# DSN",
		"doctest test",
		"type Request",
		"func Run",
	} {
		if !strings.Contains(doctest, want) {
			t.Fatalf("expected DOCTEST.md to contain %q, got:\n%s", want, doctest)
		}
	}
	setup := readProjectFile(t, req, setupRel)
	if !strings.HasPrefix(strings.TrimSpace(setup), "# Scenario") {
		t.Fatalf("expected SETUP.md to start with # Scenario, got:\n%s", setup)
	}
	for _, path := range []string{doctest, setup} {
		if strings.Contains(path, "__NAME__") {
			t.Fatalf("expected placeholders substituted, got:\n%s", path)
		}
	}
}
```