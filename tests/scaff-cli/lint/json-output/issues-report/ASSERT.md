## Expected

- Exit code is `1`.
- stdout is valid JSON describing a lint report.
- JSON contains results for `git.ignore` and `github.testing.workflow`.

## Exit Code

- `1`

```go
import (
	"encoding/json"
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 1 {
		t.Fatalf("exit code = %d, want 1\n%s", resp.ExitCode, resp.Combined)
	}
	var report struct {
		Project struct {
			Root    string `json:"Root"`
			Profile string `json:"Profile"`
		} `json:"Project"`
		Results []struct {
			ID      string `json:"ID"`
			Status  string `json:"Status"`
			Message string `json:"Message"`
		} `json:"Results"`
	}
	if err := json.Unmarshal([]byte(strings.TrimSpace(resp.Stdout)), &report); err != nil {
		t.Fatalf("stdout is not valid JSON: %v\n%s", err, resp.Stdout)
	}
	found := map[string]bool{}
	for _, r := range report.Results {
		found[r.ID] = true
	}
	for _, id := range []string{"git.ignore", "github.testing.workflow"} {
		if !found[id] {
			t.Fatalf("expected result for %q in JSON: %s", id, resp.Stdout)
		}
	}
}
```