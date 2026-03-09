# Flutter Desktop Serial Debugger

## Setup

1. Install Flutter Desktop support:
```bash
flutter config --enable-linux-desktop
flutter config --enable-windows-desktop
flutter config --enable-macos-desktop
```

2. Install dependencies:
```bash
flutter pub get
```

3. Run the app:
```bash
# Linux
flutter run -d linux

# Windows
flutter run -d windows

# macOS
flutter run -d macos
```

## Features

- Serial port connection
- Logcat format parsing
- Dark theme UI
- Real-time log display
- Log level filtering

## Dependencies

- flutter_libserialport: Serial communication
- provider: State management
- google_fonts: Typography

## Build

```bash
flutter build linux
flutter build windows
flutter build macos
```