# Serial Debugger Pro - 跨平台串口调试工具

## 📦 版本说明

提供了两个版本：

### 1. 基础版 (serial-debugger-*)
- 简单的命令行界面
- 模拟串口通信
- 适合快速测试
- 文件大小: ~2MB

### 2. 专业版 (serial-debugger-pro-*)
- 增强的命令行界面
- 支持实际串口通信（需要硬件）
- 自动检测串口
- 支持多种波特率
- 文件大小: ~2MB

## 🖥️ 支持平台

| 平台 | 基础版 | 专业版 |
|------|--------|--------|
| Windows x64 | ✅ serial-debugger-win64.exe | ✅ serial-debugger-pro-win64.exe |
| macOS Intel | ✅ serial-debugger-macos-intel | ✅ serial-debugger-pro-macos-intel |
| macOS M1/M2 | ✅ serial-debugger-macos-arm64 | ✅ serial-debugger-pro-macos-arm64 |
| Linux x64 | ✅ serial-debugger-linux-x64 | ✅ serial-debugger-pro-linux-x64 |

## 🚀 使用方法

### Windows

```powershell
# 基础版
.\serial-debugger-win64.exe COM3

# 专业版
.\serial-debugger-pro-win64.exe COM3 115200
```

### macOS

```bash
# 首次运行需要授权
chmod +x serial-debugger-macos-*

# Intel Mac
./serial-debugger-macos-intel /dev/tty.usbserial
./serial-debugger-pro-macos-intel /dev/tty.usbserial 9600

# Apple Silicon (M1/M2)
./serial-debugger-macos-arm64 /dev/tty.usbserial
./serial-debugger-pro-macos-arm64 /dev/tty.usbserial 115200
```

### Linux

```bash
# 添加执行权限
chmod +x serial-debugger-linux-x64

# 基础版
./serial-debugger-linux-x64 /dev/ttyUSB0

# 专业版
./serial-debugger-pro-linux-x64 /dev/ttyUSB0 115200

# 如果权限不足，可能需要
sudo ./serial-debugger-pro-linux-x64 /dev/ttyUSB0 115200
# 或将用户添加到 dialout 组
sudo usermod -a -G dialout $USER
```

## 📋 功能特性

### 基础版功能
- ✅ 串口参数配置
- ✅ 发送文本数据
- ✅ 接收数据显示
- ✅ 清屏功能
- ✅ 模拟