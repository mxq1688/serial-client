using System;
using System.Collections.Generic;
using Microsoft.Maui.Controls;

namespace LogcatSerialDebugger
{
    public partial class MainPage : ContentPage
    {
        private List<string> logs = new List<string>();
        private bool isConnected = false;

        public MainPage()
        {
            InitializeComponent();
            LoadPorts();
        }

        private void LoadPorts()
        {
            var ports = new List<string> { "COM1", "COM2", "COM3", "/dev/ttyUSB0" };
            PortPicker.ItemsSource = ports;
            if (ports.Count > 0)
                PortPicker.SelectedIndex = 0;
            
            BaudPicker.SelectedIndex = 4; // 115200
            LevelPicker.SelectedIndex = 0; // V
        }

        private void OnConnectClicked(object sender, EventArgs e)
        {
            isConnected = !isConnected;
            ConnectBtn.Text = isConnected ? "Disconnect" : "Connect";
            
            if (isConnected)
            {
                AddLog("Connected to " + PortPicker.SelectedItem);
            }
            else
            {
                AddLog("Disconnected");
            }
        }

        private void OnSendClicked(object sender, EventArgs e)
        {
            if (isConnected && !string.IsNullOrEmpty(CommandInput.Text))
            {
                AddLog(">> " + CommandInput.Text);
                CommandInput.Text = "";
            }
        }

        private void OnClearClicked(object sender, EventArgs e)
        {
            logs.Clear();
            LogDisplay.Text = "";
        }

        private void AddLog(string message)
        {
            logs.Add(message);
            LogDisplay.Text += message + "\n";
        }
    }
}