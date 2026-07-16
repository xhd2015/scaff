package rules

type RuleInfo struct {
	ID          string
	Description string
	Lint        bool
	Fix         bool
}

var Catalog = []RuleInfo{
	{
		ID:          "git/ignore",
		Description: "common .gitignore patterns for project profile",
		Lint:        true,
		Fix:         true,
	},
	{
		ID:          "github/testing-workflow",
		Description: ".github/workflows/test.yml with go test and doctest",
		Lint:        true,
		Fix:         true,
	},
	{
		ID:          "project/readme",
		Description: "root README.md with install instructions for project profile",
		Lint:        true,
		Fix:         true,
	},
	{
		ID:          "project/license",
		Description: "root MIT LICENSE with year and owner metadata",
		Lint:        true,
		Fix:         true,
	},
	{
		ID:          "tests/doctest",
		Description: "tests/<name>-cli/DOCTEST.md and SETUP.md doctest harness",
		Lint:        true,
		Fix:         true,
	},
	{
		ID:          "project/agents",
		Description: "root AGENTS.md with build and test instructions",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "project/layout/cmd",
		Description: "cmd/<name>/main.go CLI entry from module name",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "script/generate",
		Description: "script/generate/main.go no-op stub",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "script/install",
		Description: "script/install/install.go build-then-install helper",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "script/build",
		Description: "script/build/build.go native go build helper",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "script/dev",
		Description: "script/dev/main.go go run . --dev wrapper",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "script/bundle/for-linux",
		Description: "script/bundle/for-linux/main.go linux/amd64 bundle helper",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "git/hooks",
		Description: "script/git-hooks runner (install, pre-commit, pre-push)",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "git/hooks/install",
		Description: "install scaff hooks into .git/hooks/",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "git/pre-commit",
		Description: "script/git/pre-commit ensure paths + git add (git-hooks install comment)",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "github/release",
		Description: "script/github/release and lib helper for GitHub releases",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "install/via-curl",
		Description: "install-via-curl.sh curl installer at repo root",
		Lint:        false,
		Fix:         true,
	},
	{
		ID:          "script/github/release-assets",
		Description: "script/github/release-assets/main.go gh release asset helper",
		Lint:        false,
		Fix:         true,
	},
}

func LintRules() []RuleInfo {
	var out []RuleInfo
	for _, r := range Catalog {
		if r.Lint {
			out = append(out, r)
		}
	}
	return out
}

func FixRules() []RuleInfo {
	var out []RuleInfo
	for _, r := range Catalog {
		if r.Fix {
			out = append(out, r)
		}
	}
	return out
}
