# Scenario

**Feature**: no Catalog ID contains a dot

```
# exit criterion: No Catalog entry ID contains '.'
Catalog IDs -> none contain '.'
```

## Preconditions

- Catalog is non-empty (parent).
- Dotted legacy IDs such as `git.ignore` or `script.bundle.for-linux` must not appear.

## Steps

1. Run collects all Catalog IDs.
2. Assert no ID contains `.`.

```go
import (
	"fmt"
	"testing"

	"github.com/xhd2015/scaff/internal/rules"
)

func Setup(t *testing.T, req *Request) error {
	// Leaf forbids '.' in any Catalog ID (independent of slash presence).
	if len(rules.Catalog) == 0 {
		return fmt.Errorf("rules.Catalog is empty")
	}
	return nil
}
```
