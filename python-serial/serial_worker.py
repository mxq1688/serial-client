import serial
import serial.tools.list_ports
from PyQt6.QtCore import QThread, pyqtSignal

class SerialWorker(QThread):
    data_received = pyqtSignal(str)
    
    def __init__(self):
        super().__init__()
        self.serial_port = None
        self.running = False
        
    def connect_port(self, port, baudrate):
        try:
            self.serial_port = serial.Serial(port, baudrate, timeout=0.1)
            self.running = True
            return True
        except Exception as e:
            print(f"Connection error: {e}")
            return False
    
    def disconnect(self):
        self.running = False
        if self.serial_port:
            self.serial_port.close()
            self.serial_port = None
    
    def send_command(self, command):
        if self.serial_port and self.serial_port.is_open:
            self.serial_port.write((command + '\n').encode())
    
    def run(self):
        while self.running:
            if self.serial_port and self.serial_port.in_waiting:
                try:
                    data = self.serial_port.readline().decode('utf-8', errors='ignore').strip()
                    if data:
                        self.data_received.emit(data)
                except Exception:
                    pass

    @staticmethod
    def get_available_ports():
        return [port.device for port in serial.tools.list_ports.comports()]