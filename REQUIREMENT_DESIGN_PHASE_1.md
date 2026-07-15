# REQUIREMENT_DESIGN_PHASE_1 — plan phase P1

## Plan phase

**P1 — Canonical slash ID map (scaff)**  
**Mode: Classic TDD** (catalog still uses dotted IDs today).

## Goal

`rules.Catalog` (and any public rule list) uses **only slash rule IDs**. No dotted IDs. No backward-compat aliases.

## Authoritative mapping

| Old | New |
|-----|-----|
| `git.ignore` | `git/ignore` |
| `github.testing.workflow` | `github/testing-workflow` |
| `script.generate` | `script/generate` |
| `script.install` | `script/install` |
| `script.build` | `script/build` |
| `script.bundle.for-linux` | `script/bundle/for-linux` |
| `git.hooks` | `git/hooks` |
| `git.hooks.install` | `git/hooks/install` |
| `github.release` | `github/release` |
| `install.via.curl` | `install/via-curl` |

## Exit criteria (P1)

- Catalog lists exactly the new slash IDs (same set of rules, new ID strings).
- No Catalog entry `ID` contains `.` (dots).
- Optional: `scaff rules` CLI output uses slash IDs if it reads Catalog (if CLI still broken until P2, assert package Catalog only).

## Out of scope (P1)

- Fix/lint dispatch still may use old IDs until P2 (but prefer implementer update Catalog IDs in FixResult if trivial—**do not** change doctests outside this phase's tree yet).
- Doctest mass rewrite of all fix/ leaves (P3).
- New release-assets rule (P5).
- Behavior sketches (P4).

## Suggested tree

```text
tests/scaff-rule-ids/   # or extend tests/scaff-cli if cleaner
```

Lean leaves:

| Leaf | Expected |
|------|----------|
| `catalog/all-slash-ids` | Every Catalog ID matches mapping / is slash form |
| `catalog/no-dot-ids` | No ID contains `.` |
| `catalog/known-set` | Set equals the 10 mapped IDs (no missing/extra) |

## Designer

1. `doctest skill designer --show` first.
2. Classic TDD; RED until catalog changed.
3. Version 0.0.2.
4. Work in `/Users/xhd2015/Projects/xhd2015/scaff`.
5. No production code.

Module: scaff (github.com/xhd2015/scaff or as in go.mod).
