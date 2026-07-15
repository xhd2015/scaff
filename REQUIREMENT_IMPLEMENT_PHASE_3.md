# REQUIREMENT_IMPLEMENT_PHASE_3

## Tests sealed — do not modify tests/

Tests already GREEN (55+7). Implementer only needs **docs/** skill prose migration.

## Work

Update all `docs/**/*.md` and `docs/SKILL.md` so rule IDs use slash form per mapping.
No dotted rule IDs in `scaff fix ...` examples.

## Verify

```sh
doctest test ./tests/scaff-cli ./tests/scaff-rule-ids  # still GREEN
rg 'scaff fix git\.|scaff fix github\.|scaff fix script\.|scaff fix install\.' docs || true
# should have no dotted fix examples
```

## First step: implementer show
