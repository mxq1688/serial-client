#!/usr/bin/env python3
import sys
from PyQt6.QtWidgets import *
from PyQt6.QtCore import *
from PyQt6.QtGui import *
from serial_worker import SerialWorker
from log_parser import parse_line
from logcat_colors import LOG_COLORS, DARK_THEME

class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.serial_worker = SerialWorker()
        self.serial_worker.data_received.connect(self.add_log)
        self.log_entries = []
        self.filter_level = 'V'
        self.filter_tag = ''
        self.init_ui()
        
    def init_ui(self):
        self.setWindowTitle("Logcat Serial Debugger")
        self.setGeometry(100, 100, 1200, 800)
        
        central = QWidget()
        self.setCentralWidget(central)
        layout = QVBoxLayout(central)
        
        # Toolbar
        toolbar = QHBoxLayout()
        
        self.port_combo = QComboBox()
        self.port_combo.addItems(SerialWorker.get_available_ports())
        toolbar.addWidget(QLabel("Port:"))
        toolbar.addWidget(self.port_combo)
        
        self.baud_combo = QComboBox()
        self.baud_combo.addItems(['115200', '9600', '19200'])
        toolbar.addWidget(QLabel("Baud:"))
        toolbar.addWidget(self.baud_combo)
        
        self.connect_btn = QPushButton("Connect")
        self.connect_btn.clicked.connect(self.toggle_connect)
        toolbar.addWidget(self.connect_btn)
        
        toolbar.addStretch()
        layout.addLayout(toolbar)
        
        # Log display
        self.log_display = QTextEdit()
        self.log_display.setReadOnly(True)
        layout.addWidget(self.log_display)
        
        # Apply dark theme
        self.setStyleSheet("QWidget {background: #1e1e1e; color: #d4d4d4;}")
    
    def toggle_connect(self):
        if self.connect_btn.text() == "Connect":
            port = self.port_combo.currentText()
            baud = int(self.baud_combo.currentText())
            if self.serial_worker.connect_port(port, baud):
                self.serial_worker.start()
                self.connect_btn.setText("Disconnect")
        else:
            self.serial_worker.disconnect()
            self.connect_btn.setText("Connect")
    
    def add_log(self, line):
        entry = parse_line(line)
        self.log_entries.append(entry)
        formatted = entry.format_logcat()
        color = LOG_COLORS.get(entry.level, '#FFFFFF')
        self.log_display.append(formatted)