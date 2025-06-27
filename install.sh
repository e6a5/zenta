#!/bin/bash

# zenta installation script
# Downloads and installs the latest zenta binary for your platform

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# GitHub repository
REPO="e6a5/zenta"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case $ARCH in
        x86_64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac
    
    case $OS in
        linux)
            PLATFORM="linux-${ARCH}"
            ;;
        darwin)
            PLATFORM="darwin-${ARCH}"
            ;;
        *)
            print_error "Unsupported operating system: $OS"
            print_warning "Please download manually from: https://github.com/${REPO}/releases"
            exit 1
            ;;
    esac
}

# Get latest release info
get_latest_release() {
    print_status "Fetching latest release information..."
    
    LATEST_RELEASE=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest")
    if [ $? -ne 0 ]; then
        print_error "Failed to fetch release information"
        print_warning "Please check your internet connection or try again later"
        exit 1
    fi
    
    # Check if we got an error response
    if echo "$LATEST_RELEASE" | grep -q '"message".*"Not Found"'; then
        print_error "No releases found for this repository"
        print_warning "Please download manually from: https://github.com/${REPO}/releases"
        print_warning "Or build from source: git clone https://github.com/${REPO}.git && cd zenta && make install-system"
        exit 1
    fi
    
    # Try to extract version with multiple methods
    VERSION=$(echo "$LATEST_RELEASE" | grep -o '"tag_name": "[^"]*' | cut -d'"' -f4)
    
    # Fallback method if first one fails
    if [ -z "$VERSION" ]; then
        VERSION=$(echo "$LATEST_RELEASE" | sed -n 's/.*"tag_name":[[:space:]]*"\([^"]*\)".*/\1/p')
    fi
    
    # Another fallback using jq-like parsing
    if [ -z "$VERSION" ]; then
        VERSION=$(echo "$LATEST_RELEASE" | tr ',' '\n' | grep '"tag_name"' | cut -d'"' -f4)
    fi
    
    if [ -z "$VERSION" ]; then
        print_error "Could not determine latest version from API response"
        print_warning "This might be due to:"
        print_warning "  - Network connectivity issues"
        print_warning "  - GitHub API rate limiting"
        print_warning "  - Unexpected API response format"
        print_warning ""
        print_warning "You can try:"
        print_warning "  1. Wait a few minutes and try again"
        print_warning "  2. Download manually: https://github.com/${REPO}/releases"
        print_warning "  3. Build from source: git clone https://github.com/${REPO}.git && cd zenta && make install-system"
        exit 1
    fi
    
    print_status "Latest version: $VERSION"
}

# Download and install
install_zenta() {
    FILENAME="zenta-${VERSION}-${PLATFORM}.tar.gz"
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"
    
    print_status "Downloading zenta for ${PLATFORM}..."
    print_status "URL: $DOWNLOAD_URL"
    
    # Create temporary directory
    TMP_DIR=$(mktemp -d)
    cd "$TMP_DIR"
    
    # Download the archive
    if ! curl -L -o "$FILENAME" "$DOWNLOAD_URL"; then
        print_error "Failed to download zenta"
        print_warning "Please download manually from: https://github.com/${REPO}/releases"
        exit 1
    fi
    
    # Extract the archive
    print_status "Extracting archive..."
    if ! tar -xzf "$FILENAME"; then
        print_error "Failed to extract archive"
        exit 1
    fi
    
    # Find the binary
    BINARY=$(find . -name "zenta-${VERSION}-${PLATFORM}" -type f)
    if [ -z "$BINARY" ]; then
        print_error "Binary not found in archive"
        exit 1
    fi
    
    # Make it executable
    chmod +x "$BINARY"
    
    # Ensure /usr/local/bin exists
    sudo mkdir -p /usr/local/bin
    
    # Install to /usr/local/bin
    print_status "Installing zenta to /usr/local/bin..."
    if ! sudo mv "$BINARY" /usr/local/bin/zenta; then
        print_error "Failed to install zenta (sudo required)"
        print_warning "You can manually copy the binary from: $TMP_DIR/$BINARY"
        exit 1
    fi
    
    # Clean up
    cd /
    rm -rf "$TMP_DIR"
    
    print_success "zenta ${VERSION} installed successfully!"
}

# Verify installation
verify_installation() {
    print_status "Verifying installation..."
    
    if command -v zenta >/dev/null 2>&1; then
        INSTALLED_VERSION=$(zenta --version 2>/dev/null | head -n1 || echo "unknown")
        print_success "zenta is installed: $INSTALLED_VERSION"
        print_status "Try it out: zenta now --quick"
    else
        print_error "Installation verification failed"
        print_warning "zenta command not found in PATH"
        exit 1
    fi
}

# Main installation flow
main() {
    echo -e "${BLUE}"
    echo "🧘 zenta installer"
    echo "=================="
    echo -e "${NC}"
    
    detect_platform
    print_status "Detected platform: $PLATFORM"
    
    get_latest_release
    install_zenta
    verify_installation
    
    echo
    print_success "Installation complete! 🎉"
    echo
    print_status "Quick start:"
    echo "  breath    # Quick breathing session"
    echo "  breathe   # Standard breathing session"
    echo "  reflect   # Evening reflection"
    echo
    print_status "For more info: https://github.com/${REPO}"
}

# Run the installer
main "$@" 
