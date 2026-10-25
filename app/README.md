# Dragon's Nest app

Flutter frontend for Dragon's Nest. For now it shows an animated 256×240 test
pattern drawn through the same RGBA buffer → `ui.Image` path the emulator will
use, plus the frame rate it reaches.

```sh
flutter run -d linux      # or macos, windows, chrome, an Android/iOS device
flutter build web --release --no-web-resources-cdn
```

See the [roadmap](../docs/ROADMAP.md) for what comes next.
