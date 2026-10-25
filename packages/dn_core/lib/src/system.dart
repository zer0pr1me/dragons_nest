import 'dart:typed_data';

/// Static description of an emulated system (NES, and later others).
final class SystemInfo {
  const SystemInfo({
    required this.id,
    required this.name,
    required this.screenWidth,
    required this.screenHeight,
    required this.framesPerSecond,
    required this.audioSampleRate,
  });

  /// Stable identifier, e.g. `nes`.
  final String id;
  final String name;
  final int screenWidth;
  final int screenHeight;
  final double framesPerSecond;
  final int audioSampleRate;
}

/// A runnable machine. The frontend only ever talks to this contract.
///
/// Provisional: expected to evolve during Phase 3 (see docs/ROADMAP.md).
abstract interface class EmulatedSystem {
  SystemInfo get info;

  /// Loads a game image (e.g. an iNES file). Throws [FormatException] when
  /// the image is not supported.
  void load(Uint8List image);

  void reset({bool hard = false});

  /// Runs until the next complete video frame.
  void runFrame();

  /// Last completed frame as RGBA8888, `screenWidth * screenHeight * 4` bytes.
  Uint8List get frameBuffer;

  /// Drains audio produced since the previous call: mono, 16-bit signed PCM.
  Int16List takeAudio();

  /// Sets the pressed buttons of controller [port] as a system-defined
  /// bitmask (for the NES: A, B, Select, Start, Up, Down, Left, Right).
  void setInput(int port, int buttons);

  Uint8List saveState();

  void loadState(Uint8List state);
}
