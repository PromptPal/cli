#!/bin/bash

# Determine the OS and architecture
OS=$(uname -s)
ARCH=$(uname -m)

# Function to check if a command is available
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check for required commands
if ! command_exists curl; then
    echo "Error: curl is not installed."
    exit 1
fi

TARGET_FILE_NAME="https://github.com/PromptPal/cli/releases/latest/download/cli_"

cd /tmp

# Set the download URL based on the OS and architecture
if [[ "$OS" == "Darwin" ]]; then
    TARGET_FILE_NAME="${TARGET_FILE_NAME}${OS}"
    if [[ "$ARCH" == "x86_64" ]]; then
        URL="${TARGET_FILE_NAME}_amd64_v1.zip"
    elif [[ "$ARCH" == "arm64" ]]; then
        URL="${TARGET_FILE_NAME}_arm64.zip"
    else
        echo "Unsupported architecture: $ARCH on macOS"
        exit 1
    fi
    if ! command_exists unzip; then
        echo "Error: unzip is not installed."
        exit 1
    fi
elif [[ "$OS" == "Linux" ]]; then
    TARGET_FILE_NAME="${TARGET_FILE_NAME}${OS}"
    if [[ "$ARCH" == "x86_64" ]]; then
        URL="${TARGET_FILE_NAME}_x86_64.tar.gz"
    elif [[ "$ARCH" == "aarch64" ]]; then
        URL="${TARGET_FILE_NAME}_arm64.tar.gz"
    else
        echo "Unsupported architecture: $ARCH on Linux"
        exit 1
    fi
    if ! command_exists tar; then
        echo "Error: tar is not installed."
        exit 1
    fi
elif [[ "$OS" == "MINGW64_NT" ]] || [[ "$OS" == "MSYS_NT" ]] || [[ "$OS" == "CYGWIN_NT" ]]; then
    TARGET_FILE_NAME="${TARGET_FILE_NAME}Windows"
    if [[ "$ARCH" == "x86_64" ]]; then
        URL="${TARGET_FILE_NAME}_x86_64.zip"
    elif [[ "$ARCH" == "aarch64" ]]; then
        URL="${TARGET_FILE_NAME}_arm64.zip"
    else
        echo "Unsupported architecture: $ARCH on Windows"
        exit 1
    fi
    if ! command_exists unzip; then
        echo "Error: unzip is not installed."
        exit 1
    fi
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Download the file
FILENAME=$(basename "$URL")
curl -LO "$URL"

# Extract the file
if [[ "$FILENAME" == *.zip ]]; then
    unzip "$FILENAME"
elif [[ "$FILENAME" == *.tar.gz ]]; then
    tar zxvf "$FILENAME"
else
    echo "Unsupported file type: $FILENAME"
    exit 1
fi


# Move the binary to the appropriate location
if [[ "$OS" == "Darwin" ]] || [[ "$OS" == "Linux" ]]; then
    sudo mv promptpal /usr/local/bin
# elif [[ "$OS" == "MINGW64_NT" ]] || [[ "$OS" == "MSYS_NT" ]] || [[ "$OS" == "CYGWIN_NT" ]]; then
    # mv promptpal /c/Windows/System32/
fi

# Clean up
rm -f "$FILENAME"
rm -rf promptpal


# HELP WANTED
# I don't know where the `/bin` folder in windows. so i don't know where should the binaries be installed.
# if you know, please let me know.

# Success message
echo "Installation successful! You can try the command by running 'promptpal -h'."

if [[ "$OS" == "MINGW64_NT" ]] || [[ "$OS" == "MSYS_NT" ]] || [[ "$OS" == "CYGWIN_NT" ]]; then
	echo "Please note that the binaries are not installed yet. please move `/tmp/promptpal` to your `/bin/` folder manually."
	echo "btw if you know where should the binaries be installed, please raise an issue or pull request. (https://github.com/PromptPal/cli)."
fi