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

cd cble

# Create local configurations
cp config.example.yaml config.local.yaml
cp Caddyfile.example Caddyfile
cp docker-compose.yml docker-compose.local.yml

# Prompt for auto-configuration
read -p "Would you like to perform auto-configuration of CBLE? [Y/n] " auto_config

case $auto_config in
    [nN])
      # Give next steps instructions
      echo
      echo -e "\033[1;32mCBLE has been installed!\033[0m"
      echo
      echo -e "\033[1;31mPlease modify the local configurations this guide:\033[0m https://docs.cble.io/latest/getting-started/quick-start/#configure"
      echo
      echo -e "\033[0;33mOnce ready, start CBLE with:\033[0m"
      echo
      echo "     cd cble"
      echo "     docker compose -f docker-compose.local.yml build"
      echo "     docker compose -f docker-compose.local.yml up -d"
      echo
      exit 0
      ;;
    *)
      ;;
esac

# Replace values in config files

# Are you using SSL?
is_ssl=true
proto="https"
read -p "Are you using SSL (requires DNS to be set up)? [Y/n] " using_ssl
case $using_ssl in
  [nN])
    proto="http"
    is_ssl=false
    ;;
  *)
    sed -i "0,/ssl: false/{s/ssl: false/ssl: true/}" config.local.yaml
    ;;
esac

# Read in FQDN
domain=""
while [ -z "$domain" ]; do
  read -p "What is the domain (FQDN) you plan to use (e.g. cble.io): " domain
done
# Replace FQDN in config
sed -i "s/hostname: localhost/hostname: $domain/gm" config.local.yaml
sed -i "s,http://localhost:\(8080\|3000\),$proto://$domain,m" config.local.yaml
# Replace FQDN in Caddyfile
if ! $is_ssl; then
  sed -i "s/http:\/\/localhost/http:\/\/$domain/" Caddyfile
else
  sed -i "s/http:\/\/localhost/$domain/" Caddyfile
fi
# Replace FQDN in docker compose
sed -i "s/VITE_API_BASE_URL=http:\/\/localhost/VITE_API_BASE_URL=$proto:\/\/$domain/" docker-compose.local.yml


# Credit: https://unix.stackexchange.com/questions/462/how-to-create-strong-passwords-in-linux
function gen_pw(){
  cat /dev/urandom | tr -dc 'a-zA-Z0-9' | head -c $1
}

# Random passwords or no
read -p "Would you like to generate random passwords/keys? [Y/n] " rand_pass
# Read in all passwords
case $rand_pass in
  [nN])
    read -p "Please enter the database password (16+ chars recommended): " db_pass
    read -p "Please enter the JWT key (64+ chars recommended): " jwt_key
    read -p "Please enter the default admin password (16+ chars recommended): " admin_pass
    ;;
  *)
    db_pass=$(gen_pw 16)
    jwt_key=$(gen_pw 64)
    admin_pass=$(gen_pw 16)
    ;;
esac

# Place passwords in configs
sed -i -e ':a' -e 'N' -e '$!ba' -e "s/database:\n  username: cble\n  password: cble/database:\n  username: cble\n  password: $db_pass/g" config.local.yaml
sed -i -e "s/jwt_key: x*/jwt_key: $jwt_key/g" config.local.yaml
sed -i -e ':a' -e 'N' -e '$!ba' -e "s/password: cble\n  default_project: default/password: $admin_pass\n  default_project: default/g" config.local.yaml
# Place passwords in docker compose
sed -i -e "s/POSTGRES_PASSWORD=cble/POSTGRES_PASSWORD=$db_pass/" docker-compose.local.yml

# Read in default admin account details
read -p "Default admin account username [cble]: " username
if [ ! -z "$username" ]; then
  sed -i -e ':a' -e 'N' -e '$!ba' -e "s/last_name: Admin\n    username: cble/last_name: Admin\n    username: $username/g" config.local.yaml
else
  username="cble"
fi
read -p "Default admin account first name [CBLE]: " first_name
if [ ! -z "$first_name" ]; then
  sed -i -e "s/first_name: CBLE/first_name: $first_name/g" config.local.yaml
else
  first_name="CBLE"
fi
read -p "Default admin account last name [Admin]: " last_name
if [ ! -z "$last_name" ]; then
  sed -i -e "s/last_name: Admin/last_name: $last_name/g" config.local.yaml
else
  last_name="Admin"
fi

# Give next steps instructions
echo
echo -e "\033[1;32mCBLE has been installed and auto-configured!\033[0m"
echo
echo -e "\033[0;33mOnce ready, start CBLE with:\033[0m"
echo
echo "    cd cble"
echo "    docker compose -f docker-compose.local.yml build"
echo "    docker compose -f docker-compose.local.yml up -d"
echo
echo "Then log in with the following credentials:"
echo
echo "    Name: $first_name $last_name"
echo "    Username: $username"
echo "    Password: $admin_pass"
echo
exit 0