# Scenario

**Feature**: fix project/readme --dry-run

```
# --dry-run reports would-create without writing README.md
project/readme fix --dry-run -> preview only
```

## Preconditions

- Project has no `README.md`.

## Steps

1. Run `scaff fix project/readme --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markProjectReadmeTree()
	markFixTree()
	req.Args = []string{"fix", "project/readme", "--dry-run"}
	return nil
}
```