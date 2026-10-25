# ADR-0004: Audio output backend

Status: Proposed (decided by a spike at the start of Phase 4)

## Context

The core produces mono 16-bit PCM at the host rate. The app needs a
low-latency (≤ 50 ms) streaming sink on all five platforms, and the sink's
consumption rate drives frame pacing.

## Candidates

1. `flutter_soloud`: buffer-stream API, all platforms; check its latency and
   whether it can be fed from a background isolate.
2. Other PCM-stream plugins (e.g. `mp_audio_stream`), if they are maintained
   and cover all five platforms.
3. A thin `dart:ffi` binding to **miniaudio** (single C file, all platforms):
   the most control, at the cost of maintaining native build glue.

## Spike criteria

Steady 48 kHz output for 10 minutes with no underruns on Linux, Windows,
macOS, Android and iOS; latency measured; buffer fill level observable from
Dart (needed for pacing).
