# 各语言打包 Windows exe 方式总结

本文档说明常见语言如何产出 **Windows 可执行文件（.exe）**，并对应本仓库 `serial-client` 里各子项目的做法。

---

## 先分清两件事

| 概念 | 含义 | 典型工具 |
|------|------|----------|
| **编译（Compile）** | 源码转成机器码，运行时**不依赖**原语言解释器 | Go、Rust、C/C++、C#（Release） |
| **打包（Bundle）** | 把**解释器/运行时 + 你的代码 + 依赖**塞进一个安装包或 exe | Python（PyInstaller）、部分 PHP/Node 方案 |

动态语言（Python、PHP、JavaScript）也可以有 `.exe`，但多数是**打包**，不是传统意义上的「编译成纯机器码」。

---

## 总览对比

| 语言/技术 | 是否常做 exe | 本质 | 体积（约） | 本仓库路径 |
|-----------|--------------|------|------------|------------|
| **Go** | ✅ 很常见 | 编译为原生二进制 | 2～10 MB | `go-serial/`、`go-serial-gui/` |
| **Rust** | ✅ 很常见 | 编译为原生二进制 | 3～15 MB | `rust-serial/` |
| **C / C++** | ✅ 标准做法 | 编译 + 链接（常需带 Qt 等 DLL） | 视依赖而定 | `cpp-serial/` |
| **C# / .NET** | ✅ 很常见 | 编译 IL + 自带 .NET 运行时或单文件发布 | 10～80 MB | `csharp-serial/` |
| **Python** | ✅ 常见 | **打包**（内嵌 Python 解释器） | 80～150 MB | `python-serial/` |
| **Kotlin/Java** | ⚠️ 桌面 exe 少见 | JVM 字节码；Android 为 APK/AAB，IDE 为插件 JAR | — | `android_studio_plugins/`（插件，非 exe） |
| **JavaScript (Electron)** | ✅ 常见 | 打包 Chromium + Node + 前端 | 100 MB+ | `tauri-serial/`（含 Electron 相关文件） |
| **Tauri** | ✅ 常见 | 小窗口壳 + Rust 后端编译进 exe | 通常小于 Electron | `tauri-serial/` |
| **Flutter** | ✅ 桌面可构建 | 编译为原生 + 引擎资源 | 较大 | `flutter-serial/` |
| **PHP** | ❌ 非主流 | 一般不上 exe；Web 部署为主 | — | 本仓库无 |

---

## Go（编译型，推荐做 CLI/小工具）

**原理**：直接编译成机器码，单文件即可运行，无需安装 Go。

```powershell
cd go-serial
# 本机 Windows
go build -o serial-debugger.exe .

# 交叉编译（在任意系统打 Windows exe）
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o serial-debugger.exe .
```

本仓库脚本：`build.sh` 中 `[1/7]`、`[2/7]` 段。

| 优点 | 缺点 |
|------|------|
| 体积小、启动快 | 复杂 GUI 需 Fyne/Wails 等额外框架 |
| 无运行时依赖 | |

---

## Python（动态语言 → 打包 exe）

**原理**：**不是**把 Python 变成 C；PyInstaller 将 **Python 解释器 + 字节码 + PyQt6 等** 打进一个 exe，运行时仍是解释执行。

**本仓库推荐命令**（`python-serial/`）：

```powershell
cd python-serial
py -3 -m pip install pyinstaller
py -3 -m PyInstaller -w -F --name "LogcatSerialDebugger" --collect-all PyQt6 run.py
```

或双击 `build.bat`。

| 参数 | 含义 |
|------|------|
| `-w` | 无控制台，仅 GUI |
| `-F` | 单文件 exe |
| `--collect-all PyQt6` | 收齐 Qt 依赖，避免缺 DLL |

**输出**：`python-serial/dist/LogcatSerialDebugger.exe`

| 优点 | 缺点 |
|------|------|
| 开发快、库多 | 体积大（约 80～150 MB） |
| 对方无需装 Python | 首次打包慢 |

其他工具：`cx_Freeze`、`Nuitka`（偏真编译，配置更重）。

---

## Rust（编译型）

**原理**：`cargo build --release` 生成原生 exe。

```powershell
cd rust-serial
cargo build --release
# 输出: target\release\logcat-serial.exe（名称以 Cargo.toml 为准）
```

| 优点 | 缺点 |
|------|------|
| 性能好、内存安全 | 学习曲线陡、编译时间较长 |

---

## C / C++（编译型）

**原理**：源码 → 编译器 → 链接器 → exe；若用 **Qt** 等动态库，分发时可能要带若干 DLL 或用安装包。

```bash
cd cpp-serial
mkdir build && cd build
cmake ..
cmake --build . --config Release
# Windows 下常见输出: build\Release\LogcatSerial.exe
```

| 优点 | 缺点 |
|------|------|
| 性能极好、生态成熟 | 构建链复杂（CMake、Qt、MSVC） |

---

## C# / .NET（编译 + 运行时）

**原理**：编译为 IL，运行依赖 .NET；可用**自包含单文件**发布，看起来像「一个 exe」。

```powershell
cd csharp-serial
dotnet publish -c Release -r win-x64 --self-contained true -p:PublishSingleFile=true
```

输出在 `bin\Release\net*\win-x64\publish\` 下。

| 优点 | 缺点 |
|------|------|
| Windows 桌面开发体验好 | 单文件仍可能较大（含运行时） |

---

## Kotlin / Java

| 形态 | 打包结果 | 说明 |
|------|----------|------|
| **Android 应用** | `.apk` / `.aab` | 不是 exe，给手机/模拟器安装 |
| **Android Studio 插件** | `.zip`（Marketplace） | 如本仓库 `serial-port-plugin`，在 IDE 内加载 |
| **桌面 JVM** | `.jar` + 需安装 JRE，或 `jpackage` 打安装包 | 国内桌面较少用 exe 分发 JAR |

桌面 exe 可用 **`jpackage`**（JDK 14+）生成 Windows 安装程序，不是本串口项目的主路径。

---

## JavaScript 生态

JS 本身在浏览器或 Node 里**解释执行**；所谓「exe」几乎都是 **壳 + 运行时 + 你的 JS/HTML/CSS**。

### 方式对比（桌面 GUI）

| 方案 | 原理 | 体积 | 特点 |
|------|------|------|------|
| **Electron** | Chromium + Node | 很大（~100MB+） | 最成熟，VS Code、Discord 同类 |
| **Tauri** | 系统 WebView + Rust 后端 | 较小 | 本仓库 `tauri-serial/` |
| **NW.js** | 类似 Electron，Chromium + Node | 大 | 老牌，API 与 Electron 接近 |
| **Neutralinojs** | 系统自带 WebView，无捆绑 Chromium | 很小（几 MB 级） | 功能比 Electron 弱，适合轻量工具 |
| **Wails** | Go 后端 + 系统 WebView 前端 | 较小 | 用 Go 写逻辑，界面仍用 HTML/JS |
| **WebView2 + 壳** | 微软 Edge 内核，C#/C++ 包一层 | 中等 | Windows 上常见企业方案 |

### Electron

**原理**：打包 **Chromium + Node.js + 前端资源**。

```bash
npm install
# 常用打包: electron-builder / electron-forge
npm run build
```

### Tauri（本仓库 `tauri-serial/`）

**原理**：界面用 Web 技术，**核心逻辑用 Rust 编译进 exe**。

```bash
cd tauri-serial
npm install
npm run tauri build
# 输出: src-tauri/target/release/bundle/
```

### NW.js

```bash
npm install nw nw-builder -D
# 配置后 nwbuild 生成 exe
```

与 Electron 类似，适合已有 NW 经验或历史项目。

### Neutralinojs

```bash
npm install -g @neutralinojs/neu
neu create myapp && neu build
```

不内嵌 Chromium，依赖系统 WebView，**体积小**，串口等能力需通过扩展/本地服务配合。

### Wails（Go + 前端 JS）

```bash
wails init
wails build
# Windows 下产出 exe，前端仍是 Vue/React 等
```

逻辑在 **Go** 里写，不是纯 Node；若团队会 Go，可替代 Electron。

### 仅 Node、无窗口（CLI 工具）

若只是命令行脚本，不是要 GUI：

| 工具 | 说明 |
|------|------|
| **pkg**（Vercel） | 把 Node 项目打成单个 exe，内嵌 Node |
| **nexe** | 类似 pkg |
| **Bun** | `bun build --compile` 可编译为原生可执行文件（新方案） |

```bash
# pkg 示例
npm install -g pkg
pkg index.js --targets node18-win-x64 --output app.exe
```

适合「纯 Node、无浏览器界面」的脚本；**串口 GUI** 仍优先 Electron / Tauri / 系统 WebView 方案。

### 不算 exe、但和 JS 相关

| 方式 | 产物 | 说明 |
|------|------|------|
| **PWA** | 浏览器安装 | 无 exe，离线可用，难直接访问串口 |
| **Capacitor / Cordova** | 手机 App | 移动为主，不是 Windows exe |
| **浏览器扩展** | 插件包 | 受权限限制，不适合做通用串口工具 |

### JS 选型简记

- 要生态最全、不在乎体积 → **Electron**
- 要体积小、能接受 Rust 后端 → **Tauri**（本仓库已有）
- 要极小安装包、功能简单 → **Neutralinojs**
- 只会 Go + 会写网页 → **Wails**
- 只要命令行、无 UI → **pkg** / **Bun compile**

---

## Flutter（Dart）

**原理**：编译为原生代码 + 附带 Flutter 引擎资源。

```powershell
cd flutter-serial
flutter pub get
flutter build windows
# 输出: build\windows\x64\runner\Release\*.exe
```

---

## PHP（一般不打包 exe）

| 方式 | 说明 |
|------|------|
| 常规部署 | 服务器安装 PHP，上传 `.php` 文件 |
| Phar | `app.phar`，仍需本机 `php` 命令 |
| 桌面 exe | 第三方/套壳方案，**非主流** |

做串口桌面工具 **不推荐 PHP**；本仓库无 PHP 实现。

---

## 本仓库快速索引

| 子项目 | 打包命令摘要 | 产物示例 |
|--------|--------------|----------|
| `go-serial` | `go build -o xxx.exe` | `serial-debugger.exe` |
| `python-serial` | `py -3 -m PyInstaller ...` 或 `build.bat` | `dist/LogcatSerialDebugger.exe` |
| `rust-serial` | `cargo build --release` | `target/release/*.exe` |
| `cpp-serial` | `cmake` + `build` | `LogcatSerial.exe` |
| `csharp-serial` | `dotnet publish ...` | `publish/*.exe` |
| `tauri-serial` | `npm run tauri build` | `bundle/` 下安装包或 exe |
| `flutter-serial` | `flutter build windows` | `build/windows/.../Release/` |
| 全量脚本 | `./build.sh`（Linux/macOS 环境） | `builds/` 目录 |

更细的 Python 说明见：[python-serial/README.md](../python-serial/README.md)。

---

## 怎么选？

| 需求 | 建议 |
|------|------|
| 要小体积、单文件、无依赖 | **Go** / **Rust** |
| 快速做 GUI、逻辑多 | **Python + PyInstaller** 或 **C#** |
| 已有 Web 前端团队 | **Tauri** 或 **Electron** |
| 要强依赖 Android Studio | **Kotlin 插件**（不是 exe） |
| 做网站后台 | **PHP** 等，不必纠结 exe |

---

## 常见问题

**Q：动态语言打成 exe 后还是动态语言吗？**  
A：是。例如 Python exe 运行时仍是解释器执行字节码，只是用户看不到安装 Python 的步骤。

**Q：为什么 Python exe 比 Go exe 大很多？**  
A：Go 是单一编译产物；Python exe 内嵌了整个 Python + PyQt6。

**Q：exe 能在没装 SDK 的电脑上跑吗？**  
A：Go/Rust/C++ Release（静态或带好 DLL）、Python PyInstaller 单文件、.NET 自包含发布 — 一般可以；仅 `dotnet` 非自包含发布则需要对方安装 .NET。

---

*文档位置：`serial-client/docs/各语言打包exe方式总结.md`*
