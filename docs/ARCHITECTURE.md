# Architecture

## Layers

```
┌──────────────────────────────────────────────────────────┐
│ app/  (Flutter)                                          │
│   UI, library, settings, video/audio/input adapters      │
└──────────────▲───────────────────────────────────────────┘
               │ messages (isolate boundary)
┌──────────────┴───────────────────────────────────────────┐
│ emulation isolate: frame loop around an EmulatedSystem   │
└──────────────▲───────────────────────────────────────────┘
               │ dn_core contracts
┌──────────────┴──────────────┐  ┌────────────────────────┐
│ dn_nes  (pure Dart)         │  │ future systems ...     │
│  Cartridge+Mappers PPU APU  │  │                        │
└──────────────▲──────────────┘  └────────────────────────┘
               │ Bus
┌──────────────┴──────────────┐
│ dn_mos6502  (pure Dart)     │
└─────────────────────────────┘
```

Rules:

1. `packages/*` never import Flutter, `dart:ui` or `dart:io`. File access, audio
   and timing belong to the host, so the cores run on the VM, in tests, in
   isolates and on the web.
2. Dependencies point downwards only. A CPU knows nothing about the NES; it sees
   only a `Bus`.
3. Cores are **deterministic**: same ROM, same inputs per frame, same output.
   Save states, rewind, netplay and golden tests all rely on this. No wall
   clock time and no unseeded randomness inside a core.

## Timing model

The NES runs the PPU at 3 dots per CPU cycle (NTSC) and the APU at 1 tick per
CPU cycle. We use **CPU-driven, per-access stepping**:

- The CPU performs exactly one bus access (`read` or `write`) per cycle,
  including dummy reads and writes, as the real chip does.
- The NES bus implementation, on every access, first advances the PPU by 3 dots
  and the APU by one cycle, then performs the access.
- Interrupt lines (NMI from the PPU, IRQ from the APU and mappers) are sampled
  by the CPU at the hardware-correct point of each instruction.

This gives cycle accuracy without a scheduler, and it is the approach that
passes the timing-sensitive test ROMs. If profiling later shows it is too slow,
the PPU can switch to "catch-up" (run lazily up to the current cycle when a
register is touched) behind the same interface.

## Performance notes

- A frame is about 29,780 CPU cycles and 89,342 PPU dots, roughly 5.4 M PPU steps
  per second. That is comfortable for Dart AOT (release/profile builds).
  Debug (JIT) builds will be slower; use `--profile` to judge speed.
- Hot paths use `Uint8List` and `Uint16List`, `int` arithmetic masked with
  `& 0xFF` / `& 0xFFFF`, and no allocation per cycle (no closures, no boxed
  objects, no `List<int>` growth).
- Opcode dispatch is a `switch` on the opcode byte, which compiles to a jump
  table. Benchmark this against a table of static tear-offs before settling.
- Keep a headless benchmark (`tool/bench.dart`: frames per second on a fixed
  ROM) and watch it in CI to catch regressions.

## Threading (Flutter app)

Emulation runs in a **dedicated isolate** so UI jank never stalls the game and
heavy emulation never stalls the UI.

```
UI isolate                          Emulation isolate
──────────                          ─────────────────
input events ── SendPort ─────────▶ latch input for next frame
commands (load/reset/state) ──────▶
                                    runFrame()
◀── TransferableTypedData ───────── RGBA frame (245,760 bytes)
◀────────────────────────────────── PCM samples → audio sink
```

- Frames travel as `TransferableTypedData` (no copy). The UI keeps only the
  latest frame and turns it into a `ui.Image` with `decodeImageFromPixels`,
  drawn by a `CustomPainter` with `FilterQuality.none`.
- **Pacing**: the audio device is the master clock. The emulator runs ahead
  enough to keep the audio buffer at its target fill, and frames are shown on
  the next vsync. With 60.0988 Hz NES against a 60 Hz display, the occasional
  frame gets dropped or repeated, which is acceptable; dynamic rate control
  (resampling the audio by ±0.5%) is a later refinement.
- If the audio plugin cannot be called from a background isolate, audio goes
  through the UI isolate. ADR-0004 settles this.

## Host I/O adapters (app)

| Concern | Plan |
|---|---|
| Video | `ui.Image` per frame, `CustomPainter`; fragment shaders for filters later |
| Audio | streaming 16-bit PCM; library chosen by the Phase 4 spike (ADR-0004) |
| Input | `HardwareKeyboard`, a gamepad plugin, a touch overlay widget; all map to a per-port button bitmask |
| Files | `file_picker` / drag-and-drop; ROM bytes handed to the core as `Uint8List` |
| Storage | `path_provider` app directories for `.sav`, states, settings |

## NES package structure (planned)

```
dn_nes/lib/src/
  nes.dart            NesSystem implements EmulatedSystem; owns the components
  bus.dart            CPU memory map; steps PPU/APU per access
  cartridge/
    ines.dart         iNES / NES 2.0 header parsing
    mapper.dart       abstract Mapper (CPU/PPU address hooks, IRQ, scanline clock)
    mappers/          mapper_000.dart, mapper_001.dart, ...
  ppu/                registers, rendering pipeline, palette
  apu/                channels, frame counter, mixer
  input/              standard controller
  state/              save state serialization
```

## Save states

Every component implements `void save(StateWriter w)` / `void load(StateReader r)`
over a small binary writer with a format version header. States are
versioned, so old states either load or fail cleanly, and never silently
corrupt.
