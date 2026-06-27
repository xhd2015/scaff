package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lessflags "github.com/xhd2015/less-flags"
	"github.com/xhd2015/scaff/internal/audit"
	"github.com/xhd2015/scaff/internal/model"
	"github.com/xhd2015/scaff/internal/fix"
	"github.com/xhd2015/scaff/internal/output"
	"github.com/xhd2015/scaff/internal/rules"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "lint":
		os.Exit(runLint(os.Args[2:]))
	case "fix":
		os.Exit(runFix(os.Args[2:]))
	default:
		fmt.Fprintf(os.Stderr, "scaff: unknown command %q\n", os.Args[1])
		printUsage()
		os.Exit(2)
	}
}

func runLint(args []string) int {
	var dir string
	var jsonOut bool
	var profile string
	remain, err := lessflags.String("--dir", &dir).
		Bool("--json", &jsonOut).
		String("--profile", &profile).
		Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaff lint: %v\n", err)
		return 2
	}
	if len(remain) > 0 {
		fmt.Fprintf(os.Stderr, "scaff lint: unexpected arguments: %s\n", strings.Join(remain, " "))
		return 2
	}
	project, err := resolveProject(dir, profile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaff lint: %v\n", err)
		return 1
	}
	report := audit.Lint(project)
	if jsonOut {
		if err := output.PrintLintJSON(os.Stdout, report); err != nil {
			fmt.Fprintf(os.Stderr, "scaff lint: %v\n", err)
			return 1
		}
		if audit.HasIssues(report) {
			return 1
		}
		return 0
	}
	output.PrintLintReport(os.Stdout, report)
	if audit.HasIssues(report) {
		return 1
	}
	return 0
}

func runFix(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "scaff fix: requires rule argument")
		printFixUsage()
		return 2
	}
	ruleID := args[0]
	args = args[1:]

	var dir string
	var dryRun bool
	remain, err := lessflags.String("--dir", &dir).
		Bool("--dry-run", &dryRun).
		Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaff fix: %v\n", err)
		return 2
	}
	if len(remain) > 0 {
		fmt.Fprintf(os.Stderr, "scaff fix: unexpected arguments: %s\n", strings.Join(remain, " "))
		return 2
	}
	if !fix.IsKnownRule(ruleID) {
		fmt.Fprintf(os.Stderr, "scaff fix: unknown rule %q\n", ruleID)
		fmt.Fprintf(os.Stderr, "available rules: %s\n", strings.Join(rules.AllFixRules, ", "))
		return 2
	}
	project, err := resolveProject(dir, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaff fix: %v\n", err)
		return 1
	}
	result, err := fix.Apply(project, ruleID, dryRun)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaff fix: %v\n", err)
		return 1
	}
	for _, action := range result.Actions {
		fmt.Println(action)
	}
	return 0
}

func resolveProject(dir, profileOverride string) (model.Project, error) {
	root := "."
	if dir != "" {
		root = dir
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return model.Project{}, err
	}
	return audit.DetectProject(abs, profileOverride)
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "usage: scaff <lint|fix> [options]")
}

func printFixUsage() {
	fmt.Fprintln(os.Stderr, "usage: scaff fix <rule> [--dir DIR] [--dry-run]")
}