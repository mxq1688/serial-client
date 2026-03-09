# Rust egui Serial Debugger

## Features

- Native performance with Rust
- Immediate mode GUI with egui
- Serial port communication
- Logcat format parsing
- Dark theme interface

## Build & Run

### Debug mode:
```bash
cargo run
```

### Release mode (optimized):
```bash
cargo build --release
./target/release/logcat-serial
```

## Dependencies

- `egui` & `eframe`: GUI framework
- `serialport`: Serial communication
- `tokio`: Async runtime
- `chrono`: Date/time handling

## Requirements

- Rust 1.70+
- Linux: `libudev-dev` package
- Windows: No additional requirements
- macOS: No additional requirements

## Cross-compilation

### For Windows from Linux:
```bash
cargo build --target x86_64-pc-windows-gnu
```

### For Linux from Windows:
```bash
cargo build --target x86_64-unknown-linux-gnu
```