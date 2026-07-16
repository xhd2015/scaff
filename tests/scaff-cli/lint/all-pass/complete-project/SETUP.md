# Scenario

**Feature**: lint passes on complete scaffold

```
# complete gitignore + test.yml + README + LICENSE + doctest -> all default rules ok
lint orchestrator -> exit 0 all good
```

## Preconditions

- Project has `go.mod`, complete Go `.gitignore`, `test.yml`, `README.md`, `LICENSE`, and doctest tree.

## Steps

1. Write complete project fixtures including README, LICENSE, and doctest tree.
2. Run `scaff lint`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeGoMod(req.ProjectDir); err != nil {
		return err
	}
	if err := writeCompleteGoGitignore(req.ProjectDir); err != nil {
		return err
	}
	if err := writeTestWorkflow(req.ProjectDir); err != nil {
		return err
	}
	if err := writeREADME(req.ProjectDir, "# example.com/app\n\n## Usage\n\n...\n"); err != nil {
		return err
	}
	if err := writeLICENSE(req.ProjectDir, "MIT License\n\nCopyright (c) 2020-present, OWNER\n\nPermission is hereby granted, free of charge...\n"); err != nil {
		return err
	}
	if err := writeDoctestTree(req.ProjectDir, "app"); err != nil {
		return err
	}
	req.Args = []string{"lint"}
	return nil
}
```
