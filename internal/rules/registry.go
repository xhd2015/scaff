package rules

var AllFixRules = []string{
	"git/ignore",
	"github/testing-workflow",
	"project/readme",
	"project/license",
	"tests/doctest",
	"project/agents",
	"project/layout/cmd",
	"script/generate",
	"script/install",
	"script/build",
	"script/dev",
	"script/bundle/for-linux",
	"git/hooks",
	"git/hooks/install",
	"github/release",
	"install/via-curl",
	"script/github/release-assets",
}

var DefaultLintRules = []string{
	"git/ignore",
	"github/testing-workflow",
	"project/readme",
	"project/license",
	"tests/doctest",
}
