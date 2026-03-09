# Python PyQt6 Serial Debugger

## Requirements

- Python 3.9+
- PyQt6
- pyserial

## Installation

```bash
pip install -r requirements.txt
```

## Run

```bash
python run.py
```

## Features

- Complete implementation
- PyQt6 modern UI
- Serial port auto-detection
- Real-time log parsing
- Color-coded log levels
- Tag filtering
- Command sending
- Dark theme

## Files

- `main_simple.py` - Main UI window
- `serial_worker.py` - Serial communication thread
- `log_parser.py` - Logcat format parser
- `logcat_colors.py` - Color scheme
- `run.py` - Application entry point

## Usage

1. Connect your device via serial port
2. Select the port from dropdown
3. Set baud rate (usually 115200)
4. Click Connect
5. Logs will appear in real-time