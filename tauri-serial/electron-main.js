const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const { SerialPort } = require('serialport');
const { ReadlineParser } = require('@serialport/parser-readline');

let mainWindow;
let currentPort = null;

function createWindow() {
  mainWindow = new BrowserWindow({
    width: 1200,
    height: 800,
    webPreferences: {
      nodeIntegration: false,
      contextIsolation: true,
      preload: path.join(__dirname, 'preload.js')
    },
    icon: path.join(__dirname, 'icon.png'),
    title: 'Logcat Serial Debugger'
  });

  if (app.isPackaged) {
    mainWindow.loadFile(path.join(__dirname, 'dist', 'index.html'));
  } else {
    mainWindow.loadURL('http://localhost:5173');
    mainWindow.webContents.openDevTools();
  }
}

app.whenReady().then(() => {
  createWindow();
  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
  });
});

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

// Serial port handlers
ipcMain.handle('get-serial-ports', async () => {
  try {
    const ports = await SerialPort.list();
    return ports.map(port => port.path);
  } catch (error) {
    console.error('Error listing ports:', error);
    return [];
  }
});

ipcMain.handle('connect-serial', async (event, portPath, baudRate = 115200) => {
  try {
    if (currentPort) {
      currentPort.close();
    }

    currentPort = new SerialPort({
      path: portPath,
      baudRate: baudRate
    });

    const parser = currentPort.pipe(new ReadlineParser({ delimiter: '\n' }));
    
    parser.on('data', (data) => {
      mainWindow.webContents.send('serial-data', data);
    });

    currentPort.on('error', (err) => {
      console.error('Serial port error:', err);
      mainWindow.webContents.send('serial-error', err.message);
    });

    return { success: true };
  } catch (error) {
    console.error('Error connecting:', error);
    return { success: false, error: error.message };
  }
});

ipcMain.handle('disconnect-serial', async () => {
  try {
    if (currentPort) {
      currentPort.close();
      currentPort = null;
    }
    return { success: true };
  } catch (error) {
    return { success: false, error: error.message };
  }
});

ipcMain.handle('write-serial', async (event, data) => {
  try {
    if (currentPort && currentPort.isOpen) {
      currentPort.write(data);
      return { success: true };
    }
    return { success: false, error: 'Port not open' };
  } catch (error) {
    return { success: false, error: error.message };
  }
});