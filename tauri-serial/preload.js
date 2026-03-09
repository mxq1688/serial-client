const { contextBridge, ipcRenderer } = require('electron');

// Expose protected methods that allow the renderer process
// to communicate with the main process for serial port operations
contextBridge.exposeInMainWorld('electronAPI', {
  // Get list of available serial ports
  getSerialPorts: () => ipcRenderer.invoke('get-serial-ports'),
  
  // Connect to a serial port
  connectSerial: (portPath, baudRate) => 
    ipcRenderer.invoke('connect-serial', portPath, baudRate),
  
  // Disconnect from serial port
  disconnectSerial: () => ipcRenderer.invoke('disconnect-serial'),
  
  // Write data to serial port
  writeSerial: (data) => ipcRenderer.invoke('write-serial', data),
  
  // Listen for serial data
  onSerialData: (callback) => {
    ipcRenderer.on('serial-data', (event, data) => callback(data));
  },
  
  // Listen for serial errors
  onSerialError: (callback) => {
    ipcRenderer.on('serial-error', (event, error) => callback(error));
  },
  
  // Remove listeners
  removeAllListeners: () => {
    ipcRenderer.removeAllListeners('serial-data');
    ipcRenderer.removeAllListeners('serial-error');
  }
});