#!/bin/bash
METADATA_URL="https://raw.githubusercontent.com/VladimirsDuhovs/keyManager/main/installation_scripts/metadata.json"
APPLICATION="keyManager"

# Download the metadata.json file
curl -o metadata.json -sSL $METADATA_URL

# Get the download URL
URL=$(python -c "import json; print(json.load(open('metadata.json'))['amd64']['linux'])")

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
