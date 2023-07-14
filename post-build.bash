#!/bin/bash

# Get the file path argument
file_path=$1

# Check if the file exists
if [ ! -f "$file_path" ]; then
    echo "File does not exist: $file_path"
    exit 0
fi

# Extract the last folder name from the file path as the file name
file_name=$(basename "$(dirname "$file_path")")

# Create a template config file
timestamp=$(date +%Y%m%d%H%M%S)
config_file=".gon.temp.${timestamp}.hcl"

# Define the content of the config file using a here document
read -r -d '' config_content <<EOF
source = ["$file_path"]
bundle_id = "com.annatarhe.promptpal.cli"

apple_id {
  username = "iamhele1994@gmail.com"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Le He"
}

zip {
  output_path = "./dist/$file_name-mac.zip"
}
EOF

# Write the config content to the file
echo "$config_content" > "$config_file"

# Call 'gon' command with the config file
gon "$config_file"

# Store the exit code of the 'gon' command
gon_exit_code=$?

# Delete the temp config file
rm "$config_file"

# Check if the 'gon' command succeeded and show success info with emoji
if [ $gon_exit_code -eq 0 ]; then
    echo "Success! 😄"
else
    echo "Failed to run 'gon' command."
fi