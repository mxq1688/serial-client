# Logcat Serial Debugger（Python 桌面版）

基于 **Python 3 + PyQt6 + pyserial** 的串口调试工具，界面与日志展示参考 Android Logcat。

## 环境要求

- Windows 10/11（本文以 Windows 打包为例）
- Python 3.9+（本机使用 `py -3` 启动）
- 依赖：PyQt6、pyserial（见 `requirements.txt`）

## 安装依赖

```powershell
cd e:\stu\serial-client\python-serial
py -3 -m pip install PyQt6 pyserial
```

> `requirements.txt` 中含 `pyqt6-tools`，与部分环境可能冲突；日常运行只需安装 `PyQt6` 和 `pyserial` 即可。

## 运行（开发模式）

```powershell
py -3 run.py
```

启动后窗口标题为 **Logcat Serial Debugger**。

## 打包成 exe（PyInstaller）

### 方式一：一键脚本（推荐）

双击或在项目目录执行：

```powershell
.\build.bat
```

脚本内部执行的命令与下方「核心打包命令」相同。

### 方式二：手动命令

```powershell
cd e:\stu\serial-client\python-serial

# 1. 安装打包工具（只需一次）
py -3 -m pip install pyinstaller

# 2. 打包（首次约 5～10 分钟）
py -3 -m PyInstaller -w -F --name "LogcatSerialDebugger" --collect-all PyQt6 run.py
```

### 核心打包命令说明

```text
py -3 -m PyInstaller -w -F --name "LogcatSerialDebugger" --collect-all PyQt6 run.py
```

| 参数 | 含义 |
|------|------|
| `py -3 -m PyInstaller` | 用 Python 3 调用 PyInstaller |
| `-w` | 窗口程序，不显示黑色控制台 |
| `-F` | 打成**单个** exe 文件 |
| `--name "LogcatSerialDebugger"` | 输出文件名 |
| `--collect-all PyQt6` | 收集 PyQt6 全部依赖，避免缺 DLL |
| `run.py` | 程序入口 |

### 输出位置

```text
dist\LogcatSerialDebugger.exe
```

完整路径示例：

`e:\stu\serial-client\python-serial\dist\LogcatSerialDebugger.exe`

拷贝该 exe 到其他 Windows 电脑即可运行，**无需安装 Python**。

### 打包常见问题

- **体积**：单文件约 80～150 MB，属 PyQt6 正常范围。
- **耗时**：首次打包较慢，请耐心等待。
- **警告**：终端里关于 Qt SQL / WebEngine 等 `Library not found` 的警告可忽略，不影响串口功能。
- **`python` 找不到**：请使用 `py -3`，不要用 `python`（部分 Windows 未配置 PATH）。

### 打包产生的目录（可删）

| 目录/文件 | 说明 |
|-----------|------|
| `dist/` | **最终 exe**，需要保留 |
| `build/` | 中间文件，可删除后重新打包 |
| `LogcatSerialDebugger.spec` | PyInstaller 配置，再次打包时会复用 |

## 功能

- 自动枚举本机串口（COM / tty）
- 可选波特率，连接/断开
- 实时接收日志，Logcat 风格解析与着色
- 级别过滤、关键词搜索
- 发送命令（自动追加换行）

## 项目文件

| 文件 | 说明 |
|------|------|
| `run.py` | 程序入口 |
| `main_simple.py` | 主窗口 UI |
| `serial_worker.py` | 串口读写线程 |
| `log_parser.py` | Logcat 格式解析 |
| `logcat_colors.py` | 日志颜色 |
| `build.bat` | Windows 一键打包脚本 |

## 使用步骤

1. 用 USB 连接设备，确认系统已识别串口（设备管理器中的 COM 口）
2. 运行 `py -3 run.py` 或 `dist\LogcatSerialDebugger.exe`
3. **Port** 选择串口，**Baud** 选波特率（常见 `115200`）
4. 点击 **Connect**
5. 在底部输入框发送命令，日志区实时显示
