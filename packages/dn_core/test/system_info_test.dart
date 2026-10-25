import 'package:dn_core/dn_core.dart';
import 'package:test/test.dart';

void main() {
  test('SystemInfo holds its values', () {
    const info = SystemInfo(
      id: 'test',
      name: 'Test',
      screenWidth: 256,
      screenHeight: 240,
      framesPerSecond: 60,
      audioSampleRate: 48000,
    );
    expect(info.screenWidth * info.screenHeight, 61440);
  });
}
