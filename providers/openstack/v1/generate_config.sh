#!/bin/bash

OUTPUT_FILE="config.yaml"

# Make sure we at least pass in a rc file
if [ -z "$1" ]; then
  echo "usage: generate_config.sh <openstack rc file> [output config filename]"
  exit 1
fi

# Set output filename if specified
if [ ! -z "$2" ]; then
  OUTPUT_FILE="$2"
fi

echo -e "\033[36mGenerating config from $1...\033[0m"

# Source the file to set the env vars
source $1

# Output the YAML config to output file
cat <<EOF > $OUTPUT_FILE
auth_url: $OS_AUTH_URL/v$OS_IDENTITY_API_VERSION
identity_version: $OS_IDENTITY_API_VERSION
username: $OS_USERNAME
password: $OS_PASSWORD
project_id: $OS_PROJECT_ID
project_name: $OS_PROJECT_NAME
domain_name: $OS_USER_DOMAIN_NAME
domain_id: $OS_PROJECT_DOMAIN_ID
EOF

echo -e "\033[32mGenerated Openstack provider config $OUTPUT_FILE!\033[0m"