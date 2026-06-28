# Scenario

**Feature**: fix install.via.curl no-op when installer exists

```
# install-via-curl.sh present -> no overwrite
install.via.curl fix -> nothing to do
```

## Preconditions

- `install-via-curl.sh` already exists with custom marker content.

## Steps

1. Write custom installer script.
2. Run `scaff fix install.via.curl`.

```go
func Setup(t *testing.T, req *Request) error {
	if err := writeInstallViaCurl(req.ProjectDir, "#!/usr/bin/env bash\n# CUSTOM_INSTALLER\n"); err != nil {
		return err
	}
	req.Args = []string{"fix", "install.via.curl"}
	return nil
}
```