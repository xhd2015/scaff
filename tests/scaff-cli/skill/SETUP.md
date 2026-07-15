# Scenario

**Feature**: scaff multi-topic skill surface via `scaff skill`

```
# skillcmd SingleSkill + embedded docs TreeFS
User -> scaff skill [--list|--show|--install|--help] -> stdout/stderr/exit
# nested path/TOPIC.md topics under docs/; root docs/SKILL.md
```

## Preconditions

- The `scaff` binary and temp working directory are ready from the root setup.
- Skill content is embedded in the product (`docs/` + `skillcmd.SingleSkill`); no project fixture is required for skill commands.

## Steps

1. Descendant setups set `req.Args` for the skill action under test.
2. `Run` executes the built `scaff` binary with those args from `req.RunDir`.

## Context

- CLI surface is **skill only** — no top-level `install` or `topics` aliases.
- Modes: exactly one of `--show` | `--install` | `--list` (skillcmd); `--header` only with `--show`.

```go
func Setup(t *testing.T, req *Request) error {
	if req.RunDir == "" {
		req.RunDir = req.ProjectDir
	}
	return nil
}
```
