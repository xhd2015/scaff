# Scenario

**Feature**: lint --dir audits nested project root

```
# parent dir contains nested app/ with go.mod
scaff lint --dir app -> rules evaluated under app/
```

## Preconditions

- Parent directory contains subdirectory `app/` with `go.mod` only.

## Steps

1. Create `app/go.mod` under the project directory.
2. Run `scaff lint --dir app` from the parent.

```go
import "path/filepath"

func Setup(t *testing.T, req *Request) error {
	appDir := filepath.Join(req.ProjectDir, "app")
	if err := writeGoMod(appDir); err != nil {
		return err
	}
	req.Args = []string{"lint", "--dir", "app"}
	return nil
}
```