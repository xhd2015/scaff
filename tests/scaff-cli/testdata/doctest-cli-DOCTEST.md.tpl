# __NAME__ CLI Tests

## Version
0.0.2

# DSN (Domain Specific Notion)

The **__NAME__ CLI** is exercised by doc-style tests.

## How to Run

```sh
doctest vet ./tests/__NAME__-cli
doctest test ./tests/__NAME__-cli
```

```go
type Request struct{}

type Response struct{}

func Run(t *testing.T, req *Request) (*Response, error) {
	return &Response{}, nil
}
```