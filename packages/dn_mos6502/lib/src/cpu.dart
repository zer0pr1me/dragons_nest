import 'package:dn_core/dn_core.dart';

/// Bit masks of the status register `P`.
abstract final class Flag {
  static const int c = 0x01; // carry
  static const int z = 0x02; // zero
  static const int i = 0x04; // interrupt disable
  static const int d = 0x08; // decimal
  static const int b = 0x10; // break (only exists on the stack)
  static const int u = 0x20; // unused, reads as 1
  static const int v = 0x40; // overflow
  static const int n = 0x80; // negative
}

/// NMOS 6502 core. Every bus access is one cycle (see ADR-0003).
final class Cpu {
  Cpu(this._bus);

  final Bus _bus;

  int a = 0;
  int x = 0;
  int y = 0;
  int s = 0;
  int p = Flag.u | Flag.i;
  int pc = 0;

  /// Runs the 7-cycle reset sequence: the stack pointer drops by 3 without
  /// writing, I is set and PC is loaded from the vector at `$FFFC`.
  void reset() {
    _bus.read(pc);
    _bus.read(pc);
    for (var n = 0; n < 3; n++) {
      _bus.read(0x100 | s);
      s = (s - 1) & 0xFF;
    }
    p |= Flag.i;
    pc = _bus.read(0xFFFC) | (_bus.read(0xFFFD) << 8);
  }

  /// Executes one instruction.
  void step() {
    final opcode = _fetch();
    switch (opcode) {
      case 0xA9: // LDA #imm
        a = _fetch();
        _setZN(a);
      case 0xEA: // NOP
        _bus.read(pc);
      default:
        throw UnimplementedError(
          'opcode \$${opcode.toRadixString(16).padLeft(2, '0')}',
        );
    }
  }

  int _fetch() {
    final value = _bus.read(pc);
    pc = (pc + 1) & 0xFFFF;
    return value;
  }

  void _setZN(int value) {
    p = (p & ~(Flag.z | Flag.n)) | (value == 0 ? Flag.z : 0) | (value & Flag.n);
  }
}
