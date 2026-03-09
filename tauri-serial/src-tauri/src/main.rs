#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

use serialport::{available_ports, SerialPort};
use std::sync::Mutex;
use std::time::Duration;
use tauri::State;

struct SerialState {
    port: Mutex<Option<Box<dyn SerialPort>>>,
}

#[tauri::command]
fn get_serial_ports() -> Vec<String> {
    match available_ports() {
        Ok(ports) => ports.iter().map(|p| p.port_name.clone()).collect(),
        Err(_) => vec![],
    }
}

#[tauri::command]
fn connect_serial(state: State<SerialState>, port: String, baud_rate: u32) -> Result<(), String> {
    let port_result = serialport::new(port, baud_rate)
        .timeout(Duration::from_millis(100))
        .open();
    
    match port_result {
        Ok(port) => {
            *state.port.lock().unwrap() = Some(port);
            Ok(())
        }
        Err(e) => Err(e.to_string()),
    }
}

#[tauri::command]
fn disconnect_serial(state: State<SerialState>) -> Result<(), String> {
    *state.port.lock().unwrap() = None;
    Ok(())
}

#[tauri::command]
fn read_serial(state: State<SerialState>) -> Result<String, String> {
    let mut port_guard = state.port.lock().unwrap();
    if let Some(port) = port_guard.as_mut() {
        let mut buffer = vec![0; 1024];
        match port.read(&mut buffer) {
            Ok(n) => {
                let data = String::from_utf8_lossy(&buffer[..n]).to_string();
                Ok(data)
            }
            Err(_) => Ok(String::new()),
        }
    } else {
        Err("Not connected".to_string())
    }
}

#[tauri::command]
fn write_serial(state: State<SerialState>, data: String) -> Result<(), String> {
    let mut port_guard = state.port.lock().unwrap();
    if let Some(port) = port_guard.as_mut() {
        match port.write(data.as_bytes()) {
            Ok(_) => Ok(()),
            Err(e) => Err(e.to_string()),
        }
    } else {
        Err("Not connected".to_string())
    }
}

fn main() {
    tauri::Builder::default()
        .manage(SerialState {
            port: Mutex::new(None),
        })
        .invoke_handler(tauri::generate_handler![
            get_serial_ports,
            connect_serial,
            disconnect_serial,
            read_serial,
            write_serial
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}