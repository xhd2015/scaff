# REQUIREMENT_IMPLEMENT_PHASE_1 — plan phase P1

## Tests sealed — do not modify

`./tests/scaff-rule-ids/`

## Target: 3/3 GREEN

```sh
doctest test ./tests/scaff-rule-ids
```

## Implement

Update `internal/rules/catalog.go` Catalog IDs to slash form per mapping:

- git/ignore, github/testing-workflow, script/generate, script/install, script/build
- script/bundle/for-linux, git/hooks, git/hooks/install, github/release, install/via-curl

No dots in any ID. Do **not** add aliases. Do **not** rewrite all fix/*.go RuleID strings unless required for package compile—if Fix still uses old IDs, package may still compile; leave FixResult IDs for P2 if separate.

If something breaks compile because of string compares, minimal fix only.

## First step

`doctest skill implementer --show` or `doctest skill implementer show`
