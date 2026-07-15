# Scenario

**Feature**: every Catalog ID is slash form

```
# each RuleInfo.ID must look like path segments: a/b or a/b/c (no dots)
Catalog IDs -> all match slash form
```

## Preconditions

- Catalog is non-empty (parent).
- Slash form means: contains at least one `/` and contains no `.` character.

## Steps

1. Run collects all Catalog IDs.
2. Assert every ID is slash form.

```go
import (
	"fmt"
	"testing"

	"github.com/xhd2015/scaff/internal/rules"
)

func Setup(t *testing.T, req *Request) error {
	// Leaf will assert slash form on every ID returned by Run.
	for i, r := range rules.Catalog {
		if r.ID == "" {
			return fmt.Errorf("rules.Catalog[%d] has empty ID", i)
		}
	}
	return nil
}
```
