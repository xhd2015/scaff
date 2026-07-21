# Scenario

**Feature**: fix script/generate is idempotent

```
# existing stub -> no-op
script/generate fix -> nothing to do
```

## Preconditions

- `script/generate/main.go` already exists.

## Steps

1. Write existing stub.
2. Run `scaff fix script/generate`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeScriptGenerate(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/generate"}
	return nil
}
```