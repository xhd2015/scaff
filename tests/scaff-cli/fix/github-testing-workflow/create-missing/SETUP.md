# Scenario

**Feature**: fix creates test.yml workflow

```
# missing test.yml -> create with go test + doctest steps
github.testing.workflow fix -> embedded template
```

## Preconditions

- Project has `go.mod` but no `test.yml`.

## Steps

1. Run `scaff fix github.testing.workflow`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "github.testing.workflow"}
	return nil
}
```