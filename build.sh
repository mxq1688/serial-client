#!/bin/bash
# Multi-platform Serial Debugger Build Script

set -e

echo "========================================"
echo "Serial Debugger Multi-Platform Builder"
echo "========================================"

# Create output directory
mkdir -p builds

# Build Go CLI version
echo "\n[1/7] Building Go CLI version..."
cd go-serial
GOOS=linux GOARCH=amd64 go build -o ../builds/serial-debugger-linux
GOOS=windows GOARCH=amd64 go build -o ../builds/serial-debugger.exe
GOOS=darwin GOARCH=amd64 go build -o ../builds/serial-debugger-mac
cd ..
echo "✓ Go CLI version built"

# Build Go GUI version
echo "\n[2/7] Building Go GUI version..."
cd go-serial-gui
go get fyne.io/fyne/v2/app
go get fyne.io/fyne/v2/widget
go get github.com/tarm/serial
GOOS=linux GOARCH=amd64 go build -o ../builds/serial-gui-linux
GOOS=windows GOARCH=amd64 go build -o ../builds/serial-gui.exe
GOOS=darwin GOARCH=amd64 go build -o ../builds/serial-gui-mac
cd ..
echo "✓ Go GUI version built"

# Build Python version
echo "\n[3/7] Building Python version..."
cd python-serial
pip install pyinstaller pyserial
pyinstaller --onefile --name serial-debugger-py main.py
cp dist/serial-debugger-py ../builds/
cd ..
echo "✓ Python version built"

# Build Rust version
echo "\n[4/7] Building Rust version..."
cd rust-serial
cargo build --release
cp target/release/rust-serial ../builds/serial-debugger-rust
cd ..
echo "✓ Rust version built"

# Build C++ version
echo "\n[5/7] Building C++ version..."
cd cpp-serial
cmake -B build
cmake --build build --config Release
cp build/serial-debugger ../builds/serial-debugger-cpp
cd ..
echo "✓ C++ version built"

# Build C# version
echo "\n[6/7] Building C# version..."
cd csharp-serial
dotnet build -c Release
cp bin/Release/net*/serial-debugger.dll ../builds/
cd ..
echo "✓ C# version built"

# Build Tauri version
echo "\n[7/7] Building Tauri version..."
cd tauri-serial
npm install
npm run tauri build
cp src-tauri/target/release/tauri-serial ../builds/serial-debugger-tauri
cd ..
echo "✓ Tauri version built"

echo "\n========================================"
echo "Build Complete! Check 'builds' directory"
echo "========================================"
ls -la builds/