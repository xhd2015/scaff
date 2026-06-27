# Scenario

**Feature**: scaff lint profile override

```
# --profile overrides auto-detected project profile
Project detector <- --profile flag -> git.ignore pattern set
```

## Preconditions

- Profile detection or override affects expected `git.ignore` patterns.

## Steps

1. Prepare a project fixture for the target profile scenario.
2. Run `scaff lint` with `--profile` when overriding auto-detect.

```go
func Setup(t *testing.T, req *Request) error {
	if len(req.Args) == 0 {
		req.Args = []string{"lint"}
	}
	return nil
}
```