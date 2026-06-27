package rules

import "github.com/xhd2015/scaff/internal/model"

var (
	universalPatterns = []string{
		".DS_Store",
		".vscode/",
		"*.swp",
		"*~",
	}
	goPatterns = []string{
		"bin/",
		"*.test",
		"coverage.out",
	}
	nodePatterns = []string{
		"node_modules/",
		"dist/",
		".env",
		"build/",
		".next/",
	}
)

func PatternsForProfile(profile model.Profile) []string {
	patterns := append([]string{}, universalPatterns...)
	switch profile {
	case model.ProfileGo:
		patterns = append(patterns, goPatterns...)
	case model.ProfileNode:
		patterns = append(patterns, nodePatterns...)
	case model.ProfilePolyglot:
		patterns = append(patterns, goPatterns...)
		patterns = append(patterns, nodePatterns...)
	}
	return patterns
}