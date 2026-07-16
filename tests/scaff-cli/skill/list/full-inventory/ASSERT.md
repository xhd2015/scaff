## Expected Output

```
---
version: 2
---
scaff
fix
git/hooks
git/hooks/install
git/ignore
git/pre-commit
github/release
github/testing-workflow
github/upload
install-via-curl
lint
overview
project/agents
project/layout/cmd
project/license
project/readme
script/build
script/bundle-for-linux
script/dev
script/generate
script/github/release-assets
script/install
tests/doctest
```

## Expected

- Exit code is `0`.
- First line is exactly `scaff` (skill name).
- Stdout lists every nested topic path from the inventory, sorted, one per line
  (includes project/*, tests/doctest, script/dev, git/pre-commit).
- Trailing newline after the last topic line (CLI convention).

## Side Effects

- None (list is read-only).

## Exit Code

- `0`

```go
import (
	"testing"

	"github.com/xhd2015/doctest/assert"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("exit code = %d, want 0\n%s", resp.ExitCode, resp.Combined)
	}
	// Full inventory, skillcmd ListTreeTopics order (sort.Strings).
	assert.Output(t, resp.Stdout, `---
version: 2
---
scaff
fix
git/hooks
git/hooks/install
git/ignore
git/pre-commit
github/release
github/testing-workflow
github/upload
install-via-curl
lint
overview
project/agents
project/layout/cmd
project/license
project/readme
script/build
script/bundle-for-linux
script/dev
script/generate
script/github/release-assets
script/install
tests/doctest
`)
}
```

