# ADR-0001: Flutter + Dart for a cross-platform app

Status: Accepted (2026-10-25)

## Context

Dragon's Nest should run on desktop (Windows, macOS, Linux) and mobile (Android,
iOS) from one codebase. The earlier prototype was a Go 6502 core with no
frontend.

## Decision

Write everything in Dart. The frontend uses Flutter.

## Why it is feasible

- **CPU budget**: NES emulation needs about 1.8 M CPU cycles and 5.4 M PPU dots per
  second. Dart AOT-compiles to native code on all five targets, and comparable
  managed-language emulators (Java, C#, JS) run the NES at full speed with
  room to spare.
- **Video**: 256×240 RGBA per frame; `decodeImageFromPixels` and a
  `CustomPainter` handle that cheaply. Fragment shaders cover CRT filters.
- **Concurrency**: isolates let emulation run off the UI thread, and
  `TransferableTypedData` moves frames without copying.
- **Platform reach**: one UI codebase; gamepads, file pickers and audio are
  available as plugins.

## Risks

- **Low-latency streaming audio** is the weakest area of the Flutter ecosystem.
  Mitigation: a spike in Phase 4 (ADR-0004), with an FFI binding to miniaudio
  as the fallback.
- **Debug-mode speed** (JIT) will be lower; performance is judged in profile or
  release builds only.
- **iOS distribution**: emulators are allowed on the App Store since 2024,
  provided no copyrighted ROMs are bundled. Sideloading remains an option.

## Consequences

- One language end to end; cores are testable on the plain Dart VM.
- Native code appears only behind plugins or FFI, if at all.
