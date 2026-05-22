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
        self.serial_worker.error_occurred.connect(self.show_error)
        self.serial_worker.disconnected.connect(self.handle_disconnect)
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
        self.baud_combo.addItems(['2000000', '1500000', '921600', '115200', '57600', '38400', '19200', '9600'])
        toolbar.addWidget(QLabel("Baud:"))
        toolbar.addWidget(self.baud_combo)
        
        self.connect_btn = QPushButton("Connect")
        self.connect_btn.clicked.connect(self.toggle_connect)
        toolbar.addWidget(self.connect_btn)
        
        toolbar.addStretch()
        layout.addLayout(toolbar)
        
        # Filter Bar
        filter_layout = QHBoxLayout()
        
        self.level_combo = QComboBox()
        self.level_combo.addItems(['Verbose', 'Debug', 'Info', 'Warn', 'Error', 'Fatal'])
        self.level_combo.currentTextChanged.connect(self.apply_filter)
        filter_layout.addWidget(QLabel("Level:"))
        filter_layout.addWidget(self.level_combo)
        
        self.search_input = QLineEdit()
        self.search_input.setPlaceholderText("Search (e.g. tag:MyTag or match text...)")
        self.search_input.textChanged.connect(self.apply_filter)
        filter_layout.addWidget(self.search_input)
        
        self.clear_log_btn = QPushButton("Clear")
        self.clear_log_btn.clicked.connect(self.clear_logs)
        filter_layout.addWidget(self.clear_log_btn)
        
        layout.addLayout(filter_layout)
        
        # Log display
        self.log_display = QTextEdit()
        self.log_display.setReadOnly(True)
        layout.addWidget(self.log_display)
        
        # Command input
        command_layout = QHBoxLayout()
        self.command_input = QLineEdit()
        self.command_input.setPlaceholderText("Enter command to send...")
        self.command_input.returnPressed.connect(self.send_command)
        self.send_btn = QPushButton("Send")
        self.send_btn.clicked.connect(self.send_command)
        
        command_layout.addWidget(self.command_input)
        command_layout.addWidget(self.send_btn)
        layout.addLayout(command_layout)
        
        # Apply dark theme
        self.setStyleSheet("QWidget {background: #1e1e1e; color: #d4d4d4;}")
    
    def toggle_connect(self):
        if self.connect_btn.text() == "Connect":
            port = self.port_combo.currentText()
            if not port:
                QMessageBox.warning(self, "Warning", "No port selected.")
                return
            baud = int(self.baud_combo.currentText())
            success, msg = self.serial_worker.connect_port(port, baud)
            if success:
                self.serial_worker.start()
                self.connect_btn.setText("Disconnect")
            else:
                QMessageBox.critical(self, "Connection Error", msg)
        else:
            self.serial_worker.disconnect()
            self.connect_btn.setText("Connect")
            
    def handle_disconnect(self):
        self.connect_btn.setText("Connect")
        
    def show_error(self, error_msg):
        self.add_log(f"E/SerialError(0): {error_msg}")
        
    def send_command(self):
        cmd = self.command_input.text()
        if cmd:
            self.serial_worker.send_command(cmd)
            self.add_log(f"D/SerialSend(0): {cmd}")
            self.command_input.clear()

    def add_log(self, line):
        entry = parse_line(line)
        self.log_entries.append(entry)
        if self.should_show_log(entry):
            self.append_log_entry_to_display(entry)
            
    def apply_filter(self):
        self.log_display.clear()
        for entry in self.log_entries:
            if self.should_show_log(entry):
                self.append_log_entry_to_display(entry)

    def should_show_log(self, entry):
        # 1. Level filter
        level_map = {'V': 0, 'D': 1, 'I': 2, 'W': 3, 'E': 4, 'F': 5, 'S': 6}
        entry_level_idx = level_map.get(entry.level, 0)
        filter_level_idx = self.level_combo.currentIndex()
        if entry_level_idx < filter_level_idx:
            return False
            
        # 2. Text/Tag filter
        search_text = self.search_input.text().strip().lower()
        if not search_text:
            return True
            
        if search_text.startswith("tag:"):
            target_tag = search_text[4:].strip()
            return target_tag in entry.tag.lower()
            
        return search_text in entry.message.lower() or search_text in entry.tag.lower()

    def append_log_entry_to_display(self, entry):
        formatted = entry.format_logcat()
        color = LOG_COLORS.get(entry.level, '#FFFFFF')
        
        # Escape HTML chars to prevent rendering issues
        formatted = formatted.replace('&', '&amp;').replace('<', '&lt;').replace('>', '&gt;')
        html_formatted = f'<span style="color: {color}; white-space: pre-wrap;">{formatted}</span>'
        self.log_display.append(html_formatted)

    def clear_logs(self):
        self.log_entries.clear()
        self.log_display.clear()