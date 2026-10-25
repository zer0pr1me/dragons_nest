# Dragon's Nest

A cross-platform retro game emulator, starting with the **NES / Famicom**.

- **Core** — pure Dart, no Flutter dependency: CPU, PPU, APU, mappers. Testable
  on the plain Dart VM and reusable for future systems.
- **App** — Flutter frontend for Windows, macOS, Linux, Android and iOS (web is
  a stretch goal).

> Status: **Phase 0 — bootstrap.** The previous Go 6502 prototype is kept for
> reference in [`v1/`](v1/) and is not developed further. See the
> [roadmap](docs/ROADMAP.md) for what comes next.

## Repository layout

```
packages/
  dn_core/      system-agnostic contracts (Bus, EmulatedSystem, ...)
  dn_mos6502/   MOS 6502 / Ricoh 2A03 CPU
  dn_nes/       NES machine: cartridge, mappers, PPU, APU, input
app/            Flutter frontend (currently a test-pattern screen)
docs/           architecture, roadmap, testing strategy, decision records
v1/             the earlier Go 6502 prototype (reference only)
```

## Development

Requires the Flutter SDK (stable, with Dart ≥ 3.6). The app is a workspace
member, so resolving the workspace needs Flutter even for the pure-Dart packages.

```sh
flutter pub get                       # resolves the whole workspace
flutter analyze
dart format .
(cd packages/dn_mos6502 && dart test) # a pure-Dart package
(cd app && flutter run -d linux)      # or macos, windows, chrome, a device
```

## Documentation

- [Roadmap](docs/ROADMAP.md) — phases, milestones, current progress
- [Architecture](docs/ARCHITECTURE.md) — layers, timing model, threading, I/O
- [Testing](docs/TESTING.md) — test ROMs and how accuracy is verified
- [Decision records](docs/adr/) — why things are the way they are

## Legal

Dragon's Nest ships no games or BIOS files. Use ROMs you own.

Licensed under the [MIT License](LICENSE).
