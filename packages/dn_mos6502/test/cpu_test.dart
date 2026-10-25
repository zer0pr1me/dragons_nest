import 'package:dn_mos6502/dn_mos6502.dart';
import 'package:test/test.dart';

import 'support/ram_bus.dart';

void main() {
  late RamBus bus;
  late Cpu cpu;

  setUp(() {
    bus = RamBus();
    cpu = Cpu(bus);
  });

  /// Places [program] at `$8000` and points PC at it.
  void run(List<int> program) {
    bus.load(0x8000, program);
    cpu.pc = 0x8000;
    cpu.step();
  }

  test('reset loads PC from the vector, drops S by 3 and sets I', () {
    bus.load(0xFFFC, [0x34, 0x12]);
    cpu
      ..s = 0x00
      ..p = 0;
    cpu.reset();

    expect(cpu.pc, 0x1234);
    expect(cpu.s, 0xFD);
    expect(cpu.p & Flag.i, Flag.i);
    expect(bus.log, hasLength(7));
    expect(bus.log.where((e) => e.$3 == 'write'), isEmpty);
  });

  test('NOP takes 2 cycles with a dummy read of the next byte', () {
    run([0xEA, 0x55]);

    expect(cpu.pc, 0x8001);
    expect(bus.log, [(0x8000, 0xEA, 'read'), (0x8001, 0x55, 'read')]);
  });

  group('LDA #imm', () {
    for (final (value, z, n) in [
      (0x42, 0, 0),
      (0x00, Flag.z, 0),
      (0x80, 0, Flag.n),
    ]) {
      test('loads \$${value.toRadixString(16)} and sets Z/N', () {
        cpu.p = Flag.z | Flag.n | Flag.c;
        run([0xA9, value]);

        expect(cpu.a, value);
        expect(cpu.p, z | n | Flag.c);
        expect(cpu.pc, 0x8002);
        expect(bus.log, hasLength(2));
      });
    }
  });

  test('unimplemented opcode throws', () {
    expect(() => run([0xFF]), throwsUnimplementedError);
  });
}
