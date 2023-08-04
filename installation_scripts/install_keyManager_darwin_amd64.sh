#!/bin/bash
APPLICATION="keyManager"
URL="__PACKAGE_URL__"

# Download the tarball
curl -o app.tar.gz -sSL $URL

# Extract to the desired location
tar -C /usr/local/bin -xzf app.tar.gz

# Clean up
rm app.tar.gz
rm metadata.json

# Create alias
echo "alias $APPLICATION='/usr/local/bin/$APPLICATION'" >> ~/.bashrc
echo "alias $APPLICATION='/usr/local/bin/$APPLICATION'" >> ~/.zshrc

# Source the updated bashrc and zshrc
source ~/.bashrc
source ~/.zshrc
