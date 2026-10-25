/// An 8-bit data / 16-bit address bus, as seen by a CPU core.
///
/// Every access corresponds to exactly one CPU cycle; implementations use
/// this to advance the rest of the machine (PPU, APU, timers) in lockstep.
/// See docs/ARCHITECTURE.md, "Timing model".
abstract interface class Bus {
  int read(int address);

  void write(int address, int value);
}
