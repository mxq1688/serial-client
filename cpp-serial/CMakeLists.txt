cmake_minimum_required(VERSION 3.16)
project(serial-debugger)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

find_package(Qt6 REQUIRED COMPONENTS Core Widgets SerialPort)

add_executable(serial-debugger
    main.cpp
    mainwindow.cpp
    mainwindow.h
)

target_link_libraries(serial-debugger
    Qt6::Core
    Qt6::Widgets
    Qt6::SerialPort
)