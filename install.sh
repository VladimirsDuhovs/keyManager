#!/bin/bash

# Function to install SQLite3 on Ubuntu/Debian
function install_sqlite3_debian() {
    sudo apt-get update && sudo apt-get install sqlite3 libsqlite3-dev -y
}

# Function to install SQLite3 on CentOS
function install_sqlite3_centos() {
    sudo yum install sqlite sqlite-devel -y
}

# Function to install SQLite3 on Fedora
function install_sqlite3_fedora() {
    sudo dnf install sqlite sqlite-devel -y
}

# Function to install SQLite3 on macOS
function install_sqlite3_mac() {
    brew install sqlite
}

# Detect the OS
OS=$(uname)

# Install sqlite3 if they are not installed
if ! command -v sqlite3 &> /dev/null
then
    echo "sqlite3 is not installed, installing..."
    case $OS in
        "Linux")
            # Additional check for type of Linux distribution
            DISTRO=$(awk -F= '/^NAME/{print $2}' /etc/os-release)
            case $DISTRO in
                "\"Ubuntu\""| "\"Debian GNU/Linux\"")
                    install_sqlite3_debian
                    ;;
                "\"CentOS Linux\"")
                    install_sqlite3_centos
                    ;;
                "\"Fedora\"")
                    install_sqlite3_fedora
                    ;;
                *)
                    echo "Unsupported Linux distribution. Please install SQLite3 manually."
                    exit 1
                    ;;
            esac
            ;;
        "Darwin")
            install_sqlite3_mac
            ;;
        "Windows_NT")
            echo "Please install SQLite3 manually."
            ;;
        *)
            echo "Unsupported operating system. Please install SQLite3 manually."
            exit 1
            ;;
    esac
else
    echo "sqlite3 is already installed"
fi

# Move the binary to appropriate location
if [[ $OS == "Windows_NT" ]]; then
    echo "Moving binary to C:/Program Files"
    move key-manager_windows_amd64.exe "C:/Program Files"
else
    echo "Moving binary to /usr/local/bin"
    sudo mv key-manager_linux_amd64 /usr/local/bin
fi

# Give the binary execute permissions
if [[ $OS != "Windows_NT" ]]; then
    sudo chmod +x /usr/local/bin/key-manager_linux_amd64
fi

echo "Installation complete. You can now run the command 'key-manager_linux_amd64'"
