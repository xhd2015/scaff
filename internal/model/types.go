package model

type Profile string

const (
	ProfileGo       Profile = "go"
	ProfileNode     Profile = "node"
	ProfilePolyglot Profile = "polyglot"
	ProfileGeneric  Profile = "generic"
)

type Project struct {
	Root    string
	Profile Profile
}

type RuleStatus string

const (
	RuleOK      RuleStatus = "ok"
	RuleMissing RuleStatus = "missing"
	RulePartial RuleStatus = "partial"
)

type RuleResult struct {
	ID      string
	Status  RuleStatus
	Message string
	Paths   []string
}

type LintReport struct {
	Project Project
	Results []RuleResult
}

type FixResult struct {
	RuleID  string
	Actions []string
	Changed bool
}