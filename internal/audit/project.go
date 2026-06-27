package audit

import (
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

func DetectProject(root string, profileOverride string) (model.Project, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return model.Project{}, err
	}
	profile := detectProfile(abs)
	if profileOverride != "" {
		profile = model.Profile(profileOverride)
	}
	return model.Project{Root: abs, Profile: profile}, nil
}

func detectProfile(root string) model.Profile {
	hasGo := fileExists(filepath.Join(root, "go.mod"))
	hasNode := fileExists(filepath.Join(root, "package.json"))
	switch {
	case hasGo && hasNode:
		return model.ProfilePolyglot
	case hasGo:
		return model.ProfileGo
	case hasNode:
		return model.ProfileNode
	default:
		return model.ProfileGeneric
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}