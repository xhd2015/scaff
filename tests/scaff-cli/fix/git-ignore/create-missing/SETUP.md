# Scenario

**Feature**: fix git/ignore creates .gitignore

```
# no .gitignore -> create with full Go profile patterns
git/ignore fix -> new .gitignore with universal + Go patterns
```

## Preconditions

- Project has `go.mod` but no `.gitignore`.

## Steps

1. Ensure `.gitignore` is absent.
2. Run `scaff fix git/ignore`.

```go
func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"fix", "git/ignore"}
	return nil
}
```