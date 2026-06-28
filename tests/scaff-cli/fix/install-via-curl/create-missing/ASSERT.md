## Expected

- Exit code is `0`.
- `install-via-curl.sh` is created at the project root.
- Script is executable (`0755` or user-executable bit set).
- Content includes GitHub release URL for `xhd2015/myapp`, asset naming `myapp-v`, platform targets (`darwin-amd64`, `linux-amd64`), and Windows/WSL rejection.

## Side Effects

- `install-via-curl.sh` created at repo root.

## Exit Code

- `0`

```go
import (
	"os"
	"path/filepath"
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
	rel := "install-via-curl.sh"
	if !fileExists(t, req, rel) {
		t.Fatal("install-via-curl.sh was not created")
	}
	info, err := os.Stat(filepath.Join(req.ProjectDir, rel))
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode()&0o111 == 0 {
		t.Fatalf("expected install-via-curl.sh executable, mode=%v", info.Mode())
	}
	content := readProjectFile(t, req, rel)
	for _, want := range []string{
		"set -eo pipefail",
		"https://github.com/xhd2015/myapp/releases",
		"myapp-v",
		"darwin-amd64",
		"linux-amd64",
		"Windows_NT",
		"Windows Subsystem for Linux",
		"INSTALL_TAG",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected installer to contain %q, got:\n%s", want, content)
		}
	}
	if strings.Contains(content, "__OWNER__") || strings.Contains(content, "__REPO__") || strings.Contains(content, "__NAME__") {
		t.Fatalf("expected placeholders substituted, got:\n%s", content)
	}
}
```