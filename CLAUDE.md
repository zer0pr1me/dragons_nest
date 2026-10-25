# Dragon's Nest: notes for contributors and agents

- Read `docs/ROADMAP.md` first; it is the source of truth for what to do next.
  Tick checkboxes in the same commit that completes an item.
- Respect the layering rules in `docs/ARCHITECTURE.md`: no Flutter, `dart:ui`
  or `dart:io` in `packages/*`; cores stay deterministic.
- Record significant design decisions as a new ADR in `docs/adr/`.
- Before committing, run: `dart format .`, `dart analyze --fatal-infos`, and
  `dart test` in each changed package.
- Never commit ROMs or downloaded test suites (see `docs/TESTING.md`).
- Hot paths (CPU step, PPU dot, bus access): no allocation, use typed data,
  and mask arithmetic explicitly (`& 0xFF`, `& 0xFFFF`).
