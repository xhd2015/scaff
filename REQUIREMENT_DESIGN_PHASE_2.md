# REQUIREMENT_DESIGN_PHASE_2 — plan phase P2

## Plan phase

**P2 — Wire slash IDs through fix/lint**  
**Mode: Classic TDD** for CLI contracts; Catalog already slash (P1 GREEN).

## Goal

- `scaff fix <slash-id>` works for all catalog fix rules.
- `scaff fix <old.dotted.id>` is **unknown rule** (strict, non-zero exit).
- FixResult / lint RuleID strings use slash form (no dotted RuleID in fix path).

## Exit criteria

| Check | Expected |
|-------|----------|
| `scaff fix git/ignore --dry-run` | succeeds (or dry-run message) |
| `scaff fix git.ignore` | error unknown rule, exit ≠ 0 |
| Same for at least one other rule pair | e.g. `github/release` vs `github.release` |
| Fix implementations emit slash RuleID | package-level or CLI output |

## Out of scope

- Full doctest rewrite of all fix/ leaves under scaff-cli (P3) — but **this phase’s tree** may invoke scaff CLI.
- New release-assets rule.
- Behavior sketches.

## Suggested leaves (extend tests/scaff-rule-ids or new tree)

| Leaf | Expected |
|------|----------|
| `fix/slash-id-accepted` | fix git/ignore --dry-run exit 0 |
| `fix/dotted-id-rejected` | fix git.ignore fails unknown |
| `fix/slash-github-release` | fix github/release --dry-run ok |
| `fix/dotted-github-release-rejected` | fix github.release unknown |

Use real scaff binary build from ModuleRoot.

## Designer

`doctest skill designer --show` first. Classic TDD. Workspace: /Users/xhd2015/Projects/xhd2015/scaff
No production code. Keep P1 leaves intact.
