# Scenario

**Feature**: skill help lists actions and available topics

```
scaff skill --help -> mentions --show, --install, --list + Available topics
```

## Preconditions

- Parent sets `skill --help`.

## Steps

1. Assert usage markers and topic index.

```go
func Setup(t *testing.T, req *Request) error {
	if len(req.Args) == 0 {
		req.Args = []string{"skill", "--help"}
	}
	return nil
}
```
