package output

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/xhd2015/scaff/internal/audit"
	"github.com/xhd2015/scaff/internal/model"
)

func PrintLintReport(w io.Writer, report model.LintReport) {
	if !audit.HasIssues(report) {
		passing := len(report.Results)
		fmt.Fprintf(w, "scaff lint: all good (%d/%d rules passing)\n", passing, passing)
		return
	}
	fmt.Fprintf(w, "scaff lint: %d issue(s) in .\n\n", audit.IssueCount(report))
	for _, result := range report.Results {
		if result.Status == model.RuleOK {
			continue
		}
		fmt.Fprintf(w, "  %-28s %s\n", result.ID, result.Message)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "hint: scaff fix <rule>")
}

func PrintLintJSON(w io.Writer, report model.LintReport) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(report)
}