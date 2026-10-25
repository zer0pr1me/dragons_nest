import 'dart:typed_data';
import 'dart:ui' as ui;

import 'package:dn_core/dn_core.dart';
import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';

/// Screen geometry the real NES core will report.
const nesInfo = SystemInfo(
  id: 'nes',
  name: 'Nintendo Entertainment System',
  screenWidth: 256,
  screenHeight: 240,
  framesPerSecond: 60.0988,
  audioSampleRate: 48000,
);

/// Placeholder screen: draws a moving test pattern through the same
/// RGBA buffer → `ui.Image` → painter path the emulator will use, and
/// shows the frame rate actually achieved.
class TestPatternScreen extends StatefulWidget {
  const TestPatternScreen({super.key});

  @override
  State<TestPatternScreen> createState() => _TestPatternScreenState();
}

class _TestPatternScreenState extends State<TestPatternScreen>
    with SingleTickerProviderStateMixin {
  final _pixels = Uint8List(nesInfo.screenWidth * nesInfo.screenHeight * 4);
  late final Ticker _ticker;
  ui.Image? _image;
  bool _decoding = false;
  int _frame = 0;
  int _framesThisSecond = 0;
  int _fps = 0;
  Duration _secondStart = Duration.zero;

  @override
  void initState() {
    super.initState();
    _ticker = createTicker(_onTick)..start();
  }

  void _onTick(Duration elapsed) {
    if (elapsed - _secondStart >= const Duration(seconds: 1)) {
      _fps = _framesThisSecond;
      _framesThisSecond = 0;
      _secondStart = elapsed;
    }
    if (_decoding) return;
    _decoding = true;
    _renderPattern(_frame++);
    ui.decodeImageFromPixels(
      _pixels,
      nesInfo.screenWidth,
      nesInfo.screenHeight,
      ui.PixelFormat.rgba8888,
      (image) {
        _decoding = false;
        if (!mounted) {
          image.dispose();
          return;
        }
        _framesThisSecond++;
        setState(() {
          _image?.dispose();
          _image = image;
        });
      },
    );
  }

  /// Colour bars scrolling sideways over a checkerboard.
  void _renderPattern(int frame) {
    const bars = [
      0xFFFFFF, 0xFFFF00, 0x00FFFF, 0x00FF00, //
      0xFF00FF, 0xFF0000, 0x0000FF, 0x000000,
    ];
    final w = nesInfo.screenWidth;
    final h = nesInfo.screenHeight;
    var i = 0;
    for (var y = 0; y < h; y++) {
      for (var x = 0; x < w; x++) {
        int rgb;
        if (y < h * 2 ~/ 3) {
          rgb = bars[((x + frame) & 0xFF) >> 5];
        } else {
          final checker = (((x + frame) >> 3) ^ (y >> 3)) & 1;
          rgb = checker == 0 ? 0x202020 : 0xB3261E;
        }
        _pixels[i++] = (rgb >> 16) & 0xFF;
        _pixels[i++] = (rgb >> 8) & 0xFF;
        _pixels[i++] = rgb & 0xFF;
        _pixels[i++] = 0xFF;
      }
    }
  }

  @override
  void dispose() {
    _ticker.dispose();
    _image?.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(title: const Text("Dragon's Nest")),
      body: Column(
        children: [
          Expanded(
            child: Center(
              child: AspectRatio(
                aspectRatio: nesInfo.screenWidth / nesInfo.screenHeight,
                child: CustomPaint(painter: _FramePainter(_image)),
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(16),
            child: Text(
              '${nesInfo.screenWidth}×${nesInfo.screenHeight} test pattern'
              ' · $_fps fps · no emulator core yet',
              style: theme.textTheme.bodyMedium?.copyWith(
                color: Colors.white70,
              ),
              textAlign: TextAlign.center,
            ),
          ),
        ],
      ),
    );
  }
}

class _FramePainter extends CustomPainter {
  _FramePainter(this.image);

  final ui.Image? image;

  @override
  void paint(Canvas canvas, Size size) {
    final image = this.image;
    if (image == null) return;
    canvas.drawImageRect(
      image,
      Rect.fromLTWH(0, 0, image.width.toDouble(), image.height.toDouble()),
      Offset.zero & size,
      Paint()..filterQuality = FilterQuality.none,
    );
  }

  @override
  bool shouldRepaint(_FramePainter oldDelegate) => oldDelegate.image != image;
}
