# Scenario

**Feature**: scaff fix github/release

```
# scaffold script/github/release + lib from go.mod metadata
fix executor -> github/release -> release scripts under script/github/
```

## Preconditions

- A Go project fixture is prepared per leaf case (typically `github.com/xhd2015/myapp`).

## Steps

1. Materialize release script file state for the scenario.
2. Run `scaff fix github/release` with optional `--dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoModGitHubScaffold(req.ProjectDir); err != nil {
		return err
	}
	return nil
}

```