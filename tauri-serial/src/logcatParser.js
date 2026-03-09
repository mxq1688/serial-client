export function parseLogcat(line) {
  // Parse Android logcat format: MM-DD HH:MM:SS.mmm PID TID LEVEL TAG: MESSAGE
  const regex = /^(\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+)\s+(\d+)\s+([VDIWEF])\s+([^:]+)\s*:\s*(.*)$/;
  const match = line.match(regex);
  
  if (match) {
    return {
      time: match[1],
      pid: match[2],
      tid: match[3],
      level: match[4],
      tag: match[5].trim(),
      message: match[6]
    };
  }
  
  // Try to parse as simple log if not in standard format
  if (line.trim()) {
    return {
      time: new Date().toLocaleTimeString(),
      pid: '-',
      tid: '-',
      level: 'I',
      tag: 'Serial',
      message: line
    };
  }
  
  return null;
}