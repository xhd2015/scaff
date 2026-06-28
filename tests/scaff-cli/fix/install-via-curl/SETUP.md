# Scenario

**Feature**: scaff fix install.via.curl

```
# scaffold install-via-curl.sh at repo root from go.mod metadata
fix executor -> install.via.curl -> curl installer script
```

## Preconditions

- A Go project fixture is prepared per leaf case (typically `github.com/xhd2015/myapp`).

## Steps

1. Materialize installer script state for the scenario.
2. Run `scaff fix install.via.curl` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}
```