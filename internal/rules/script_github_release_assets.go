package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const scriptGithubReleaseAssetsPath = "script/github/release-assets/main.go"

// scriptGithubReleaseAssetsStub is the scaffold written by fix script/github/release-assets.
// Generic release-asset helper: pack a directory and optionally upload via gh
// (opt-in --upload; clobber same-named assets). Not tied to a single product.
const scriptGithubReleaseAssetsStub = `// usage: go run ./script/github/release-assets [options]
//
// Proposed behavior (sketch):
//   1. Parse --dir (asset source directory) and optional --tag / --title.
//   2. Pack / list files under --dir as release-ready assets.
//   3. Without --upload: print the plan only (safe default; upload is opt-in).
//   4. With --upload: use gh to create/update the release and clobber assets.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fs := flag.NewFlagSet("script/github/release-assets", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	dir := fs.String("dir", "", "directory of assets to pack / upload (required)")
	tag := fs.String("tag", "", "release tag (default: current git tag or \"latest\")")
	title := fs.String("title", "", "release title (default: tag)")
	upload := fs.Bool("upload", false, "opt-in: upload assets with gh (create/clobber release assets)")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, ` + "`" + `Usage: go run ./script/github/release-assets [options]

Pack a directory of release assets and optionally upload them to a GitHub
Release via the gh CLI. Upload is opt-in so inspection is the default.

Options:
  --dir DIR      directory of assets to pack / upload (required)
  --tag TAG      release tag (default: current git tag or "latest")
  --title TEXT   release title (default: tag)
  --upload       opt-in: upload assets with gh (create/clobber release assets)
  -h, --help     show this help message
` + "`" + `)
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	}

	if err := run(*dir, *tag, *title, *upload); err != nil {
		fmt.Fprintf(os.Stderr, "script/github/release-assets: %v\n", err)
		os.Exit(1)
	}
}

func run(dir, tag, title string, upload bool) error {
	if dir == "" {
		return fmt.Errorf("--dir is required (directory of assets to pack)")
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	info, err := os.Stat(absDir)
	if err != nil {
		return fmt.Errorf("asset dir: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("--dir must be a directory: %s", absDir)
	}

	entries, err := os.ReadDir(absDir)
	if err != nil {
		return err
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		files = append(files, filepath.Join(absDir, e.Name()))
	}
	if len(files) == 0 {
		return fmt.Errorf("no files found under --dir %s", absDir)
	}

	if tag == "" {
		tag = "latest"
		if out, err := exec.Command("git", "describe", "--tags", "--exact-match").Output(); err == nil {
			tag = strings.TrimSpace(string(out))
		}
	}
	if title == "" {
		title = tag
	}

	fmt.Printf("asset dir: %s\n", absDir)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("title: %s\n", title)
	fmt.Println("assets:")
	for _, f := range files {
		fmt.Printf("  %s\n", f)
	}

	if !upload {
		fmt.Println("upload skipped (pass --upload to publish via gh; clobbers same-named assets)")
		return nil
	}

	// Opt-in upload: ensure release exists, then upload each asset (clobber).
	if err := runGH("release", "view", tag); err != nil {
		fmt.Printf("creating release %s...\n", tag)
		if err := runGH("release", "create", tag, "--title", title, "--notes", title); err != nil {
			return fmt.Errorf("gh release create: %w", err)
		}
	}
	for _, f := range files {
		fmt.Printf("uploading %s (clobber if exists)...\n", filepath.Base(f))
		if err := runGH("release", "upload", tag, f, "--clobber"); err != nil {
			return fmt.Errorf("gh release upload %s: %w", filepath.Base(f), err)
		}
	}
	fmt.Println("upload complete")
	return nil
}

func runGH(args ...string) error {
	cmd := exec.Command("gh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
`

func FixScriptGithubReleaseAssets(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, scriptGithubReleaseAssetsPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "script/github/release-assets",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", scriptGithubReleaseAssetsPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "script/github/release-assets"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", scriptGithubReleaseAssetsPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(scriptGithubReleaseAssetsStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", scriptGithubReleaseAssetsPath)}
	return result, nil
}
