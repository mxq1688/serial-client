#include <QApplication>
#include <QMainWindow>
#include <QComboBox>
#include <QPushButton>
#include <QTextEdit>
#include <QLineEdit>
#include <QVBoxLayout>
#include <QHBoxLayout>
#include <QLabel>
#include <QSerialPort>
#include <QSerialPortInfo>
#include <QDateTime>
#include <QRegularExpression>

class LogcatSerialDebugger : public QMainWindow {
    Q_OBJECT

public:
    LogcatSerialDebugger(QWidget *parent = nullptr) : QMainWindow(parent) {
        setupUI();
        loadAvailablePorts();