# ADR-0003: CPU-driven, per-access cycle stepping

Status: Accepted (2026-10-25)

## Context

Many NES games and most test ROMs depend on the exact relative timing of the
CPU, the PPU and mapper IRQs. Instruction-level emulation (run a whole
instruction, then catch up the PPU) breaks these games and ROMs in subtle ways.

## Decision

The CPU makes exactly one bus access per cycle, including dummy accesses. The
NES bus steps the PPU 3 dots and the APU 1 cycle on every access, before
servicing it. Interrupts are sampled at the hardware-defined points.

## Consequences

- Cycle accuracy falls out of the structure, and per-cycle CPU tests
  (SingleStepTests) map directly onto the `Bus` interface.
- The PPU is stepped about 5.4 M times per second; its step must be
  allocation-free and branch-light. If needed, the PPU can move to a lazy
  catch-up model behind the same bus without changing the CPU.
