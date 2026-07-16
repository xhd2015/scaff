# Scenario

**Feature**: shared setup for __NAME__ CLI tests

```
# shared doctest harness for __NAME__
__NAME__ CLI -> doctest harness -> assertions
```

```go
func Setup(t *testing.T, req *Request) error {
	t.Helper()
	return nil
}
```