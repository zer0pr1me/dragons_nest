import 'dart:typed_data';

import 'package:dn_core/dn_core.dart';

/// Flat 64 KiB RAM that logs every access as `(address, value, kind)`,
/// matching the per-cycle format of the SingleStepTests suite.
final class RamBus implements Bus {
  final Uint8List ram = Uint8List(0x10000);
  final List<(int, int, String)> log = [];

  @override
  int read(int address) {
    final value = ram[address];
    log.add((address, value, 'read'));
    return value;
  }

  @override
  void write(int address, int value) {
    ram[address] = value;
    log.add((address, value, 'write'));
  }

  /// Copies [bytes] into RAM starting at [address].
  void load(int address, List<int> bytes) =>
      ram.setRange(address, address + bytes.length, bytes);
}
