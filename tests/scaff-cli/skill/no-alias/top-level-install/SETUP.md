# Scenario

**Feature**: `scaff install` is not a skill-install alias

```
# top-level install must not invoke skill --install
scaff install -> unknown command, exit non-zero
```

## Preconditions

- No top-level `install` command is registered.

## Steps

1. Run `scaff install --dry-run`.

```go
func Setup(t *testing.T, req *Request) error {
	markNoAliasTree()
	markSkillTree()
	req.Args = []string{"install", "--dry-run"}
	return nil
}
```
