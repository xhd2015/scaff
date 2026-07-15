# REQUIREMENT_DESIGN_PHASE_4 — plan phase P4

## Plan phase

**P4 — Behavior sketch headers on all scaffold Go templates**  
**Mode: Classic TDD**

## Goal

Every scaffolded Go template in `internal/rules/*.go` that writes a `.go` file must include a top-of-file comment with:

```
// usage: ...
//
// Proposed behavior (sketch):
//   1. ...
```

(or equivalent "Proposed behavior" wording)

## Leaves (lean)

Create `tests/scaff-template-sketches/` or extend scaff-cli:

| Leaf | Expected |
|------|----------|
| `templates/script-build-has-sketch` | After fix script/build on empty dir, build.go contains "Proposed behavior" |
| `templates/github-release-has-sketch` | release main.go contains "Proposed behavior" |
| `templates/script-generate-has-sketch` | generate main has sketch |

Cover at least 3 different templates; implementer applies to ALL templates.

## Out of scope

- New release-assets rule (P5)
- browser-agent

## Designer

doctest skill designer --show. Classic TDD. Workspace scaff.
