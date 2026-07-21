# Scenario

**Feature**: fix script/dev is idempotent

```
# existing dev stub -> no-op
script/dev fix -> nothing to do
```

## Preconditions

- `script/dev/main.go` already exists.

## Steps

1. Write existing dev stub.
2. Run `scaff fix script/dev`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if err := writeScriptDev(d, req.ProjectDir); err != nil {
		return err
	}
	req.Args = []string{"fix", "script/dev"}
	return nil
}
```