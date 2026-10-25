# ADR-0002: Pure-Dart cores in a pub workspace

Status: Accepted (2026-10-25)

## Decision

- Emulation code lives in `packages/` as pure Dart packages with no Flutter,
  `dart:ui` or `dart:io` imports.
- Packages are split by reuse boundary: `dn_core` (contracts), `dn_mos6502`
  (CPU, reusable by future 6502 systems), `dn_nes` (the machine).
- The repo is a Dart **pub workspace** (Dart ≥ 3.6): one `pub get` and one
  lockfile, with cross-package references resolved locally.
- The Flutter app (`app/`) joins the workspace in Phase 4. Until then the repo
  builds with just the Dart SDK.

## Consequences

- Fast tests (`dart test`) with no Flutter tooling.
- Cores can later run headless (CLI test runner, benchmarks) and on the web.
- The package prefix `dn_` keeps names short and avoids clashes on pub.dev.
