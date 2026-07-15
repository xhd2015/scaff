## Expected

- No Catalog entry `ID` contains the character `.`.

## Errors

- None: `Run` must succeed.

```go
import (
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
	if len(resp.IDs) == 0 {
		t.Fatal("expected at least one Catalog ID")
	}
	for _, id := range resp.IDs {
		if strings.Contains(id, ".") {
			t.Fatalf("Catalog ID %q must not contain '.'", id)
		}
	}
}
```
