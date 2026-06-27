package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/xhd2015/scaff/internal/rules"
)

type RulesReport struct {
	Lint []rules.RuleInfo `json:"lint"`
	Fix  []rules.RuleInfo `json:"fix"`
}

func PrintRulesJSON(w io.Writer) error {
	report := RulesReport{
		Lint: rules.LintRules(),
		Fix:  rules.FixRules(),
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(report)
}

func PrintRules(w io.Writer) {
	fmt.Fprintln(w, "lint:")
	for _, r := range rules.LintRules() {
		fmt.Fprintf(w, "  %-28s %s\n", r.ID, r.Description)
	}
	fmt.Fprintln(w, "fix:")
	for _, r := range rules.FixRules() {
		fmt.Fprintf(w, "  %-28s %s\n", r.ID, r.Description)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "usage:")
	fmt.Fprintln(w, "  scaff lint")
	fmt.Fprintln(w, "  scaff fix <rule>")
}

func FormatFixRuleList() string {
	ids := make([]string, len(rules.AllFixRules))
	copy(ids, rules.AllFixRules)
	return strings.Join(ids, ", ")
}