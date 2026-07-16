package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const licensePath = "LICENSE"

func LintProjectLicense(project model.Project) model.RuleResult {
	result := model.RuleResult{
		ID:     "project/license",
		Paths:  []string{licensePath},
		Status: model.RuleOK,
	}
	path := filepath.Join(project.Root, licensePath)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			result.Status = model.RuleMissing
			result.Message = fmt.Sprintf("missing: %s", licensePath)
			return result
		}
		result.Status = model.RuleMissing
		result.Message = err.Error()
		return result
	}
	result.Message = "present"
	return result
}

func FixProjectLicense(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, licensePath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "project/license",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", licensePath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}
	content := licenseTemplate(meta)
	result := model.FixResult{RuleID: "project/license"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", licensePath)}
		return result, nil
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", licensePath)}
	return result, nil
}

func licenseTemplate(meta ProjectMeta) string {
	return substituteMeta(`MIT License

Copyright (c) __YEAR__-present, __OWNER__

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`, meta)
}
