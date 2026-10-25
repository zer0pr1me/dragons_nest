/// System-agnostic emulation contracts shared by every Dragon's Nest core.
///
/// Nothing in this package depends on Flutter, `dart:io` or `dart:ui`, so it
/// runs on the VM, in isolates, in tests and (later) on the web.
library;

export 'src/bus.dart';
export 'src/system.dart';
