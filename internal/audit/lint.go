package audit

import (
	"github.com/xhd2015/scaff/internal/model"
	"github.com/xhd2015/scaff/internal/rules"
)

func Lint(project model.Project) model.LintReport {
	report := model.LintReport{Project: project}
	for _, ruleID := range rules.DefaultLintRules {
		switch ruleID {
		case "git.ignore":
			report.Results = append(report.Results, rules.LintGitIgnore(project))
		case "github.testing.workflow":
			report.Results = append(report.Results, rules.LintGitHubTestingWorkflow(project))
		}
	}
	return report
}

func HasIssues(r model.LintReport) bool {
	for _, result := range r.Results {
		if result.Status != model.RuleOK {
			return true
		}
	}
	return false
}

func IssueCount(r model.LintReport) int {
	n := 0
	for _, result := range r.Results {
		if result.Status != model.RuleOK {
			n++
		}
	}
	return n
}