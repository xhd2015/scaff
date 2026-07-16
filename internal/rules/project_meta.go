package rules

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ProjectMeta struct {
	Module string
	Name   string
	Owner  string
	Repo   string
	Year   string
}

func DetectProjectMeta(root string) (ProjectMeta, error) {
	meta := ProjectMeta{Year: strconv.Itoa(time.Now().Year())}
	base := filepath.Base(root)
	meta.Name = base
	meta.Repo = base
	meta.Owner = "OWNER"

	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		if os.IsNotExist(err) {
			return meta, nil
		}
		return meta, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			meta.Module = strings.TrimSpace(strings.TrimPrefix(line, "module "))
			break
		}
	}
	if meta.Module != "" {
		parts := strings.Split(meta.Module, "/")
		if len(parts) > 0 {
			meta.Name = parts[len(parts)-1]
			meta.Repo = meta.Name
		}
		if len(parts) >= 3 && parts[0] == "github.com" {
			meta.Owner = parts[1]
			meta.Repo = parts[2]
		}
	}
	return meta, nil
}

func substituteMeta(s string, meta ProjectMeta) string {
	r := strings.NewReplacer(
		"__MODULE__", meta.Module,
		"__NAME__", meta.Name,
		"__OWNER__", meta.Owner,
		"__REPO__", meta.Repo,
		"__YEAR__", meta.Year,
	)
	return r.Replace(s)
}
