use eframe::egui;
use serialport::SerialPort;
use std::sync::mpsc::{channel, Receiver, Sender};
use std::thread;
use std::time::Duration;

fn main() {
    let options = eframe::NativeOptions::default();
    eframe::run_native(
        "Logcat Serial Debugger - Rust",
        options,
        Box::new(|_cc| Box::new(SerialApp::default())),
    );
}

struct SerialApp {
    logs: Vec<String>,
    command: String,
    port_name: String,
    baud_rate: String,
    connected: bool,
}

impl Default for SerialApp {
    fn default() -> Self {
        Self {
            logs: Vec::new(),
            command: String::new(),
            port_name: String::from("/dev/ttyUSB0"),
            baud_rate: String::from("115200"),
            connected: false,
        }
    }
}

impl eframe::App for SerialApp {
    fn update(&mut self, ctx: &egui::Context, _frame: &mut eframe::Frame) {
        egui::CentralPanel::default().show(ctx, |ui| {
            ui.heading("Logcat Serial Debugger");
            
            ui.horizontal(|ui| {
                ui.label("Port:");
                ui.text_edit_singleline(&mut self.port_name);
                ui.label("Baud:");
                ui.text_edit_singleline(&mut self.baud_rate);
                
                if ui.button(if self.connected { "Disconnect" } else { "Connect" }).clicked() {
                    self.connected = !self.connected;
                }
            });
            
            ui.separator();
            
            egui::ScrollArea::vertical().show(ui, |ui| {
                for log in &self.logs {
                    ui.label(log);
                }
            });
            
            ui.separator();
            
            ui.horizontal(|ui| {
                ui.text_edit_singleline(&mut self.command);
                if ui.button("Send").clicked() {
                    self.logs.push(format!(">> {}", self.command));
                    self.command.clear();
                }
            });
        });
    }
}