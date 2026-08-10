# Scenario

**Feature**: fix script/install creates stub

```
# missing script/install/install.go -> create install helper stub
script/install fix -> script/install/install.go
```

## Preconditions

- `script/install/install.go` does not exist.

## Steps

1. Run `scaff fix script/install`.

```go
func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	req.Args = []string{"fix", "script/install"}
	return nil
}
```