import serial
import serial.tools.list_ports
from PyQt6.QtCore import QThread, pyqtSignal

class SerialWorker(QThread):
    data_received = pyqtSignal(str)
    error_occurred = pyqtSignal(str)
    disconnected = pyqtSignal()
    
    def __init__(self):
        super().__init__()
        self.serial_port = None
        self.running = False
        
    def connect_port(self, port, baudrate):
        try:
            self.serial_port = serial.Serial(port, baudrate, timeout=0.1)
            self.running = True
            return True, "Success"
        except Exception as e:
            error_msg = f"Connection error: {e}"
            print(error_msg)
            return False, error_msg
    
    def disconnect(self):
        self.running = False
        if self.serial_port:
            try:
                self.serial_port.close()
            except Exception:
                pass
            self.serial_port = None
    
    def send_command(self, command):
        if self.serial_port and self.serial_port.is_open:
            try:
                self.serial_port.write((command + '\n').encode('utf-8'))
            except Exception as e:
                self.error_occurred.emit(f"Write error: {e}")
    
    def run(self):
        while self.running:
            if self.serial_port and self.serial_port.is_open:
                try:
                    if self.serial_port.in_waiting:
                        data = self.serial_port.readline().decode('utf-8', errors='ignore').strip()
                        if data:
                            self.data_received.emit(data)
                except serial.SerialException as e:
                    self.error_occurred.emit(f"Serial exception: {e}")
                    self.disconnect()
                    self.disconnected.emit()
                    break
                except Exception as e:
                    self.error_occurred.emit(f"Read error: {e}")
                    self.disconnect()
                    self.disconnected.emit()
                    break

    @staticmethod
    def get_available_ports():
        return [port.device for port in serial.tools.list_ports.comports()]