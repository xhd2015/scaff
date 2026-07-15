package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lessflags "github.com/xhd2015/less-flags"
	"github.com/xhd2015/scaff/docs"
	"github.com/xhd2015/scaff/internal/audit"
	"github.com/xhd2015/scaff/internal/fix"
	"github.com/xhd2015/scaff/internal/model"
	"github.com/xhd2015/scaff/internal/output"
	"github.com/xhd2015/skills/skillcmd"
)

func main() {
	if len(os.Args) < 2 {
		printUsage(os.Stdout)
		os.Exit(0)
	}
	if isHelpArg(os.Args[1]) {
		printUsage(os.Stdout)
		os.Exit(0)
	}
	switch os.Args[1] {
	case "lint":
		os.Exit(runLint(os.Args[2:]))
	case "fix":
		os.Exit(runFix(os.Args[2:]))
	case "rules":
		os.Exit(runRules(os.Args[2:]))
	case "skill":
		os.Exit(runSkill(os.Args[2:]))
	default:
		fmt.Fprintf(os.Stderr, "scaff: unknown command %q\n", os.Args[1])
		printUsage(os.Stderr)
		os.Exit(2)
	}
}

func singleSkill() *skillcmd.SingleSkill {
	return &skillcmd.SingleSkill{
		Name:        docs.Name,
		RootContent: docs.SkillMD,
		TreeFS:      docs.TreeFS,
		Usage:       "scaff skill --install",
		Help:        skillHelp,
	}
}

func runSkill(args []string) int {
	if err := singleSkill().Handle(args); err != nil {
		fmt.Fprintf(os.Stderr, "scaff skill: %v\n", err)
		return 1
	}
	return 0
}

func isHelpArg(arg string) bool {
	return arg == "-h" || arg == "--help" || arg == "help"
}

func runLint(args []string) int {
	var dir string
	var jsonOut bool
	var profile string
	remain, err := lessflags.String("--dir", &dir).
		Bool("--json", &jsonOut).
		String("--profile", &profile).
		Help("-h,--help", lintHelp).
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

func runRules(args []string) int {
	var jsonOut bool
	remain, err := lessflags.Bool("--json", &jsonOut).
		Help("-h,--help", rulesHelp).
		Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaff rules: %v\n", err)
		return 2
	}
	if len(remain) > 0 {
		fmt.Fprintf(os.Stderr, "scaff rules: unexpected arguments: %s\n", strings.Join(remain, " "))
		return 2
	}
	if jsonOut {
		if err := output.PrintRulesJSON(os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "scaff rules: %v\n", err)
			return 1
		}
		return 0
	}
	output.PrintRules(os.Stdout)
	return 0
}

func runFix(args []string) int {
	if len(args) == 0 || isHelpArg(args[0]) {
		fmt.Print(fixHelp)
		return 0
	}
	ruleID := args[0]
	args = args[1:]

	var dir string
	var dryRun bool
	remain, err := lessflags.String("--dir", &dir).
		Bool("--dry-run", &dryRun).
		Help("-h,--help", fixHelp).
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
		fmt.Fprintf(os.Stderr, "available rules: %s\n", output.FormatFixRuleList())
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

const topHelp = `scaff — amend scaffolding to existing projects

Usage:
  scaff lint [options]
  scaff fix <rule> [options]
  scaff rules [options]
  scaff skill --show|--install|--list [options]

Commands:
  lint    audit default rules (git/ignore, github/testing-workflow)
  fix     apply one scaffolding rule
  rules   list lint and fix rules
  skill   show, install, or list embedded multi-topic skill docs

Run scaff <command> --help for command-specific options.
Run scaff skill --help for skill surface.
Run scaff skill --install --help for install flags.
`

const skillHelp = `Usage: scaff skill --show [--header] [<topic-path>]
       scaff skill <topic-path> --show [--header]
       scaff skill --install [OPTIONS] [<dir>]
       scaff skill --list

Show the root SKILL.md index or a nested topic (path/TOPIC.md).
Install copies SKILL.md and nested TOPIC.md topics into agent skill directories.
List prints the skill name and every available topic path.
--help also lists available topics (see below).

Examples:
  scaff skill --show
  scaff skill --show git/ignore
  scaff skill git/ignore --show
  scaff skill --list
  scaff skill --install --dry-run
  scaff skill --install --help

Options:
  --show [--header] [path]   Print skill or topic content (header-only with --header)
  --install [OPTIONS] [dir]  Install skill files (see --install --help)
  --list                     Print skill name and all topic paths
  -h, --help                 Show this help and available topics
`

const rulesHelp = `scaff rules — list lint and fix rules

Usage:
  scaff rules [--json]

Options:
  --json        emit machine-readable JSON
  -h, --help    show this help
`

const lintHelp = `scaff lint — audit project scaffolding

Usage:
  scaff lint [--dir DIR] [--json] [--profile PROFILE]

Options:
  --dir DIR          project directory (default: current directory)
  --json             emit machine-readable JSON report
  --profile PROFILE  go, node, polyglot, or generic (overrides auto-detect)
  -h, --help         show this help
`

const fixHelp = `scaff fix — apply one scaffolding rule

Usage:
  scaff fix <rule> [--dir DIR] [--dry-run]

Run scaff rules for the full rule list.

Options:
  --dir DIR     project directory (default: current directory)
  --dry-run     show planned changes without writing files
  -h, --help    show this help
`

func printUsage(w *os.File) {
	fmt.Fprint(w, topHelp)
}