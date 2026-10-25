import 'package:dragons_nest/main.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('shows the app title', (tester) async {
    await tester.pumpWidget(const DragonsNestApp());
    expect(find.text("Dragon's Nest"), findsOneWidget);
  });
}
