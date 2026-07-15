## Expected

- Catalog lists **exactly** these eleven slash IDs (set equality; order does not matter).
  Updated from the original P1 ten-ID set to include P5 rule `script/github/release-assets`:
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
  - `script/github/release-assets`
- No dotted legacy IDs, no aliases, no missing rules, no extra rules.
- No duplicate IDs.
- Catalog entry for `script/github/release-assets` is Fix-capable (covered by catalog inventory inclusion; lint optional).

## Errors

- None: `Run` must succeed.

```go
import (
	"sort"
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("nil response")
	}

	// Phase expansion (P5): known-set grows from 10 → 11 with script/github/release-assets.
	want := []string{
		"git/ignore",
		"github/testing-workflow",
		"script/generate",
		"script/install",
		"script/build",
		"script/bundle/for-linux",
		"git/hooks",
		"git/hooks/install",
		"github/release",
		"install/via-curl",
		"script/github/release-assets",
	}

	got := append([]string(nil), resp.IDs...)
	sort.Strings(want)
	sort.Strings(got)

	if len(got) != len(want) {
		t.Fatalf("Catalog ID count = %d, want %d\ngot:  %v\nwant: %v",
			len(got), len(want), got, want)
	}

	seen := make(map[string]int, len(got))
	for _, id := range got {
		seen[id]++
	}
	for id, n := range seen {
		if n != 1 {
			t.Fatalf("Catalog ID %q appears %d times; want unique", id, n)
		}
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Catalog ID set mismatch\ngot:  %s\nwant: %s\nmissing: %v\nextra: %v",
				strings.Join(got, ", "),
				strings.Join(want, ", "),
				setDiff(want, got),
				setDiff(got, want),
			)
		}
	}
}

func setDiff(a, b []string) []string {
	inB := make(map[string]bool, len(b))
	for _, x := range b {
		inB[x] = true
	}
	var out []string
	for _, x := range a {
		if !inB[x] {
			out = append(out, x)
		}
	}
	if out == nil {
		return []string{}
	}
	return out
}
```
