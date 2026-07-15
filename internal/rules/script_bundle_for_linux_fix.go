package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const scriptBundleForLinuxPath = "script/bundle/for-linux/main.go"

const scriptBundleForLinuxStub = `// usage: go run ./script/bundle/for-linux
//
// Cross-compiles the binary for linux/amd64.
// The resulting artifact is written to ./app-linux-amd64 in the module root.
//
// Proposed behavior (sketch):
//   1. Parse -o/--output (default app-linux-amd64).
//   2. Cross-compile with GOOS=linux GOARCH=amd64 CGO_ENABLED=0.
//   3. Write the binary under the module root and print its path.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/xhd2015/less-gen/flags"
)

const defaultOutput = "app-linux-amd64"

` + "var help = `\nUsage: go run ./script/bundle/for-linux [options]\n\nCross-compiles the binary for linux/amd64.\n\nOptions:\n  -o, --output PATH   Output binary path (default: ` + defaultOutput + `)\n  -h, --help          Show this help message\n`\n\n" + `func main() {
	if err := Handle(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Handle(args []string) error {
	var output string
	_, err := flags.
		String("-o,--output", &output).
		Help("-h,--help", help).
		Parse(args)
	if err != nil {
		return err
	}
	if output == "" {
		output = defaultOutput
	}

	root, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("resolve working directory: %w", err)
	}
	fmt.Printf("module root: %s\n", root)

	env := append(os.Environ(), "GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0")
	cmd := exec.Command("go", "build", "-o", output, ".")
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	out := filepath.Join(root, output)
	fmt.Printf("\nBundle ready: %s\n", out)
	return nil
}
`

func FixScriptBundleForLinux(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, scriptBundleForLinuxPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "script/bundle/for-linux",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", scriptBundleForLinuxPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "script/bundle/for-linux"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", scriptBundleForLinuxPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(scriptBundleForLinuxStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", scriptBundleForLinuxPath)}
	return result, nil
}