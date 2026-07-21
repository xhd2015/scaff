# Scenario

**Feature**: show `git/ignore` with path before `--show`

```
# path-before-flag order (skillcmd both orders)
scaff skill git/ignore --show -> same topic body as --show git/ignore
```

## Preconditions

- Topic `git/ignore` is embedded.
- skillcmd accepts topic path before `--show`.

## Steps

1. Run `scaff skill git/ignore --show`.

```go
func Setup(t *testing.T, req *Request) error {
	markShowTopicTree()
	markShowTree()
	markSkillTree()
	req.Args = []string{"skill", "git/ignore", "--show"}
	return nil
}
```
