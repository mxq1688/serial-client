import { useState, useEffect, useRef } from 'react';
import './App.css';
import { parseLogcat } from './logcatParser';

function App() {
  const [isConnected, setIsConnected] = useState(false);
  const [logs, setLogs] = useState([]);
  const [portPath, setPortPath] = useState('/dev/ttyUSB0');
  const [availablePorts, setAvailablePorts] = useState([]);
  const [logLevel, setLogLevel] = useState('verbose');
  const logContainerRef = useRef(null);
  const isElectron = window.electronAPI !== undefined;

  useEffect(() => {
    if (logContainerRef.current) {
      logContainerRef.current.scrollTop = logContainerRef.current.scrollHeight;
    }
  }, [logs]);

  useEffect(() => {
    if (isElectron) {
      window.electronAPI.getSerialPorts().then(ports => {
        setAvailablePorts(ports);
        if (ports.length > 0 && !portPath) {
          setPortPath(ports[0]);
        }
      });

      window.electronAPI.onSerialData((data) => {
        const parsed = parseLogcat(data);
        if (parsed) {
          setLogs(prev => [...prev, parsed]);
        }
      });

      window.electronAPI.onSerialError((error) => {
        console.error('Serial error:', error);
        setIsConnected(false);
      });

      return () => {
        if (window.electronAPI.removeAllListeners) {
          window.electronAPI.removeAllListeners();
        }
      };
    }
  }, []);

  const handleConnect = async () => {
    if (isElectron) {
      const result = await window.electronAPI.connectSerial(portPath, 115200);
      setIsConnected(result.success);
      if (!result.success) alert('Failed to connect');
    } else {
      setIsConnected(true);
    }
    setLogs([]);
  };

  const handleDisconnect = async () => {
    if (isElectron) await window.electronAPI.disconnectSerial();
    setIsConnected(false);
  };

  const clearLogs = () => setLogs([]);

  const getLogLevelClass = (level) => {
    const map = {'V':'log-verbose','D':'log-debug','I':'log-info','W':'log-warning','E':'log-error'};
    return map[level] || 'log-verbose';
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>Logcat Serial Debugger</h1>
        <div className="controls">
          <input type="text" value={portPath} onChange={(e)=>setPortPath(e.target.value)} disabled={isConnected}/>
          <button onClick={isConnected ? handleDisconnect : handleConnect}>
            {isConnected ? 'Disconnect' : 'Connect'}
          </button>
          <select value={logLevel} onChange={(e)=>setLogLevel(e.target.value)}>
            <option value="verbose">Verbose</option>
            <option value="debug">Debug</option>
            <option value="info">Info</option>
            <option value="warning">Warning</option>
            <option value="error">Error</option>
          </select>
          <button onClick={clearLogs}>Clear</button>
        </div>
      </header>
      <div className="log-container" ref={logContainerRef}>
        {logs.map((log, i) => (
          <div key={i} className={`log-entry ${getLogLevelClass(log.level)}`}>
            <span className="log-time">{log.time}</span>
            <span className="log-level">[{log.level}]</span>
            <span className="log-tag">{log.tag}:</span>
            <span className="log-message">{log.message}</span>
          </div>
        ))}
        {logs.length === 0 && <div className="no-logs">No logs yet. Connect to start receiving data.</div>}
      </div>
    </div>
  );
}

export default App;