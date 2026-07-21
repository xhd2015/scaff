# Scenario

**Feature**: fix script/bundle/for-linux is idempotent

```
# existing bundle stub -> no-op
script/bundle/for-linux fix -> nothing to do
```

## Preconditions

- `script/bundle/for-linux/main.go` already exists.

## Steps

1. Write existing stub.
2. Run `scaff fix script/bundle/for-linux`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeScriptBundleForLinux(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/bundle/for-linux"}
	return nil
}
```