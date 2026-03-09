import re
from datetime import datetime

class LogEntry:
    def __init__(self, level='I', tag='Serial', message='', pid=0):
        self.timestamp = datetime.now()
        self.level = level
        self.tag = tag
        self.message = message
        self.pid = pid
        self.tid = 0
    
    def format_logcat(self):
        time_str = self.timestamp.strftime('%m-%d %H:%M:%S.%f')[:-3]
        return f"{time_str} {self.pid:5d} {self.tid:5d} {self.level} {self.tag:12s}: {self.message}"

def parse_line(line):
    # D/TAG(1234): Message
    match = re.match(r'^([VDIWEFS])/([^(]+)\((\d+)\): (.+)$', line)
    if match:
        return LogEntry(match.group(1), match.group(2).strip(), match.group(4), int(match.group(3)))
    # [D] TAG: Message
    match = re.match(r'^\[([VDIWEFS])\] ([^:]+): (.+)$', line)
    if match:
        return LogEntry(match.group(1), match.group(2).strip(), match.group(3))
    # Plain text
    return LogEntry('I', 'Serial', line)