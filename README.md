# PasteCTL CLI

A command-line interface for PasteCTL, allowing developers to create, retrieve, and manage code snippets directly from the terminal.

## Overview

PasteCTL CLI is a terminal-based client for the PasteCTL platform. It provides a fast and efficient way to share code snippets without leaving your development environment. The CLI integrates with your system's default editor and supports automatic language detection based on file extensions.

## Features

- **Create Pastes**: Create pastes from files or directly in your editor
- **Automatic Language Detection**: Detects language from file extensions
- **Custom Expiry Times**: Set paste expiration (10m, 1h, never, etc.)
- **Retrieve Pastes**: Fetch paste content with formatted output or raw text
- **Update Pastes**: Edit existing pastes in your preferred editor
- **Configuration Management**: Store API endpoint and frontend URL locally
- **Cross-Platform**: Works on Linux, macOS, and Windows

## Installation

### Go Install

If you have Go 1.24.0 or higher installed:

```bash
go install github.com/Sumedhvats/pastectl@latest
```

Make sure your `$GOPATH/bin` is in your `PATH`.

### Pre-built Binaries

Download the appropriate binary for your platform from the [releases page](https://github.com/Sumedhvats/pastectl/releases).

#### Linux

```bash
wget https://github.com/Sumedhvats/pastectl/releases/download/v0.1.1/pastectl_0.1.1_linux_amd64.tar.gz
tar -xvf pastectl_0.1.1_linux_amd64.tar.gz
sudo mv pastectl /usr/local/bin/
```

Verify installation:
```bash
pastectl --help
```

#### macOS

```bash
wget https://github.com/Sumedhvats/pastectl/releases/download/v0.1.1/pastectl_0.1.1_darwin_amd64.tar.gz
tar -xvf pastectl_0.1.1_darwin_amd64.tar.gz
sudo mv pastectl /usr/local/bin/
```

Verify installation:
```bash
pastectl --help
```

#### Windows

1. Download `pastectl_0.1.1_windows_amd64.zip` from the [releases page](https://github.com/Sumedhvats/pastectl/releases)
2. Extract the archive
3. Move `pastectl.exe` to a directory in your `PATH` (e.g., `C:\Windows\System32`)

Verify installation in PowerShell:
```powershell
pastectl.exe --help
```

## Configuration

PasteCTL CLI stores configuration in `~/.config/pastectl/config.yaml`.

### Initial Setup

Configure the backend API URL:

```bash
pastectl config set backend_url https://api.paste.sumedh.app
```

Configure the frontend URL (for generating shareable links):

```bash
pastectl config set frontend_url https://paste.sumedh.app/paste
```

### Default Configuration

The CLI comes with the following defaults:
- **backend_url**: `https://api.paste.sumedh.app`
- **frontend_url**: `https://paste.sumedh.app/paste`

### View Configuration

Configuration is automatically loaded on each command. To modify settings, use the `config set` command.

## Usage

### Create a Paste

#### From Your Editor

Open your system's default editor to create a paste:

```bash
pastectl create
```

The CLI will use the editor specified in your `$EDITOR` environment variable (defaults to `vim` on Unix/Linux/macOS, `notepad` on Windows).

#### From a File

Create a paste from an existing file:

```bash
pastectl create --file /path/to/file.go
```

Language is automatically detected from the file extension.

#### With Custom Options

Specify language and expiry time:

```bash
pastectl create --file script.py --language python --expire 1h
```

Expiry options:
- `10m` - 10 minutes
- `1h` - 1 hour
- `24h` - 24 hours
- `never` - No expiration (default)

### Retrieve a Paste

#### Formatted Output

Get paste with metadata:

```bash
pastectl get <paste-id>
```

Output example:
```
--- Paste Details ---
ID:       abc123
Language: go
Created:  2025-10-03 10:21:05
--- Content ---
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

#### Raw Content

Get only the paste content (useful for piping):

```bash
pastectl get <paste-id> --raw
```

Example usage with output redirection:
```bash
pastectl get abc123 --raw > downloaded.go
```

### Update a Paste

Edit an existing paste in your editor:

```bash
pastectl update <paste-id>
```

The CLI will:
1. Fetch the current paste content
2. Open it in your editor
3. Update the paste with your changes

### Configuration Management

#### Set Configuration

```bash
pastectl config set <key> <value>
```

Available keys:
- `backend_url`: API endpoint URL
- `frontend_url`: Frontend base URL for shareable links

Example:
```bash
pastectl config set backend_url http://localhost:8080
```

## Supported Languages

The CLI automatically detects the following languages from file extensions:

| Extension | Language   |
|-----------|------------|
| `.js`     | javascript |
| `.py`     | python     |
| `.go`     | go         |
| `.java`   | java       |
| `.c`      | c          |
| `.cpp`    | cpp        |
| `.json`   | json       |
| `.md`     | markdown   |
| `.txt`    | plain      |

For other file types or editor-created pastes, you can manually specify the language using the `--language` flag.

## Editor Configuration

PasteCTL respects your system's `$EDITOR` environment variable. To set your preferred editor:

### Linux/macOS

Add to your shell configuration file (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export EDITOR=nano  # or vim, emacs, code, etc.
```

### Windows (PowerShell)

```powershell
$env:EDITOR = "notepad"
```

Common editor options:
- `vim` - Vim text editor
- `nano` - Nano text editor
- `emacs` - Emacs editor
- `code` - Visual Studio Code (use `code --wait`)
- `notepad` - Windows Notepad

## Examples

### Quick Code Sharing

Create and share a code snippet:

```bash
# Create from file
pastectl create --file main.go --expire 24h

# Output:
# Creating paste...
# Paste created successfully!
# Link: https://paste.sumedh.app/paste/abc123
```

### Downloading a Paste

Download and save a paste:

```bash
pastectl get abc123 --raw > downloaded.go
```

### Updating a Paste

Edit an existing paste:

```bash
pastectl update abc123
```

### Custom Backend

Use a self-hosted PasteCTL instance:

```bash
pastectl config set backend_url http://localhost:8080
pastectl config set frontend_url http://localhost:3000/paste
pastectl create --file script.sh
```

## API Integration

The CLI communicates with the PasteCTL backend via REST API:

- **POST** `/api/pastes` - Create paste
- **GET** `/api/pastes/:id` - Get paste with metadata
- **GET** `/api/pastes/:id/raw` - Get raw paste content
- **PUT** `/api/pastes/:id` - Update paste

## Project Structure

```
pastectl/
├── cmd/
│   ├── config.go      # Configuration management
│   ├── create.go      # Create paste command
│   ├── get.go         # Retrieve paste command
│   ├── update.go      # Update paste command
│   └── root.go        # Root command setup
├── internal/
│   ├── api/           # API client
│   ├── config/        # Configuration handling
│   └── editor/        # Editor integration
├── main.go
└── go.mod
```

## Dependencies

- **cobra**: CLI framework
- **viper**: Configuration management
- **Go 1.24.0+**: Required Go version

## Building from Source

Clone the repository and build:

```bash
git clone https://github.com/Sumedhvats/pastectl.git
cd pastectl
go build -o pastectl
```

Install locally:

```bash
go install
```

## Related Projects

- [PasteCTL Web](https://github.com/sumedhvats/pastectl_web) - Full-stack web application with Next.js frontend and Go backend

## Troubleshooting

### Backend URL Not Set

If you see the error:
```
backend_url is not set. Please use 'pastectl config set backend_url <url>'
```

Configure the backend URL:
```bash
pastectl config set backend_url https://api.paste.sumedh.app
```

### Editor Not Found

If your editor is not opening, set the `EDITOR` environment variable:

```bash
export EDITOR=vim  # or your preferred editor
```

### Permission Denied (Linux/macOS)

If you get a permission error when running the binary:

```bash
chmod +x pastectl
```

## Support

For issues, questions, or contributions, please visit the [GitHub Issues](https://github.com/Sumedhvats/pastectl/issues) page.

Made with ❤️ by Sumedh