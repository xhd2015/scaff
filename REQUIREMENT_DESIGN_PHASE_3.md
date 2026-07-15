# REQUIREMENT_DESIGN_PHASE_3 — plan phase P3

## Plan phase

**P3 — Doctests + skill docs use slash IDs only**  
**Mode: Classic TDD relative to docs/tests still on dots** — designer **rewrites** sealed expectations to slash form. After rewrite, many leaves may go GREEN immediately (P2 already fixed CLI). Remaining RED = skill/docs asserts still mentioning old IDs until implementer.

## Goal

- All `tests/scaff-cli/**` use slash rule IDs in SETUP/ASSERT/DOCTEST prose and code.
- Skill TOPIC docs and SKILL.md examples use slash IDs for rules.
- `rg` for old dotted rule IDs as **command/rule identifiers** in tests/docs is clean.

## Mapping (same as P1)

git/ignore, github/testing-workflow, script/generate, script/install, script/build,
script/bundle/for-linux, git/hooks, git/hooks/install, github/release, install/via-curl

## Designer work

1. Mass-update `tests/scaff-cli/` SETUP.md ASSERT.md DOCTEST.md: replace dotted rule IDs with slash.
2. Update skill show ASSERTS under tests/scaff-cli/skill that expect old IDs in docs.
3. Prefer editing existing tree (not new feature tree). Version stays 0.0.2.

## Out of scope

- release-assets new rule
- behavior sketches
- browser-agent

## Designer

`doctest skill designer --show` first. Workspace: /Users/xhd2015/Projects/xhd2015/scaff
You MAY edit test markdown under tests/scaff-cli (this phase is the test migration).
Do not edit production go except if designer is blocked—prefer leave to implementer for docs under docs/.

Actually: skill docs live in docs/ — designer can update docs if needed for skill show leaves GREEN, OR leave docs to implementer. Prefer designer updates tests; implementer updates docs/**.
