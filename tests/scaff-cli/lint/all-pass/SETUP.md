# Scenario

**Feature**: scaff lint passes when scaffold is complete

```
# complete default-rule coverage yields exit 0
Project detector -> lint orchestrator -> all rules ok
```

## Preconditions

- The project satisfies both default lint rules.

## Steps

1. Prepare a project with complete `.gitignore` and `test.yml`.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	markLintTree()
	req.Args = []string{"lint"}
	return nil
}

// markAllPassTree keeps hierarchical child packages importing this package live.
func markAllPassTree() {}
```