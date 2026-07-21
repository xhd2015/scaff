# Scenario

**Feature**: fix script/build is idempotent

```
# existing build stub -> no-op
script/build fix -> nothing to do
```

## Preconditions

- `script/build/build.go` already exists.

## Steps

1. Write existing stub.
2. Run `scaff fix script/build`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeScriptBuild(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/build"}
	return nil
}
```