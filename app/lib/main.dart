import 'package:flutter/material.dart';

import 'src/test_pattern_screen.dart';

void main() {
  runApp(const DragonsNestApp());
}

class DragonsNestApp extends StatelessWidget {
  const DragonsNestApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Dragon's Nest",
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
          seedColor: const Color(0xFFB3261E),
          brightness: Brightness.dark,
        ),
      ),
      home: const TestPatternScreen(),
    );
  }
}
