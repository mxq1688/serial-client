# C# .NET MAUI Serial Debugger

## Requirements

- .NET 7.0 SDK or later
- Visual Studio 2022 or VS Code
- .NET MAUI workload

## Setup

```bash
# Install MAUI workload
dotnet workload install maui

# Restore packages
dotnet restore
```

## Build & Run

```bash
# Run in debug mode
dotnet build
dotnet run

# Build for release
dotnet publish -c Release
```

## Features

- Cross-platform (Windows, macOS, Linux)
- Native performance
- Modern UI with MAUI
- Serial port communication
- Logcat format parsing