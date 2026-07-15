# Scenario

**Feature**: Catalog ID set equals the authoritative slash map

```
# exactly the 11 slash IDs — no missing, no extra, no aliases
Catalog IDs as set == known slash set (P1 map + script/github/release-assets)
```

## Preconditions

- Catalog is non-empty (parent).
- Authoritative set (exactly these eleven, order-independent):
  - `git/ignore`
  - `github/testing-workflow`
  - `script/generate`
  - `script/install`
  - `script/build`
  - `script/bundle/for-linux`
  - `git/hooks`
  - `git/hooks/install`
  - `github/release`
  - `install/via-curl`
  - `script/github/release-assets` (P5)

## Steps

1. Run collects all Catalog IDs.
2. Assert multiset/set equality with the known slash IDs (no duplicates, no extras, no missing).

```go
import (
	"fmt"
	"testing"

	"github.com/xhd2015/scaff/internal/rules"
)

func Setup(t *testing.T, req *Request) error {
	// Leaf compares Catalog IDs to the fixed slash map (11 rules after P5).
	if len(rules.Catalog) == 0 {
		return fmt.Errorf("rules.Catalog is empty")
	}
	return nil
}
```
