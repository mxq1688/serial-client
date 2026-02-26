// Arduino Serial Debugger Test Code
// Supports color-coded log levels

void setup() {
  Serial.begin(115200);
  pinMode(LED_BUILTIN, OUTPUT);
  
  Serial.println("[INFO] Arduino Serial Debugger Test Started");
  Serial.println("[INFO] Firmware Version: 1.0.0");
  Serial.println("[DEBUG] Setup completed");
}

void loop() {
  static unsigned long counter = 0;
  static unsigned long lastTime = 0;
  unsigned long currentTime = millis();
  
  // Send different log levels every second
  if (currentTime - lastTime >= 1000) {
    lastTime = currentTime;
    counter++;
    
    // Cycle through different log levels
    switch(counter % 5) {
      case 0:
        Serial.print("[DEBUG] Counter value: ");
        Serial.println(counter);
        break;
      case 1:
        Serial.print("[INFO] System running for: ");
        Serial.print(currentTime / 1000);
        Serial.println(" seconds");
        break;
      case 2:
        Serial.println("[WARNING] This is a warning message");
        break;
      case 3:
        Serial.println("[ERROR] Simulated error for testing");
        break;
      case 4:
        // Send sensor data
        float temp = 20.0 + random(-50, 50) / 10.0;
        Serial.print("[DATA] Temperature: ");
        Serial.print(temp);
        Serial.println(" C");
        break;
    }
    
    // Toggle LED
    digitalWrite(LED_BUILTIN, counter % 2);
  }
  
  // Check for commands
  if (Serial.available() > 0) {
    String cmd = Serial.readStringUntil('\n');
    cmd.trim();
    
    if (cmd == "AT") {
      Serial.println("OK");
    } else if (cmd == "AT+VERSION") {
      Serial.println("Version: 1.0.0");
    } else if (cmd == "AT+RST") {
      Serial.println("[INFO] Resetting...");
      delay(100);
      setup();
    } else if (cmd.startsWith("LED:")) {
      int state = cmd.substring(4).toInt();
      digitalWrite(LED_BUILTIN, state);
      Serial.print("[INFO] LED set to: ");
      Serial.println(state ? "ON" : "OFF");
    } else if (cmd == "SENSOR:READ") {
      float temp = 20.0 + random(-50, 50) / 10.0;
      float humidity = 40.0 + random(-100, 100) / 10.0;
      Serial.print("[DATA] Temp: ");
      Serial.print(temp);
      Serial.print(" C, Humidity: ");
      Serial.print(humidity);
      Serial.println(" %");
    } else {
      Serial.print("[ERROR] Unknown command: ");
      Serial.println(cmd);
    }
  }
}