# Dragon's Nest: notes for contributors and agents

- Read `docs/ROADMAP.md` first; it is the source of truth for what to do next.
  Tick checkboxes in the same commit that completes an item.
- Respect the layering rules in `docs/ARCHITECTURE.md`: no Flutter, `dart:ui`
  or `dart:io` in `packages/*`; cores stay deterministic.
- Record significant design decisions as a new ADR in `docs/adr/`.
- Before committing, run: `dart format .`, `flutter analyze --fatal-infos`,
  `dart test` in each changed package and `flutter test` in `app/`.
- Never commit ROMs or downloaded test suites (see `docs/TESTING.md`).
- Hot paths (CPU step, PPU dot, bus access): no allocation, use typed data,
  and mask arithmetic explicitly (`& 0xFF`, `& 0xFFFF`).

## Ground rules

- Minimalism and modularity: write the least code that does the job, in small
  focused units. Refactor periodically rather than letting cruft pile up.
- Test everything, thoroughly but not excessively: cover behaviour and edge
  cases, skip redundant or trivial tests.
- Small, easy-to-follow steps: one coherent change per commit, no
  thousand-line drops. Larger commits only in exceptional, justified cases.
