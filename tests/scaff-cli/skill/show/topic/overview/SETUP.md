# Scenario

**Feature**: show `overview` topic (product model)

```
# path overview -> docs/overview/TOPIC.md
scaff skill --show overview -> name: scaff/overview
```

## Preconditions

- Topic `overview` is embedded.

## Steps

1. Run `scaff skill --show overview`.

```go
func Setup(t *testing.T, req *Request) error {
	markShowTopicTree()
	markShowTree()
	markSkillTree()
	req.Args = []string{"skill", "--show", "overview"}
	return nil
}
```
