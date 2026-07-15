# Scenario

**Feature**: inspect the public `rules.Catalog` ID list

```
# grouping: all leaves target Catalog.IDs (not CLI, not FixRules dispatch)
rules.Catalog -> []ID
```

## Preconditions

- Root setup confirmed Catalog is importable.
- Assertions apply to every entry's `ID` field only (Lint/Fix flags unchanged by this suite).

## Steps

1. Require Catalog to have at least one entry so empty inventory fails early.
2. Leaves specialize which ID invariant they assert after `Run`.

```go
import (
	"fmt"
	"testing"

	"github.com/xhd2015/scaff/internal/rules"
)

func Setup(t *testing.T, req *Request) error {
	if len(rules.Catalog) == 0 {
		return fmt.Errorf("rules.Catalog is empty")
	}
	return nil
}
```
