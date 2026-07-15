# scaff rule ID Catalog and Fix Dispatch Tests

Doc-style tests for canonical **slash-form** rule IDs on `rules.Catalog` and
on the `scaff fix` CLI contract.

## Version
0.0.2

# DSN (Domain Specific Notion)

The **rule Catalog** is the public inventory of scaff lint/fix rules. Each
entry is a `rules.RuleInfo` with an `ID` string used by humans and by CLI
surfaces (`scaff fix <id>`, `scaff rules`).

**Slash form** is the only allowed ID shape: path-like segments joined by `/`
(e.g. `git/ignore`, `script/bundle/for-linux`). Dotted historical IDs
(e.g. `git.ignore`) are forbidden. There are no backward-compat aliases —
Catalog lists exactly the new IDs, and `scaff fix` accepts only those IDs.

**Authoritative slash map** (old → new):

| Old (forbidden) | New (required) |
|-----------------|----------------|
| `git.ignore` | `git/ignore` |
| `github.testing.workflow` | `github/testing-workflow` |
| `script.generate` | `script/generate` |
| `script.install` | `script/install` |
| `script.build` | `script/build` |
| `script.bundle.for-linux` | `script/bundle/for-linux` |
| `git.hooks` | `git/hooks` |
| `git.hooks.install` | `git/hooks/install` |
| `github.release` | `github/release` |
| `install.via.curl` | `install/via-curl` |

Plus P5 addition (no dotted form): `script/github/release-assets`.

**Known set** is exactly those **eleven** slash IDs (ten mapped +
`script/github/release-assets`) — no missing entries, no extras.

**Catalog leaves** read `rules.Catalog` in-process.

**Fix leaves** build the real `scaff` binary and run `scaff fix <id> [--dry-run]`
against an isolated temp project. Slash IDs must be accepted; dotted IDs must
fail as unknown rule (non-zero exit). FixResult / dispatch wiring must use
slash form end-to-end (no dotted aliases).

## Decision Tree

```
tests/scaff-rule-ids/                         [Rule ID surface]
│
├── catalog/                                  public rules.Catalog (P1)
│   ├── all-slash-ids/                        every ID is slash form (has /, no .)
│   ├── no-dot-ids/                           no ID contains '.'
│   └── known-set/                            set equals the 11 slash IDs (incl. release-assets)
│
└── fix/                                      scaff fix CLI (P2)
    ├── slash-id-accepted/                    fix git/ignore --dry-run → exit 0
    ├── dotted-id-rejected/                   fix git.ignore → unknown, exit ≠ 0
    ├── slash-github-release/                 fix github/release --dry-run → exit 0
    └── dotted-github-release-rejected/       fix github.release → unknown, exit ≠ 0
```

## Test Index

| Leaf | Description |
|------|-------------|
| `catalog/all-slash-ids` | Every Catalog ID uses slash form (contains `/`, does not contain `.`) |
| `catalog/no-dot-ids` | No Catalog entry `ID` contains `.` |
| `catalog/known-set` | Catalog ID set equals exactly the 11 authoritative slash IDs |
| `fix/slash-id-accepted` | `scaff fix git/ignore --dry-run` exits 0 |
| `fix/dotted-id-rejected` | `scaff fix git.ignore` fails as unknown rule (exit ≠ 0) |
| `fix/slash-github-release` | `scaff fix github/release --dry-run` exits 0 |
| `fix/dotted-github-release-rejected` | `scaff fix github.release` fails as unknown rule (exit ≠ 0) |

## How to Run

```sh
doctest vet ./tests/scaff-rule-ids
doctest test -v ./tests/scaff-rule-ids
```

```go
import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/xhd2015/scaff/internal/rules"
)

// Request selects catalog inspection (empty Args) or CLI fix (Args + binary).
type Request struct {
	Args       []string
	ProjectDir string
	RunDir     string
	ScaffBin   string
}

// Response carries Catalog IDs and/or CLI process outcome.
type Response struct {
	IDs      []string
	Stdout   string
	Stderr   string
	Combined string
	ExitCode int
}

func Run(t *testing.T, req *Request) (*Response, error) {
	// CLI mode: fix (or other) command via built binary.
	if len(req.Args) > 0 {
		if req.ScaffBin == "" {
			t.Fatal("CLI mode requires ScaffBin")
		}
		runDir := req.ProjectDir
		if req.RunDir != "" {
			runDir = req.RunDir
		}
		cmd := exec.Command(req.ScaffBin, req.Args...)
		cmd.Dir = runDir
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		runErr := cmd.Run()
		exitCode := 0
		if runErr != nil {
			if exitErr, ok := runErr.(*exec.ExitError); ok {
				exitCode = exitErr.ExitCode()
			} else {
				return nil, runErr
			}
		}
		out := stdout.String()
		errOut := stderr.String()
		combined := out
		if errOut != "" {
			if combined != "" {
				combined += "\n"
			}
			combined += errOut
		}
		return &Response{
			Stdout:   out,
			Stderr:   errOut,
			Combined: combined,
			ExitCode: exitCode,
		}, nil
	}

	// Catalog mode: package-level ID inventory.
	ids := make([]string, 0, len(rules.Catalog))
	for _, r := range rules.Catalog {
		ids = append(ids, r.ID)
	}
	return &Response{IDs: ids}, nil
}
```
