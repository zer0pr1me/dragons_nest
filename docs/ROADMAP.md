# Roadmap

Each phase ends with a **definition of done** we can check. Phases 1–3 are pure
Dart packages; the Flutter app exists from Phase 0 as a shell and grows into
the real frontend in Phase 4. Tick boxes as work
lands, so this file always shows where the project stands.

## Phase 0: Bootstrap ✅

- [x] Retire the Go prototype (kept for reference in `v1/`)
- [x] Dart pub workspace with `dn_core`, `dn_mos6502`, `dn_nes`
- [x] Lints, formatting and CI (format, analyze, test, web build)
- [x] Minimal Flutter app (`app/`) for all five platforms plus web: draws an
      animated 256×240 test pattern through the RGBA → `ui.Image` path
- [x] Architecture, roadmap, testing docs and initial ADRs

## Phase 1: 6502 CPU (`dn_mos6502`)

Goal: a cycle-accurate NMOS 6502 that passes the standard CPU test suites.

- [ ] Registers, status flags, `Bus`-driven fetch/execute loop
- [ ] All 151 official opcodes with every addressing mode
- [ ] Correct cycle counts, including page-cross and branch penalties
- [ ] Dummy reads and writes (they matter for NES I/O registers)
- [ ] Interrupts: RESET, NMI, IRQ, BRK, with the right polling and hijacking behaviour
- [ ] Decimal mode behind a flag (off for the 2A03, on for a generic 6502)
- [ ] Stable unofficial opcodes (LAX, SAX, DCP, ISC, SLO, RLA, SRE, RRA, ...)
- [ ] Disassembler and trace logger (nestest.log format)

**Done when:** the Tom Harte `6502` SingleStepTests pass for all official and
stable unofficial opcodes (registers, memory and per-cycle bus activity), and
the Klaus Dormann functional test passes.

## Phase 2: NES core, headless (`dn_nes`)

Goal: real games boot and render correctly to an in-memory framebuffer.

- [ ] iNES / NES 2.0 parser, `Cartridge` and `Mapper` abstraction
- [ ] CPU memory map: RAM mirroring, PPU and APU registers, OAM DMA
- [ ] Mapper 0 (NROM)
- [ ] `nestest.nes` automation mode matches `nestest.log`, cycle counts included
- [ ] PPU: registers, VRAM and palette, nametable mirroring, scrolling (loopy v/t/x/w)
- [ ] PPU: dot-accurate background and sprite pipeline, sprite 0 hit, overflow, NMI timing
- [ ] Controllers ($4016/$4017)
- [ ] Headless runner in `tool/` that runs a ROM for N frames and dumps PNGs or hashes
- [ ] Mappers 1 (MMC1), 2 (UxROM), 3 (CNROM), 4 (MMC3 with scanline IRQ), 7 (AxROM)

**Done when:** blargg's `ppu_vbl_nmi`, `sprite_hit_tests` and `instr_test-v5`
pass, and Super Mario Bros., Zelda, Mega Man 2 and Super Mario Bros. 3 render
the right title screens in golden-image tests.

## Phase 3: APU, timing and save states

- [ ] APU: pulse ×2, triangle, noise, DMC, frame counter, IRQs
- [ ] Mixer, band-limited resampling to the host rate (48 kHz), filters
- [ ] DMC DMA cycle stealing
- [ ] Battery-backed PRG-RAM (`.sav`)
- [ ] Versioned binary save states for every component
- [ ] NTSC timing; PAL and Dendy as configuration
- [ ] Finalise the `EmulatedSystem` contract in `dn_core`

**Done when:** blargg's `apu_test` and `dmc_dma_during_read4` pass, and a save
state taken mid-game restores to an identical framebuffer and audio stream.

## Phase 4: Flutter app, desktop first (`app/`)

- [ ] Emulation isolate and message protocol (see ARCHITECTURE.md, "Threading")
- [ ] Video: `ui.Image` per frame, integer or aspect-correct scaling, nearest-neighbour filtering
- [ ] Audio: streaming PCM output, with audio-driven frame pacing (spike first, see ADR-0004)
- [ ] Keyboard input and remappable bindings
- [ ] Open ROM (file picker and drag-and-drop), recent games list
- [ ] Pause, reset, save and load state slots, fast-forward
- [ ] Linux, Windows and macOS builds in CI

**Done when:** SMB plays at a steady 60 fps on desktop with glitch-free audio
and under 3 frames of input latency.

## Phase 5: Mobile

- [ ] Touch controller overlay (customisable layout and opacity)
- [ ] Bluetooth and USB gamepads
- [ ] App lifecycle: auto-pause and auto-save state on background
- [ ] Android and iOS builds; performance pass on a mid-range phone
- [ ] Storage: ROM library folder, scoped storage on Android

**Done when:** the app is playable on a mid-range Android phone and an iPhone
at full speed with the touch overlay and a Bluetooth gamepad.

## Phase 6: Polish and extras (unordered backlog)

- Game library with box art and metadata (from ROM hash databases)
- Shaders and filters (CRT, NTSC composite), using Flutter fragment shaders
- Rewind (ring buffer of save states)
- Cheats (Game Genie)
- Debugger UI: CPU trace, PPU nametable, pattern and OAM viewers
- More mappers, driven by coverage of the commercial library
- Famicom Disk System
- Netplay (rollback, built on deterministic save states)
- Web target (dart2wasm)

## Beyond the NES

The `dn_core` contracts and the reusable CPU package make further systems
possible. Plausible next ones, by how much they reuse:

| System | CPU | Notes |
|---|---|---|
| Atari 2600 | 6507 (reuses `dn_mos6502`) | small, good second system |
| Game Boy / Color | SM83 (new `dn_sm83`) | very popular on mobile |
| Sega Master System | Z80 (new `dn_z80`) | |
| SNES | 65C816 | a big step up in complexity |
