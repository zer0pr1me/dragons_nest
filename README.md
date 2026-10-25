# Dragon's Nest

A cross-platform retro game emulator, starting with the **NES / Famicom**.

- **Core** — pure Dart, no Flutter dependency: CPU, PPU, APU, mappers. Testable
  on the plain Dart VM and reusable for future systems.
- **App** — Flutter frontend for Windows, macOS, Linux, Android and iOS (web is
  a stretch goal).

> Status: **Phase 0 — bootstrap.** The previous Go 6502 prototype was retired;
> it is still in git history at commit `49c63f3`. See the
> [roadmap](docs/ROADMAP.md) for what comes next.

## Repository layout

```
packages/
  dn_core/      system-agnostic contracts (Bus, EmulatedSystem, ...)
  dn_mos6502/   MOS 6502 / Ricoh 2A03 CPU
  dn_nes/       NES machine: cartridge, mappers, PPU, APU, input
app/            Flutter frontend (created in Phase 4)
docs/           architecture, roadmap, testing strategy, decision records
```

## Development

Requires Dart SDK ≥ 3.6 (Flutter SDK once `app/` exists).

```sh
dart pub get                 # resolves the whole workspace
dart analyze
dart format .
(cd packages/dn_mos6502 && dart test)
```

## Documentation

- [Roadmap](docs/ROADMAP.md) — phases, milestones, current progress
- [Architecture](docs/ARCHITECTURE.md) — layers, timing model, threading, I/O
- [Testing](docs/TESTING.md) — test ROMs and how accuracy is verified
- [Decision records](docs/adr/) — why things are the way they are

## Legal

Dragon's Nest ships no games or BIOS files. Use ROMs you own.

Licensed under the [MIT License](LICENSE).
