# Testing strategy

Emulator correctness is measured against hardware-verified test suites, not
hand-written expectations alone.

## Levels

1. **Unit tests**: per-component behaviour (flag logic, iNES parsing, mapper
   banking, PPU register side effects). Plain `package:test`.
2. **CPU single-step tests**: [SingleStepTests/65x02](https://github.com/SingleStepTests/65x02)
   (formerly Tom Harte's ProcessorTests): about 10,000 JSON cases per opcode with
   initial state, final state and the exact bus activity per cycle. These are
   the main oracle for Phase 1.
3. **Test ROMs**: run headlessly; results are read from the ROM's status
   protocol (blargg ROMs write a status byte at `$6000` and text at `$6004`), or
   compared to a reference trace (`nestest.log`).
4. **Golden frames**: run a game ROM for N frames with scripted input and compare
   a framebuffer hash. These catch regressions in the PPU and mappers.

## Test assets are not committed

Test ROMs and JSON suites are large and come under assorted licences, so they
live outside git:

- `test-roms/` (git-ignored) holds downloaded suites.
- A fetch script (added in Phase 1, `tool/fetch_test_assets.dart`) downloads
  pinned versions and checks their SHA-256.
- Tests that need assets are tagged (`@Tags(['roms'])`) and **skip with a
  clear message** when the assets are missing, so `dart test` always works on a
  fresh clone. CI fetches the assets and runs everything.
- Commercial game ROMs are never downloaded or committed. Golden tests on them
  run only locally, from a path in the `DN_GAME_ROMS` environment variable.

## Suites by phase

| Phase | Suite | Checks |
|---|---|---|
| 1 | SingleStepTests 6502 | every opcode, cycle-level bus activity |
| 1 | Klaus Dormann `6502_functional_test` | whole-CPU sanity, decimal mode |
| 2 | `nestest.nes` + `nestest.log` | CPU in NES context, cycle counts |
| 2 | blargg `instr_test-v5`, `instr_timing`, `cpu_interrupts_v2` | CPU edge cases |
| 2 | blargg `ppu_vbl_nmi`, `sprite_hit_tests`, `sprite_overflow_tests` | PPU timing |
| 2 | `mmc3_test_2` | MMC3 IRQ |
| 3 | blargg `apu_test`, `dmc_dma_during_read4` | APU, DMA |

The [nes-test-roms](https://github.com/christopherpow/nes-test-roms) collection
mirrors most of these.

## Running

```sh
dart run tool/fetch_test_assets.dart   # once (from Phase 1)
cd packages/dn_mos6502 && dart test
dart test --exclude-tags roms          # fast loop
```
