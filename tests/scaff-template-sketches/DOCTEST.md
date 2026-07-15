# scaff template behavior sketch Tests

Doc-style tests that scaffolded **Go templates** include a top-of-file
**Proposed behavior** sketch after `scaff fix <rule>` on an empty project.

## Version
0.0.2

# DSN (Domain Specific Notion)

**Scaffold Go templates** are the embedded stub sources written by fix rules
under `internal/rules` into project paths such as `script/build/build.go`,
`script/generate/main.go`, and `script/github/release/main.go`.

Every such template must open with a usage comment and a behavior sketch:

```
// usage: ...
//
// Proposed behavior (sketch):
//   1. ...
```

(or equivalent wording that includes the phrase **Proposed behavior**).

**Empty project** means a temp directory with only a minimal `go.mod` (and for
release, a GitHub module path) and **no** pre-existing scaffold files. Fix
creates the Go file from the template.

These leaves are **Classic TDD**: current stubs may only have `// usage:` lines
without the sketch; asserts require `Proposed behavior` in the generated file.

Implementer must add sketches to **all** Go-writing templates; this suite
covers a representative sample of three distinct rules.

## Decision Tree

```
tests/scaff-template-sketches/                [fix → generated .go sketch]
│
└── templates/                                empty project + scaff fix
    ├── script-build-has-sketch/              script/build/build.go has Proposed behavior
    ├── github-release-has-sketch/            script/github/release/main.go has sketch
    └── script-generate-has-sketch/           script/generate/main.go has sketch
```

## Test Index

| Leaf | Description |
|------|-------------|
| `templates/script-build-has-sketch` | After `scaff fix script/build`, `script/build/build.go` contains `Proposed behavior` |
| `templates/github-release-has-sketch` | After `scaff fix github/release`, release `main.go` contains `Proposed behavior` |
| `templates/script-generate-has-sketch` | After `scaff fix script/generate`, generate `main.go` contains `Proposed behavior` |

## How to Run

```sh
doctest vet ./tests/scaff-template-sketches
doctest test -v ./tests/scaff-template-sketches
```

```go
import (
	"bytes"
	"os/exec"
	"testing"
)

type Request struct {
	Args       []string
	ProjectDir string
	RunDir     string
	ScaffBin   string
}

type Response struct {
	Stdout   string
	Stderr   string
	Combined string
	ExitCode int
}

func Run(t *testing.T, req *Request) (*Response, error) {
	if req.ScaffBin == "" {
		t.Fatal("ScaffBin required")
	}
	if len(req.Args) == 0 {
		t.Fatal("Args required")
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
```
