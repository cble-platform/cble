#!/bin/bash

# Function to check if a command is already installed (from docker install script)
command_exists() {
  command -v "$@" > /dev/null 2>&1
}

# Install Docker using docker auto-install script (if not installed)
if command_exists docker; then
  echo "Docker already installed... skipping install"
else
  echo "Installing Docker..."
  curl -fsSL https://get.docker.com | sh
fi

# Ensure git is installed
if command_exists git; then
  # Clone the CBLE repo
  git clone https://github.com/cble-platform/cble
else
  cat >&2 <<-'EOF'
	Error: "git" was not found. Please install "git" and re-run this installer.
	EOF
  exit 1
fi

# Create local configurations
cp cble/config.example.yaml cble/config.local.yaml
cp cble/Caddyfile.example cble/Caddyfile
cp cble/docker-compose.yml cble/docker-compose.local.yml

# Give next steps instructions
echo "CBLE has been installed!"
echo
echo "Please modify the local configurations this guide: https://docs.cble.io/latest/getting-started/installation/quick-start/"
echo
echo "Once ready, start CBLE with:"
echo
echo "     cd cble"
echo "     docker compose -f docker-compose.local.yml build"
echo "     docker compose -f docker-compose.local.yml up -d"
echo
exit 0