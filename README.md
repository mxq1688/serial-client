# 🚀 Serial Debugger Suite - 专业串口调试工具套件

## 📦 项目概览

这是一个功能强大的跨平台串口调试工具套件，包含三个版本：

1. **基础版** - 简单快速的串口模拟工具
2. **专业版** - 支持真实硬件的串口通信
3. **GUI高级版** - 图形界面，数据分析，协议解析

## 🎯 快速开始

### 下载预编译版本

所有可执行文件都在 `go-serial/` 目录下：

```bash
# Windows 用户
./go-serial/serial-debugger-pro-win64.exe COM3 115200

# macOS 用户 (Intel)
./go-serial/serial-debugger-pro-macos-intel /dev/tty.usbserial 115200

# macOS 用户 (M1/M2)
./go-serial/serial-debugger-pro-macos-arm64 /dev/tty.usbserial 115200

# Linux 用户
./go-serial/serial-debugger-pro-linux-x64 /dev/ttyUSB0 115200
```

## 📊 版本功能对比

| 功能特性 | 基础版 | 专业版 | GUI版 |
|---------|-------|--------|-------|
| **界面类型** | 命令行 | 命令行 | 图形界面 |
| **串口通信** | 模拟 | ✅ 真实硬件 | ✅ 真实硬件 |
| **自动端口检测** | ❌ | ✅ | ✅ |
| **数据记录** | ❌ | ❌ | ✅ JSON/CSV |
| **协议解析** | ❌ | ❌ | ✅ Modbus/MQTT |
| **波形显示** | ❌ | ❌ | ✅ 实时图表 |
| **数据回放** | ❌ | ❌ | ✅ |
| **统计分析** | ❌ | ❌ | ✅ |
| **文件大小** | ~2MB | ~2.1MB | ~10MB |

## 🛠️ 技术架构

```
串口调试器套件/
├── go-serial/              # 命令行版本
│   ├── simple-serial.go    # 基础版源码
│   ├── main.go             # 专业版源码
│   └── [8个可执行文件