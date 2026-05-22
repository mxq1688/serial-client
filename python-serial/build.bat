@echo off
cd /d "%~dp0"
echo Installing PyInstaller...
py -3 -m pip install pyinstaller -q
echo Building exe (may take several minutes)...
py -3 -m PyInstaller -w -F --name "LogcatSerialDebugger" --collect-all PyQt6 run.py
if %ERRORLEVEL% NEQ 0 exit /b %ERRORLEVEL%
echo.
echo Done: dist\LogcatSerialDebugger.exe
pause
