package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const (
	githubReleaseMainPath = "script/github/release/main.go"
	githubReleaseLibPath  = "script/github/lib/build_release.go"
)

const githubReleaseMainTemplate = `// usage: go run ./script/github/release [--dry-run]
//
// Proposed behavior (sketch):
//   1. Parse --dry-run / --help flags.
//   2. Dry-run: print tag, planned artifacts, and upload target without writing.
//   3. Live: load credentials, build multi-platform assets, create/upload GitHub Release.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xhd2015/kool/pkgs/github"
	"github.com/xhd2015/kool/pkgs/release"
	"github.com/xhd2015/less-flags"
)

const help = ` + "`" + `
Usage: go run ./script/github/release [--dry-run]

Release __NAME__ to GitHub Releases.

Options:
  --dry-run    print what would be done without actually building or uploading
  -h,--help    show help message
` + "`" + `

func main() {
	if err := handle(); err != nil {
		fmt.Fprintf(os.Stderr, "__NAME__ release: %v\n", err)
		os.Exit(1)
	}
}

func handle() error {
	var dryRun bool
	args, err := lessflags.
		Bool("--dry-run", &dryRun).
		Help("-h,--help", help).
		Parse(os.Args[1:])
	if err != nil {
		return err
	}
	if len(args) > 0 {
		return fmt.Errorf("unrecognized extra args: %s", strings.Join(args, " "))
	}

	if dryRun {
		return handleDryRun()
	}

	creds, err := release.LoadCredentials(".upload-credentials.json")
	if err != nil {
		return err
	}

	result, err := release.BuildRelease("__NAME__", nil, release.DefaultSpecs)
	if err != nil {
		return err
	}

	client := github.NewReleaseClient(creds.Token, creds.Owner, creds.Repo)

	rel, err := client.GetOrCreateRelease(result.Tag)
	if err != nil {
		return fmt.Errorf("failed to get or create release for tag %s: %v", result.Tag, err)
	}

	for _, file := range result.Files {
		if err := client.UploadReleaseAsset(rel.ID, file); err != nil {
			return fmt.Errorf("failed to upload %s: %v", file, err)
		}
		fmt.Printf("Uploaded %s\n", file)
	}
	return nil
}

func handleDryRun() error {
	tag, tagErr := release.GetTag()
	if tagErr != nil {
		fmt.Fprintf(os.Stderr, "[dry-run] warning: %v\n", tagErr)
		tag = "(unknown)"
	}

	creds, credsErr := release.LoadCredentials(".upload-credentials.json")
	if credsErr != nil {
		fmt.Fprintf(os.Stderr, "[dry-run] warning: %v\n", credsErr)
		creds = &release.Credentials{Owner: "__OWNER__", Repo: "__REPO__"}
	}

	fmt.Printf("[dry-run] tag: %s\n", tag)
	for _, spec := range release.DefaultSpecs {
		fmt.Printf("[dry-run] would build: __NAME__-%s-%s-%s\n", tag, spec.OS, spec.Arch)
	}
	fmt.Printf("[dry-run] would upload to %s/%s release (creates if 404)\n", creds.Owner, creds.Repo)
	return nil
}
`

const githubReleaseLibTemplate = `// usage: imported by go run ./script/github/release (shared release helpers)
//
// Proposed behavior (sketch):
//   1. Expose DefaultSpecs for multi-platform release builds.
//   2. BuildRelease runs optional pre-build steps then release.BuildRelease.
//   3. Callers pass specs; name/module placeholders are substituted at scaffold time.
package lib

import (
	"github.com/xhd2015/kool/pkgs/release"
)

var DefaultSpecs = release.DefaultSpecs

func BuildRelease(specs []*release.Spec) (*release.BuildReleaseResult, error) {
	// Add custom pre-build steps here (e.g. frontend build, asset generation)
	return release.BuildRelease("__NAME__", nil, specs)
}
`

func FixGithubRelease(project model.Project, dryRun bool) (model.FixResult, error) {
	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}

	mainPath := filepath.Join(project.Root, githubReleaseMainPath)
	libPath := filepath.Join(project.Root, githubReleaseLibPath)

	mainExists := fileExistsAt(mainPath)
	libExists := fileExistsAt(libPath)

	if mainExists && libExists {
		return model.FixResult{
			RuleID:  "github/release",
			Actions: []string{fmt.Sprintf("%s and %s already exist, nothing to do", githubReleaseMainPath, githubReleaseLibPath)},
		}, nil
	}

	result := model.FixResult{RuleID: "github/release"}
	var actions []string

	if !mainExists {
		if dryRun {
			actions = append(actions, fmt.Sprintf("dry-run: would create %s", githubReleaseMainPath))
		} else {
			if err := os.MkdirAll(filepath.Dir(mainPath), 0o755); err != nil {
				return model.FixResult{}, err
			}
			content := substituteMeta(githubReleaseMainTemplate, meta)
			if err := os.WriteFile(mainPath, []byte(content), 0o644); err != nil {
				return model.FixResult{}, err
			}
			result.Changed = true
			actions = append(actions, fmt.Sprintf("created %s", githubReleaseMainPath))
		}
	}

	if !libExists {
		if dryRun {
			actions = append(actions, fmt.Sprintf("dry-run: would create %s", githubReleaseLibPath))
		} else {
			if err := os.MkdirAll(filepath.Dir(libPath), 0o755); err != nil {
				return model.FixResult{}, err
			}
			content := substituteMeta(githubReleaseLibTemplate, meta)
			if err := os.WriteFile(libPath, []byte(content), 0o644); err != nil {
				return model.FixResult{}, err
			}
			result.Changed = true
			actions = append(actions, fmt.Sprintf("created %s", githubReleaseLibPath))
		}
	}

	result.Actions = actions
	return result, nil
}

func fileExistsAt(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}