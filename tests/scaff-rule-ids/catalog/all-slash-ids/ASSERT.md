## Expected

- Every Catalog ID contains at least one `/`.
- Every Catalog ID contains no `.` character.
- Together these define **slash form** for scaff rule IDs.

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
		if !strings.Contains(id, "/") {
			t.Fatalf("Catalog ID %q is not slash form: missing '/'", id)
		}
		if strings.Contains(id, ".") {
			t.Fatalf("Catalog ID %q is not slash form: contains '.'", id)
		}
	}
}
```
