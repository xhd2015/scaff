# Review: scaff vs go-best-practice

**Date:** 2026-08-06  
**Branch:** `master-2026-08-06-use-go-best-practice-to-review-current-project`  
**Module:** `github.com/xhd2015/scaff`  
**Scope:** codebase structure, CLI design, flag handling, package layout  
**Method:** Inspect `cmd/`, `internal/`, `docs/`, generated rule stubs; map findings to go-best-practice topics (`cli/*`, `flags-parsing/*`, `cmd-exec`, `go-embed-assets`, `kool-create`).  
**Status:** Review only — no functional fixes applied.

---

## 1. Executive summary

**scaff** is a solid *amend-not-create* scaffolding CLI: clear subcommands (`lint` / `fix` / `rules` / `skill`), slash-form rule IDs, less-flags on the host binary, Shape-3 skill packaging via `skillcmd.SingleSkill` + `//go:embed`, and a consistent dry-run gate on fix paths.

The largest gap is **what scaff teaches through the code it generates**. Host tooling largely follows go-best-practice; several **scaffold stubs** diverge from the same recipes (mixed flag libraries, raw `os/exec`, missing per-level `--help`). Second-order debt is **rule registration in five places**, which will drift as rules grow.

| Area | Verdict |
|------|---------|
| Package layout (`cmd` / `internal` / `docs`) | Good |
| Host CLI dispatch + help | Good (matches “no toplevel flags”) |
| Host flags (less-flags) | Good, with small UX gaps |
| Host dry-run | Good pattern; prefix style differs from recipe |
| Skill CLI (Shape 3) | Strong |
| Generated script CLIs | **Inconsistent** with flags / cmd-exec / help recipes |
| Rule registry / package growth | Needs consolidation |
| Color / JSON polish | Optional upgrades |
| kool-create / fat go:embed hydrate | N/A for this product shape |

---

## 2. Project snapshot

```text
cmd/scaff/main.go          # binary entry: lint | fix | rules | skill
docs/                      # Shape-3 skill tree + embed.go (SKILL.md + TOPIC.md)
internal/
  audit/                   # project detect + lint orchestration
  fix/                     # fix dispatcher
  model/                   # shared types
  output/                  # human + JSON printers
  rules/                   # catalog, registry, per-rule lint/fix + stubs
script/generate/           # no-op generator entry (dogfood of script/generate)
tests/                     # doctest trees for CLI + rule IDs
```

**Dependencies (direct):** `less-flags`, `xgo` (and `skills` for `skillcmd`).

**Domain model:** profiles `go` | `node` | `polyglot` | `generic`; lint subset of catalog; fix is one-rule-at-a-time with `--dry-run`.

---

## 3. What already aligns well

Grounded in go-best-practice:

| Topic | Evidence |
|-------|----------|
| **flags-parsing / subcommand — no toplevel flags** | Root dispatches on `os.Args[1]`; each command owns less-flags + `Help("-h,--help", …)`. Matches the “no toplevel flags” recipe. |
| **flags-parsing — less-flags on host** | `lint` / `fix` / `rules` use `github.com/xhd2015/less-flags` with typed targets and remain-arg rejection. |
| **cli/skill-cli Shape 3** | `docs.SkillMD` + `docs.TreeFS` + `skillcmd.SingleSkill`; actions as flags (`--show` / `--install` / `--list`); both path orders documented and tested; no show/install subcommands. |
| **cli/dry-run (host fix)** | Single `dryRun bool` through `fix.Apply` → each `Fix*`; discover/plan first, gate writes — not a separate dry-run codepath. |
| **cli/dry-run (github/release stub)** | Template is textbook: soft-fail enrichment on dry-run, same naming formula, `[dry-run]` lines, less-flags. |
| **cmd-exec (some stubs)** | `script/build`, `script/install`, `script/dev` use `cmd.Debug().Run(...)`. |
| **go:embed (skill docs)** | Correct Shape-3 layout: root SKILL.md separate from nested topic dirs; compile-safe markdown tree (no fat UI assets). |
| **kool-create** | Product correctly scopes to *amend*; greenfield stays outside scaff (docs/overview). |
| **Exit codes** | Parse/usage-ish failures → `2`; runtime/fix failures or lint issues → `1`; success → `0`. |
| **Idempotent fix** | Existing targets → “nothing to do” without overwrite — good amend semantics. |

---

## 4. Findings (severity order)

Severity guide:

- **High** — wrong library/pattern in user-facing output of scaff, or structural drift that will cause real bugs as the catalog grows.
- **Medium** — UX/consistency gaps or missing recipe compliance that users will feel or that complicates maintenance.
- **Low** — polish, dogfooding, or optional alignment.

### H1. Generated CLIs use three flag stacks (High)

**Topic:** `flags-parsing`, `flags-parsing/types`

| Stub | Flag stack |
|------|------------|
| Host `cmd/scaff` | `github.com/xhd2015/less-flags` |
| `github/release` template | `less-flags` |
| `script/bundle/for-linux` | `github.com/xhd2015/less-gen/flags` |
| `script/github/release-assets` | stdlib `flag` |
| `script/build`, `install`, `dev` | no flags (ok) |
| `git/hooks`, `git/pre-commit`, `project/layout/cmd` | manual / none |

As a scaffolding tool, scaff *defines* house style for other repos. Emitting `less-gen/flags` and stdlib `flag` fights the less-flags recipe and the host binary’s own stack.

**Recommend:** Standardize all scaffolded Go CLIs on `github.com/xhd2015/less-flags` (same API style as host + `github/release` stub). Drop `less-gen/flags` from templates; rewrite release-assets sketch to less-flags.

---

### H2. Generated external commands inconsistently ignore `cmd-exec` (High)

**Topic:** `cmd-exec`

| Stub | Runner |
|------|--------|
| `script/build`, `script/install`, `script/dev` | `xgo/support/cmd` (`Debug().Run`) |
| `script/bundle/for-linux` | raw `os/exec.Command` + manual `Env`/`Stdout`/`Stderr` |
| `script/github/release-assets` | raw `os/exec` for `git` and `gh` |
| `git/pre-commit` | raw `os/exec` for `git rev-parse` / `git add` |

Recipe preference: fluent `cmd.Debug()` (or `cmd.Output` for capture) with inherit I/O, `.Dir`, `.Env`.

**Recommend:** Prefer `cmd.Debug().Env(...).Run(...)` for fire-and-forget; `cmd.Output(...)` for captured git metadata. Keep raw `os/exec` only where the recipe does not cover a need (rare).

Example direction for bundle:

```go
return cmd.Debug().
    Env([]string{"GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0"}).
    Run("go", "build", "-o", output, ".")
```

---

### H3. Rule surface registered in five places (High)

**Topic:** package layout / maintainability (feeds CLI correctness)

To add a fix rule today you must touch:

1. `internal/rules/catalog.go` (`Catalog`)
2. `internal/rules/registry.go` (`AllFixRules`, and `DefaultLintRules` if linted)
3. `internal/fix/fix.go` (`Apply` switch + relies on `IsKnownRule` → `AllFixRules`)
4. `internal/audit/lint.go` (lint switch, if lint-enabled)
5. Docs topic + doctests (expected)

`IsKnownRule` does **not** read `Catalog`; it reads `AllFixRules`. Catalog tests assert ID set equality, but nothing forces `AllFixRules` ≡ `FixRules()` or that every catalog Fix entry has an `Apply` case.

**Recommend:** Single source of truth:

- Keep `Catalog` as the inventory.
- Derive fix ID list and lint ID list from `Catalog` (`Lint` / `Fix` bits).
- Register handlers once, e.g. `map[string]func(Project, bool) (FixResult, error)` and `map[string]func(Project) RuleResult`, or a small `Rule` interface with optional `Lint`/`Fix` methods.
- Delete duplicated `AllFixRules` / `DefaultLintRules` slices (or generate them in `init` from Catalog + handler presence checks in tests).

---

### M1. Many stubs omit per-level `--help` (Medium)

**Topic:** `flags-parsing/subcommand`, `cli/skill-cli` (“every command level needs `--help`”)

Examples:

- `git/hooks` stub: usage only when `len(os.Args) < 2`; no `-h`/`--help` handling on subcommands.
- `git/pre-commit` stub: no help.
- `project/layout/cmd` stub: `run` always returns `"not implemented"`; no help surface.
- Host is fine: root, lint, fix, rules, skill (via skillcmd).

**Recommend:** Every scaffolded `main` that accepts args should wire `Help("-h,--help", help)` (or skillcmd-equivalent). For multi-command stubs (`git-hooks`), handle help at root **and** document subcommands.

---

### M2. Dry-run output prefix differs from recipe (and from scaff’s own release stub) (Medium)

**Topic:** `cli/dry-run`

| Source | Style |
|--------|--------|
| Recipe + `github/release` template | `[dry-run] would …` on stdout; `[dry-run] warning:` on stderr |
| Host `scaff fix --dry-run` | `dry-run: would create …` / `dry-run: would append …` |

Behavior (single path + side-effect gate) is correct; **prefix/convention** is not.

**Recommend:** Align host fix messages with `[dry-run]` so dry-run is greppable and consistent with scaff-generated release scripts and the dry-run recipe.

---

### M3. `scaff fix` peels rule ID before flags (Medium)

**Topic:** `flags-parsing` UX / flexible flag order

```go
ruleID := args[0]
args = args[1:]
// then lessflags.Parse(args)
```

Works: `scaff fix git/ignore --dry-run`  
Fails as unknown rule: `scaff fix --dry-run git/ignore` (first token is the flag).

Skill surface intentionally supports both flag/path orders; fix does not.

**Recommend (pick one and document):**

1. **Preferred:** Parse flags with less-flags first (or `StopOnFirstArg` after optional global-ish flags), then take remaining positional as rule ID — both orders work; or  
2. Keep rule-first, but improve the error when `args[0]` looks like a flag (`starts with -`) with a hint: `usage: scaff fix <rule> [options]`.

---

### M4. `--profile` is not validated (Medium)

**Topic:** `flags-parsing/types`, CLI UX

```go
profile = model.Profile(profileOverride) // any string accepted
```

Typos (`--profile goo`) become a non-matching profile: gitignore patterns fall through to **universal-only** (same as generic), and `tests/doctest` may still run for non-node/generic switch fall-through (polyglot-like behavior for doctest when profile is unknown — see `LintTestsDoctest` switch). Surprising and hard to debug.

**Recommend:** After parse, require one of `go|node|polyglot|generic` or empty (auto); exit `2` with enumerated values on mismatch. Optional: `**string` / unset vs set if you need tri-state later.

---

### M5. Lint human summary ignores real project path (Medium)

**Topic:** `cli` UX / honesty of output

```go
fmt.Fprintf(w, "scaff lint: %d issue(s) in .\n\n", ...)
```

When users pass `--dir /path/to/svc`, the report still says `.`.

**Recommend:** Print `report.Project.Root` (or a relative path when under cwd).

---

### M6. Machine JSON lacks conventional `json` tags (Medium)

**Topic:** `cli/streaming` (machine-readable stability)

`model.Project`, `RuleResult`, `LintReport` and `rules.RuleInfo` rely on default Go field names (`Root`, `ID`, `Lint`, …). Works and is tested, but:

- Not conventional for public machine APIs (`root` / `id` snake or lower-camel).
- `rules --json` exposes unexported-style keys `Lint`/`Fix` bools with Go names.

**Recommend:** If JSON is a supported contract, add explicit `json` tags and freeze them in doctests. Keep full JSON documents (small reports) — buffering a single document is justified per streaming recipe.

---

### M7. Large stub bodies live as raw Go string constants (Medium)

**Topic:** package layout; light use of `go:embed` (not full `go-embed-assets`)

`internal/rules/*.go` embeds multi-hundred-line templates as backtick strings (escaping hell for nested backticks in bundle/release-assets). Editing templates is error-prone; diffs are noisy.

**Recommend:** Move stubs to e.g. `internal/rules/templates/<rule>/...` and `//go:embed templates/...`. This is **not** the fat-UI hydrate stack from `go-embed-assets` (placeholders + release tarballs) — just compile-time embed of source sketches. Keep placeholders only if you later ship generated binary assets.

---

### M8. Skill topic path ≠ rule ID for some rules (Medium)

**Topic:** `cli/skill-cli` discovery UX

Documented intentionally in `docs/fix/TOPIC.md`:

| Rule ID | Topic path |
|---------|------------|
| `script/bundle/for-linux` | `script/bundle-for-linux` |
| `install/via-curl` | `install-via-curl` |

Users who mirror slash rule IDs in `skill --show` may miss topics. Skill tree layout (filesystem paths) forces some renames.

**Recommend:** Prefer 1:1 where FS allows; where not, keep the table prominent and add `skill --show` aliases or “see also” at the slash path if skillcmd ever supports redirects. At minimum, ensure `scaff fix <id> --help` (or fix help) points at the topic path.

---

### L1. No ANSI color policy on human output (Low)

**Topic:** `cli/color`

Lint/fix human mode is monochrome. For a status CLI, green OK / red issues / gray meta is a natural fit with `--color` / `--no-color` / `NO_COLOR` auto.

**Recommend:** Optional follow-up; never colorize `--json` output.

---

### L2. Host CLI logic concentrated in one `main.go` (Low)

**Topic:** package layout

`cmd/scaff/main.go` (~264 lines) holds dispatch, resolve, and all help strings. Acceptable size today; will grow with flags.

**Recommend:** When it next grows, split to `cmd/scaff` packages or `internal/cli` (`lint.go`, `fix.go`, `rules.go`, `usage.go`) while keeping a thin `main`.

---

### L3. Scaffolded `project/layout/cmd` does not model best-practice CLI (Low)

**Topic:** `flags-parsing/subcommand`, `kool-create` spirit (good defaults)

Generated `cmd/<name>/main.go` is a stub that always errors `"not implemented"` with no flags/help. Fine as a minimal entry, but scaff could ship a less-flags skeleton (top help + `StopOnFirstArg` placeholder) so new CLIs start on the recipe path.

---

### L4. Dogfooding gaps on scaff itself (Low)

**Topic:** product rules, not a go-best-practice recipe

This repo has no root `README.md` / `LICENSE` / full script tree that its own default lint set expects. `.gitignore` is good for a Go profile. Not a CLI-design defect; note for self-host completeness (`scaff lint` on self would fail several default rules).

---

### L5. `kool-create` / `go-embed-assets` hydrate — intentionally out of scope (Info)

| Topic | Applicability |
|-------|----------------|
| **kool-create** | Greenfield templates; scaff is amend-only. No change required; keep the boundary in overview docs. |
| **go-embed-assets** (placeholder → fat → hydrate) | For UI/extension assets. scaff embeds markdown skill trees only; current `docs/embed.go` is correct. Do not add hydrate layers unless binary assets appear. |

---

## 5. Recommended changes (actionable backlog)

Ordered for impact vs effort. Implementation deferred unless/until requested.

### P0 — Generated code consistency

1. **Unify flag library in all Go stubs to less-flags** (H1)  
   - Rewrite `script/bundle/for-linux` off `less-gen/flags`.  
   - Rewrite `script/github/release-assets` off stdlib `flag`.  
   - Update `tests/scaff-cli/testdata/*` mirrors.

2. **Unify external process execution to `xgo/support/cmd`** (H2)  
   - Bundle: `cmd.Debug().Env(...).Run`.  
   - release-assets: `cmd` for `gh`; `cmd.Output` for `git describe`.  
   - git-pre-commit: `cmd.Output` / `cmd.Debug().Dir(root).Run("git", "add", ...)`.

3. **Add `-h`/`--help` to multi-arg stubs** (M1)  
   - Especially `git-hooks` and any new script with options.

### P1 — Host CLI correctness & registry

4. **Collapse rule registries** (H3)  
   - Catalog-driven lists + handler maps; test: every `Fix: true` has a handler; every `Lint: true` has a linter.

5. **Validate `--profile`** (M4).

6. **Fix dry-run prefix to `[dry-run]`** (M2); update doctest ASSERT strings.

7. **Improve fix argv UX** (M3): flags-before-rule or clearer error.

8. **Lint report path** (M5): print real root.

### P2 — Structure & polish

9. Embed rule templates from files (M7).  
10. JSON tags if machine API is public (M6).  
11. Optional color (L1).  
12. Richer `project/layout/cmd` skeleton (L3).  
13. Align skill topic paths or document aliases louder (M8).

---

## 6. Topic coverage matrix

| go-best-practice topic | Relevance | Assessment |
|------------------------|-----------|------------|
| `flags-parsing` | Core | Host good; stubs inconsistent |
| `flags-parsing/types` | Core | Host uses `*string`/`*bool`; profile validation missing |
| `flags-parsing/subcommand` | Core | Host matches no-toplevel-flags pattern; stubs often lack help |
| `flags-parsing/cut` | Low | No need today (no “run external rest-of-line” host cmd) |
| `flags-parsing/collect` | Low | No parent→child flag forwarding yet |
| `cli/dry-run` | Core | Host path correct; prefix + stub variance |
| `cli/skill-cli` | Core | Shape 3 well done |
| `cli/color` | Optional | Not implemented |
| `cli/streaming` | Low | Small reports; full JSON OK |
| `cli/inline-tui-mouse` | N/A | No TUI |
| `cmd-exec` | Core for stubs | Partial adoption |
| `go-embed-assets` | N/A (hydrate) / light embed OK | Skill markdown embed is correct |
| `kool-create` | Boundary only | Correctly out of product scope |

---

## 7. Package layout notes

**Strengths**

- Standard Go layout: `cmd/scaff` binary, `internal/*` non-importable core, module path clean.
- `docs` as a first-class embed package is the right place for Shape-3 skill content.
- `model` keeps shared DTOs free of I/O.
- `audit` vs `fix` vs `rules` separation is clear: orchestration vs apply vs implementations.

**Pressures**

- `internal/rules` is a **god package** (catalog + 18 rules + huge templates). Acceptable while rules share helpers (`DetectProjectMeta`, `fileExistsAt`); if it keeps growing, split `internal/rules/<area>/` or `internal/fixrules/` with a thin registry package.
- `internal/fix` is a thin switch — either merge into registry or keep as the only dispatcher once maps exist.
- `script/generate` exists as product dogfood; host does not need more top-level cmds.

---

## 8. CLI surface reference (current)

```text
scaff                         # usage (exit 0)
scaff -h | --help | help
scaff lint  [--dir DIR] [--json] [--profile PROFILE]
scaff fix   <rule> [--dir DIR] [--dry-run]
scaff rules [--json]
scaff skill --show|--install|--list …   # skillcmd; both path orders
```

**Notable conventions already worth keeping**

- One rule per `fix` (no fix-all).
- Slash rule IDs only (dotted rejected — tested).
- Skill actions are flags, not subcommands.
- Dry-run is a gate, not a second planner.

---

## 9. Conclusion

Treat **scaff host** as mostly recipe-compliant and **scaff-generated scripts** as the primary remediation surface. Closing H1–H3 (less-flags everywhere in stubs, `cmd` for external processes, single rule registry) would bring the project into strong alignment with go-best-practice and prevent the tool from teaching conflicting patterns.

No code changes were made for this review beyond writing this report.
`)