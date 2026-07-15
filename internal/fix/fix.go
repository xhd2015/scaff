package fix

import (
	"fmt"

	"github.com/xhd2015/scaff/internal/model"
	"github.com/xhd2015/scaff/internal/rules"
)

func Apply(project model.Project, ruleID string, dryRun bool) (model.FixResult, error) {
	switch ruleID {
	case "git/ignore":
		return rules.FixGitIgnore(project, dryRun)
	case "github/testing-workflow":
		return rules.FixGitHubTestingWorkflow(project, dryRun)
	case "script/generate":
		return rules.FixScriptGenerate(project, dryRun)
	case "script/install":
		return rules.FixScriptInstall(project, dryRun)
	case "script/build":
		return rules.FixScriptBuild(project, dryRun)
	case "script/bundle/for-linux":
		return rules.FixScriptBundleForLinux(project, dryRun)
	case "git/hooks":
		return rules.FixGitHooks(project, dryRun)
	case "git/hooks/install":
		return rules.FixGitHooksInstall(project, dryRun)
	case "github/release":
		return rules.FixGithubRelease(project, dryRun)
	case "install/via-curl":
		return rules.FixInstallViaCurl(project, dryRun)
	case "script/github/release-assets":
		return rules.FixScriptGithubReleaseAssets(project, dryRun)
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
