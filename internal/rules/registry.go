package rules

var AllFixRules = []string{
	"git.ignore",
	"github.testing.workflow",
	"script.generate",
	"script.install",
	"script.build",
	"script.bundle.for-linux",
	"git.hooks",
	"git.hooks.install",
	"github.release",
	"install.via.curl",
}

var DefaultLintRules = []string{
	"git.ignore",
	"github.testing.workflow",
}