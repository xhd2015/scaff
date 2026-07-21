# Scenario

**Feature**: show unknown topic path errors

```
# missing TreeFS path -> non-zero exit, stderr indicates unknown/missing topic
scaff skill --show not-a-real-topic -> error
```

## Preconditions

- Topic path `not-a-real-topic` is not part of the skill inventory.

## Steps

1. Run `scaff skill --show not-a-real-topic`.

```go
func Setup(t *testing.T, req *Request) error {
	markShowTree()
	markSkillTree()
	req.Args = []string{"skill", "--show", "not-a-real-topic"}
	return nil
}
```
