# PasteCTL CLI

A command-line interface for [PasteCTL](https://paste.sumedh.app), allowing developers to create, retrieve, and manage code snippets directly from the terminal.

## Overview

PasteCTL CLI is a terminal-based client for the PasteCTL platform. It provides a fast and efficient way to share code snippets without leaving your development environment. The CLI integrates with your system's default editor, supports automatic language detection from file extensions, and displays fetched pastes with syntax highlighting in your terminal.

## Features

- **Create Pastes** — from files or directly in your editor
- **Syntax Highlighted Output** — fetched pastes are displayed with monokai-themed ANSI colors
- **Automatic Language Detection** — detects language from file extensions
- **Custom Expiry Times** — set paste expiration: `1h` (default), `24h`, `7d`, or `never`
- **Retrieve Pastes** — formatted output with metadata or raw text for piping
- **Update Pastes** — edit existing pastes in your preferred editor
- **Configuration Management** — store API endpoint and frontend URL locally
- **Cross-Platform** — works on Linux, macOS, and Windows

## Installation

### Go Install

If you have Go installed:

```bash
go install github.com/Sumedhvats/pasteCTL@latest
```

Make sure your `$GOPATH/bin` is in your `PATH`.

### Pre-built Binaries

Download the appropriate binary for your platform from the [Releases page](https://github.com/Sumedhvats/pasteCTL/releases).

#### Linux

```bash
wget https://github.com/Sumedhvats/pasteCTL/releases/latest/download/pastectl_linux_amd64.tar.gz
tar -xvf pastectl_linux_amd64.tar.gz
sudo mv pastectl /usr/local/bin/
pastectl --version
```

#### macOS

```bash
wget https://github.com/Sumedhvats/pasteCTL/releases/latest/download/pastectl_darwin_amd64.tar.gz
tar -xvf pastectl_darwin_amd64.tar.gz
sudo mv pastectl /usr/local/bin/
pastectl --version
```

#### Windows

1. Download `pastectl_windows_amd64.zip` from the [Releases page](https://github.com/Sumedhvats/pasteCTL/releases)
2. Extract the archive
3. Move `pastectl.exe` to a directory in your `PATH`

```powershell
pastectl.exe --version
```

## Usage

### Create a Paste

**From your editor** (opens `$EDITOR`, defaults to `vim` / `notepad`):

```bash
pastectl create
```

**From a file** (language is auto-detected from the extension):

```bash
pastectl create --file main.go
```

**With custom options:**

```bash
pastectl create --file script.py --language python --expire 24h
```

| Flag | Short | Default | Description |
|---|---|---|---|
| `--file` | `-f` | — | Path to file |
| `--language` | `-l` | auto-detect | Override language detection |
| `--expire` | `-e` | `1h` | Expiry: `1h`, `24h`, `7d`, or `never` |

### Retrieve a Paste

**With syntax highlighting and metadata:**

```bash
pastectl get <paste-id>
```

Example output:
```
--- Paste Details ---
ID:       abc123
Language: go
Created:  2025-10-03 10:21:05
--- Content ---
package main              ← (highlighted in terminal)

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Raw content** (for piping or saving to a file):

```bash
pastectl get abc123 --raw > downloaded.go
```

**Without colors** (plain text with metadata):

```bash
pastectl get abc123 --no-color
```

| Flag | Default | Description |
|---|---|---|
| `--raw` | `false` | Output raw content only |
| `--no-color` | `false` | Disable syntax highlighting |

### Update a Paste

Edit an existing paste in your editor:

```bash
pastectl update <paste-id>
```

The CLI will:
1. Fetch the current paste content
2. Open it in your editor
3. Update the paste with your changes

## Supported Languages

The CLI automatically detects the following languages from file extensions:

| Extensions | Language |
|---|---|
| `.js` `.jsx` `.ts` `.tsx` `.mjs` `.cjs` | JavaScript |
| `.py` `.pyw` | Python |
| `.java` | Java |
| `.cpp` `.cc` `.cxx` `.hpp` `.hxx` `.h` | C++ |
| `.c` | C |
| `.go` | Go |
| `.sql` | SQL |
| `.txt` `.md` `.json` | Plain Text |

For other file types, use the `--language` flag: `pastectl create -f query.txt -l sql`

## Configuration

PasteCTL CLI stores configuration in `~/.config/pastectl/config.yaml`.

### Setup

```bash
# Set the backend API URL
pastectl config set backend_url https://api.paste.sumedh.app

# Set the frontend URL (used for generating shareable links)
pastectl config set frontend_url https://paste.sumedh.app/paste
```

### Defaults

The CLI comes with sensible defaults that work out of the box:

| Key | Default Value |
|---|---|
| `backend_url` | `https://api.paste.sumedh.app` |
| `frontend_url` | `https://paste.sumedh.app/paste` |

## Editor Configuration

PasteCTL respects your system's `$EDITOR` environment variable. To set your preferred editor:

**Linux / macOS** — add to `~/.bashrc` or `~/.zshrc`:

```bash
export EDITOR=nano  # or vim, emacs, code --wait
```

**Windows (PowerShell):**

```powershell
$env:EDITOR = "notepad"
```

## Examples

### Quick Code Sharing

```bash
# Create from file
pastectl create --file main.go --expire 24h

# Output:
# Creating paste...
# Paste created successfully!
# Link: https://paste.sumedh.app/paste/abc123
```

### Download a Paste

```bash
pastectl get abc123 --raw > downloaded.go
```

### Self-Hosted Instance

```bash
pastectl config set backend_url http://localhost:8080
pastectl config set frontend_url http://localhost:3000/paste
pastectl create --file script.sh
```

## API Endpoints

The CLI communicates with the PasteCTL backend via REST API:

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/api/pastes` | Create a paste |
| `GET` | `/api/pastes/:id` | Get paste with metadata |
| `GET` | `/api/pastes/:id/raw` | Get raw paste content |
| `PUT` | `/api/pastes/:id` | Update a paste |

## Project Structure

```
pasteCTL/
├── cmd/
│   ├── root.go           # Root command setup
│   ├── create.go         # Create paste command
│   ├── get.go            # Retrieve paste command
│   ├── update.go         # Update paste command
│   ├── config.go         # Configuration management
│   └── version.go        # Version command
├── internal/
│   ├── api/              # HTTP API client
│   ├── config/           # Viper config handling
│   ├── editor/           # System editor integration
│   └── highlight/        # Syntax highlighting (chroma)
├── main.go
└── go.mod
```

## Dependencies

| Dependency | Purpose |
|---|---|
| [cobra](https://github.com/spf13/cobra) | CLI framework |
| [viper](https://github.com/spf13/viper) | Configuration management |
| [chroma](https://github.com/alecthomas/chroma) | Syntax highlighting |

## Building from Source

```bash
git clone https://github.com/Sumedhvats/pasteCTL.git
cd pasteCTL
go build -o pastectl .
```

Install directly:

```bash
go install
```

## Troubleshooting

**Backend URL not set:**
```
backend_url is not set. Please use 'pastectl config set backend_url <url>'
```
Fix: `pastectl config set backend_url https://api.paste.sumedh.app`

**Editor not opening:**
Set the `EDITOR` environment variable: `export EDITOR=vim`

**Permission denied (Linux/macOS):**
```bash
chmod +x pastectl
```

## Related Projects

- [PasteCTL Web](https://github.com/Sumedhvats/pasteCTL_web) — Full-stack web application with Next.js frontend and Go backend
- [paste.sumedh.app](https://paste.sumedh.app) — Live instance

## Support

For issues, questions, or contributions, visit the [GitHub Issues](https://github.com/Sumedhvats/pasteCTL/issues) page.

Made with love by Sumedh