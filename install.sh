#!/bin/bash

# Function to install SQLite3 and libsqlite3-dev on Ubuntu/Debian
function install_sqlite3_debian() {
    sudo apt-get update && sudo apt-get install sqlite3 libsqlite3-dev -y
}

# Function to install SQLite3 and libsqlite3-dev on CentOS
function install_sqlite3_centos() {
    sudo yum install sqlite sqlite-devel -y
}

# Function to install SQLite3 and libsqlite3-dev on Fedora
function install_sqlite3_fedora() {
    sudo dnf install sqlite sqlite-devel -y
}

# Detect the OS
OS=$(awk -F= '/^NAME/{print $2}' /etc/os-release)

# Install sqlite3 and libsqlite3-dev if they are not installed
if ! command -v sqlite3 &> /dev/null
then
    echo "sqlite3 is not installed, installing..."
    case $OS in
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
            echo "Unsupported operating system. Please install SQLite3 manually."
            exit 1
            ;;
    esac
else
    echo "sqlite3 is already installed"
fi

# Move the binary to /usr/local/bin
echo "Moving binary to /usr/local/bin"
sudo mv key-app /usr/local/bin

# Give the binary execute permissions
sudo chmod +x /usr/local/bin/key-app

echo "Installation complete. You can now run the command 'key-app'"
