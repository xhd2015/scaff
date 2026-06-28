## Expected

- Exit code is `0`.
- `script/github/release/main.go` and `script/github/lib/build_release.go` are created.
- `main.go` references `go run ./script/github/release`, kool `release` packages, and `--dry-run`.
- Substituted binary/owner metadata uses `myapp` and `xhd2015` (no `__NAME__` / `__OWNER__` placeholders).

## Side Effects

- Release entrypoint and lib helper created under `script/github/`.

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
	mainRel := "script/github/release/main.go"
	libRel := "script/github/lib/build_release.go"
	if !fileExists(t, req, mainRel) {
		t.Fatal("release main.go was not created")
	}
	if !fileExists(t, req, libRel) {
		t.Fatal("build_release.go was not created")
	}
	main := readProjectFile(t, req, mainRel)
	lib := readProjectFile(t, req, libRel)
	for _, want := range []string{
		"go run ./script/github/release",
		"github.com/xhd2015/kool/pkgs/release",
		"--dry-run",
		"myapp",
	} {
		if !strings.Contains(main, want) {
			t.Fatalf("expected main.go to contain %q, got:\n%s", want, main)
		}
	}
	if strings.Contains(main, "__NAME__") || strings.Contains(main, "__OWNER__") {
		t.Fatalf("expected placeholders substituted in main.go, got:\n%s", main)
	}
	if !strings.Contains(lib, "release.BuildRelease(\"myapp\"") {
		t.Fatalf("expected lib to call BuildRelease with myapp, got:\n%s", lib)
	}
}
```