import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Logcat Serial Debugger',
      theme: ThemeData.dark(),
      home: SerialDebugger(),
    );
  }
}

class SerialDebugger extends StatefulWidget {
  @override
  _SerialDebuggerState createState() => _SerialDebuggerState();
}

class _SerialDebuggerState extends State<SerialDebugger> {
  List<String> logs = [];
  bool connected = false;
  String port = '/dev/ttyUSB0';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Logcat Serial Debugger - Flutter'),
        backgroundColor: Color(0xFF2D2D30),
      ),
      backgroundColor: Color(0xFF1E1E1E),
      body: Column(
        children: [
          Container(
            padding: EdgeInsets.all(10),
            color: Color(0xFF2D2D30),
            child: Row(
              children: [
                Text('Port: ', style: TextStyle(color: Colors.white)),
                SizedBox(width: 10),
                Expanded(
                  child: TextField(
                    decoration: InputDecoration(
                      hintText: port,
                      filled: true,
                      fillColor: Color(0xFF3C3C3C),
                    ),
                    style: TextStyle(color: Colors.white),
                  ),
                ),
                SizedBox(width: 10),
                ElevatedButton(
                  onPressed: () {
                    setState(() {
                      connected = !connected;
                      logs.add(connected ? 'Connected' : 'Disconnected');
                    });
                  },
                  child: Text(connected ? 'Disconnect' : 'Connect'),
                ),
              ],
            ),
          ),
          Expanded(
            child: Container(
              padding: EdgeInsets.all(10),
              child: ListView.builder(
                itemCount: logs.length,
                itemBuilder: (context, index) {
                  return Text(
                    logs[index],
                    style: TextStyle(color: Colors.green, fontFamily: 'monospace'),
                  );
                },
              ),
            ),
          ),
        ],
      ),
    );
  }
}