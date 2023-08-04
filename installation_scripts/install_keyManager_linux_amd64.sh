#!/bin/bash
APPLICATION="keyManager"
URL="https://dl.cloudsmith.io/public/vladimirsduhovs/key-manager/raw/names/keyManager_linux_amd64/versions/5323e1b9d0f71f098ead38778da167df4e90c7ad/keyManager.tar.gz"

# Download the tarball
curl -o app.tar.gz -sSL $URL

# Extract to the desired location
sudo tar -C /usr/local/bin -xzf app.tar.gz

# Clean up
rm app.tar.gz
rm metadata.json

# Create alias
echo "alias $APPLICATION='/usr/local/bin/$APPLICATION'" >> ~/.bashrc
echo "alias $APPLICATION='/usr/local/bin/$APPLICATION'" >> ~/.zshrc

# Source the updated bashrc and zshrc
source ~/.bashrc
source ~/.zshrc
