# pasteCTL

A command-line interface for pasteCTL - share code snippets directly from your terminal.

[![Go Report Card](https://goreportcard.com/badge/github.com/Sumedhvats/pasteCTL_cli)](https://goreportcard.com/report/github.com/Sumedhvats/pasteCTL_cli)
[![Go Reference](https://pkg.go.dev/badge/github.com/Sumedhvats/pasteCTL_cli.svg)](https://pkg.go.dev/github.com/Sumedhvats/pasteCTL_cli)
[![Release](https://img.shields.io/github/v/release/Sumedhvats/pasteCTL_cli)](https://github.com/Sumedhvats/pasteCTL_cli/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

pastectl is a command-line client for the pasteCTL platform that enables developers to share code snippets and text directly from their terminal. Built with Go and designed for seamless integration into development workflows.

## Features

- 📁 File-based paste creation from any file in your filesystem
- ✏️ Interactive editor integration for creating pastes
- 🔍 Automatic language detection based on file extensions
- ⏰ Configurable expiration times for paste lifecycle management
- 📤 Raw output support for scripting and automation
- 🌐 Cross-platform compatibility (Linux, macOS, Windows)

## Installation

### Using Go

```bash
go install github.com/Sumedhvats/pastectl@latest
```

### Binary Releases

Download pre-compiled binaries from [GitHub Releases](https://github.com/Sumedhvats/pasteCTL_cli/releases).

### Building From Source

```bash
git clone https://github.com/Sumedhvats/pasteCTL_cli.git
cd pasteCTL_cli
go build -o pastectl ./cmd/pastectl
```

After building, you'll have a `pastectl` binary (or `pastectl.exe` on Windows) in your current directory.

## Adding Binary to PATH

To use `pastectl` from anywhere in your terminal, add it to your system PATH:

### Linux

```bash
# Move the binary to a directory in your PATH
sudo mv pastectl /usr/local/bin/

# Or add the current directory to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/path/to/pasteCTL_cli
```

**Verify installation:**
```bash
which pastectl
pastectl --help
```

### macOS

```bash
# Move the binary to a directory in your PATH
sudo mv pastectl /usr/local/bin/

# Or add to PATH (add to ~/.zshrc or ~/.bash_profile)
export PATH=$PATH:/path/to/pasteCTL_cli

# You may need to allow the binary in System Preferences > Security & Privacy
```

**Verify installation:**
```bash
which pastectl
pastectl --help
```

### Windows

**Method 1: Move to System Directory**
```powershell
# Run PowerShell as Administrator
Move-Item pastectl.exe C:\Windows\System32\
```

**Method 2: Add Directory to PATH (Recommended)**

1. Open System Properties:
   - Press `Win + X` and select "System"
   - Click "Advanced system settings"
   - Click "Environment Variables"

2. Edit PATH variable:
   - Under "User variables" or "System variables", find `Path`
   - Click "Edit" → "New"
   - Add the full path to the directory containing `pastectl.exe`
   - Click "OK" to save

3. Restart your terminal/PowerShell

**Verify installation:**
```powershell
where pastectl
pastectl --help
```

## Initial Configuration

Before using pastectl, configure the frontend and backend URLs:

```bash
# Set the frontend URL (where pastes will be accessed)
pastectl config set frontend_url https://paste.example.com

# Set the backend API URL
pastectl config set backend_url https://api.paste.example.com
```

**Configuration file location:**
- Linux/macOS: `~/.pastectl.yaml`
- Windows: `%USERPROFILE%\.pastectl.yaml`

You can override the config location with the `PASTECTL_CONFIG` environment variable.

## Usage Guide

### Creating Pastes

#### Create from a File

```bash
# Basic file paste
pastectl create --file main.go

# With custom language and expiration
pastectl create --file script.py --language python --expire 24h

# Short flags
pastectl create -f config.json -l json -e 1h
```

#### Create Using an Editor

```bash
# Opens your default editor (set via $EDITOR environment variable)
pastectl create

# Will use vi/vim by default if $EDITOR is not set
```

**Setting your preferred editor:**

```bash
# Linux/macOS (add to ~/.bashrc or ~/.zshrc)
export EDITOR=nano        # or vim, emacs, code, etc.

# Windows (PowerShell)
$env:EDITOR = "notepad"   # or code, notepad++, etc.
```

#### Language Detection

The CLI automatically detects language from file extensions:

| Extension | Language | Extension | Language |
|-----------|----------|-----------|----------|
| `.go` | go | `.js` | javascript |
| `.py` | python | `.java` | java |
| `.c` | c | `.cpp` | cpp |
| `.json` | json | `.md` | markdown |
| `.txt` | plain | | |

**Override detection:**
```bash
pastectl create -f script.sh -l bash
```

#### Setting Expiration

Available expiration formats:
- `10m` - 10 minutes
- `1h` - 1 hour  
- `24h` - 24 hours
- `7d` - 7 days
- `never` - No expiration (default)

```bash
# Temporary paste (1 hour)
pastectl create -f debug.log -e 1h

# Week-long paste
pastectl create -f report.txt -e 7d

# Permanent paste (default)
pastectl create -f reference.md
```

### Retrieving Pastes

#### Get Paste with Metadata

```bash
pastectl get abc123def
```

**Output:**
```
--- Paste Details ---
ID:       abc123def
Language: python
Created:  2025-01-15 14:30:22
--- Content ---
def hello():
    print("Hello, World!")
```

#### Get Raw Content Only

```bash
# Display raw content
pastectl get abc123def --raw

# Save to file
pastectl get abc123def --raw > local-copy.py

# Pipe to other commands
pastectl get abc123def --raw | grep "TODO"
```

### Updating Pastes

Update an existing paste using your editor:

```bash
pastectl update abc123def
```

This will:
1. Fetch the current paste content
2. Open it in your default editor
3. Submit the updated content when you save and close the editor

**Note:** The language of the paste is preserved during updates.

### Configuration Management

View and modify CLI configuration:

```bash
# Set configuration values
pastectl config set frontend_url https://paste.example.com
pastectl config set backend_url https://api.paste.example.com

# Configuration keys:
# - frontend_url: Base URL for accessing pastes in browser
# - backend_url: API endpoint for paste operations
```

## Practical Examples

### Share Build Logs

```bash
# Capture and share build output
make build 2>&1 | tee build.log
pastectl create -f build.log --expire 7d
```

### Share Git Diffs

```bash
# Share uncommitted changes
git diff > changes.diff
pastectl create -f changes.diff -l diff -e 1h

# Or pipe directly
git diff | pastectl create -l diff -e 1h
```

### Quick Code Sharing

```bash
# Share a specific file with team
pastectl create -f src/utils/helper.js -e 24h

# Create and copy to clipboard (Linux with xclip)
pastectl create -f main.go | grep "Link:" | awk '{print $3}' | xclip -selection clipboard
```

### System Administration

```bash
# Share configuration files
pastectl create -f /etc/nginx/nginx.conf -e 1h

# Share recent logs
tail -n 100 /var/log/syslog | pastectl create -l log -e 24h
```

### Code Review Workflow

```bash
# Create paste for review
pastectl create -f feature.go -e 7d

# After review, update with changes
pastectl update abc123def

# Share final version
pastectl get abc123def --raw > reviewed-feature.go
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `EDITOR` | Default text editor for creating/updating pastes | `vi` |
| `PASTECTL_CONFIG` | Custom config file path | `~/.pastectl.yaml` |

## Troubleshooting

### Command not found
- Ensure the binary is in your PATH (see installation section)
- Try using the full path: `/path/to/pastectl` or `./pastectl`

### Config not set error
```bash
# You'll see: "Error: frontend_url is not set"
# Solution:
pastectl config set frontend_url https://your-paste-url.com
pastectl config set backend_url https://your-api-url.com
```

### Editor not opening
```bash
# Set your preferred editor
export EDITOR=nano  # or vim, code, etc.
```

### Permission denied (Linux/macOS)
```bash
chmod +x pastectl
```

## Contributing

We welcome contributions to pastectl. Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
git clone https://github.com/Sumedhvats/pasteCTL_cli.git
cd pasteCTL_cli
go mod download
go build ./...
```

### Running Tests

```bash
go test ./...
```

## Security

For security concerns, please see our [Security Policy](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [pasteCTL](https://github.com/Sumedhvats/pasteCTL) - Web application and API server
- [pasteCTL Documentation](https://github.com/Sumedhvats/pasteCTL/docs) - Complete documentation

## Support

- **Issues**: [GitHub Issues](https://github.com/Sumedhvats/pasteCTL_cli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/Sumedhvats/pasteCTL_cli/discussions)
- **Documentation**: [Wiki](https://github.com/Sumedhvats/pasteCTL_cli/wiki)

---

Made with ❤️ by Sumedh