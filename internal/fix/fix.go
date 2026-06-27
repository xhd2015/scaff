package fix

import (
	"fmt"

	"github.com/xhd2015/scaff/internal/model"
	"github.com/xhd2015/scaff/internal/rules"
)

func Apply(project model.Project, ruleID string, dryRun bool) (model.FixResult, error) {
	switch ruleID {
	case "git.ignore":
		return rules.FixGitIgnore(project, dryRun)
	case "github.testing.workflow":
		return rules.FixGitHubTestingWorkflow(project, dryRun)
	case "script.generate":
		return rules.FixScriptGenerate(project, dryRun)
	case "git.hooks":
		return rules.FixGitHooks(project, dryRun)
	case "git.hooks.install":
		return rules.FixGitHooksInstall(project, dryRun)
	default:
		return model.FixResult{}, fmt.Errorf("unknown rule %q", ruleID)
	}
}

func IsKnownRule(ruleID string) bool {
	for _, id := range rules.AllFixRules {
		if id == ruleID {
			return true
		}
	}
	return false
}