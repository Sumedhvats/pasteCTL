# pasteCTL

A command-line interface for pasteCTL - share code snippets directly from your terminal.

[![Go Report Card](https://goreportcard.com/badge/github.com/Sumedhvats/pasteCTL_cli)](https://goreportcard.com/report/github.com/Sumedhvats/pasteCTL_cli)
[![Go Reference](https://pkg.go.dev/badge/github.com/Sumedhvats/pasteCTL_cli.svg)](https://pkg.go.dev/github.com/Sumedhvats/pasteCTL_cli)
[![Release](https://img.shields.io/github/v/release/Sumedhvats/pasteCTL_cli)](https://github.com/Sumedhvats/pasteCTL_cli/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

pastectl is a command-line client for the pasteCTL platform that enables developers to share code snippets and text directly from their terminal. Built with Go and designed for seamless integration into development workflows.

## Features

- File-based paste creation from any file in your filesystem
- Interactive editor integration for creating pastes
- Automatic language detection based on file extensions
- Configurable expiration times for paste lifecycle management
- Raw output support for scripting and automation
- Cross-platform compatibility (Linux, macOS, Windows)

## Installation

### Using Go

```bash
go install github.com/Sumedhvats/pastectl@latest
```

### Binary Releases

Download pre-compiled binaries from [GitHub Releases](https://github.com/Sumedhvats/pasteCTL_cli/releases).

### From Source

```bash
git clone https://github.com/Sumedhvats/pasteCTL_cli.git
cd pasteCTL_cli
go build -o pastectl ./cmd/pastectl
```

## Configuration

Configure the CLI to connect to your pasteCTL instance:

```bash
pastectl config set frontend_url https://paste.example.com
pastectl config set backend_url https://api.paste.example.com
```

## Usage

### Creating Pastes

Create a paste from a file:
```bash
pastectl create --file main.go
```

Create a paste using your default editor:
```bash
pastectl create
```

Specify language and expiration:
```bash
pastectl create --file script.py --language python --expire 24h
```

### Retrieving Pastes

Get paste content with metadata:
```bash
pastectl get <paste-id>
```

Get raw content only:
```bash
pastectl get <paste-id> --raw
```

### Updating Pastes

Update an existing paste:
```bash
pastectl update <paste-id>
```

## Command Reference

### pastectl create

Create a new paste from a file or using an editor.

```
pastectl create [flags]
```

**Flags:**
- `-f, --file string`: Create paste from file path
- `-l, --language string`: Override automatic language detection
- `-e, --expire string`: Set expiration (default: never)

**Examples:**
```bash
pastectl create -f ./main.go
pastectl create -f config.json -l json -e 1h
```

### pastectl get

Retrieve a paste by ID.

```
pastectl get <id> [flags]
```

**Flags:**
- `--raw`: Output only the paste content

**Examples:**
```bash
pastectl get abc123def
pastectl get abc123def --raw > local-copy.py
```

### pastectl update

Update an existing paste using an editor.

```
pastectl update <id>
```

### pastectl config

Manage CLI configuration.

```
pastectl config set <key> <value>
```

## Supported Languages

The CLI automatically detects language from file extensions:

| Extension | Language |
|-----------|----------|
| `.go` | go |
| `.py` | python |
| `.js` | javascript |
| `.java` | java |
| `.c` | c |
| `.cpp` | cpp |
| `.json` | json |
| `.md` | markdown |

## Expiration Options

Set paste expiration using these formats:
- `10m` - 10 minutes
- `1h` - 1 hour  
- `24h` - 24 hours
- `7d` - 7 days
- `never` - No expiration (default)

## Integration Examples

### CI/CD Pipeline

```bash
# Share build logs
make build 2>&1 | tee build.log
pastectl create -f build.log --expire 7d
```

### Development Workflow

```bash
# Share current diff
git diff HEAD~1 | pastectl create -l diff --expire 1h
```

### System Administration

```bash
# Share system information
pastectl create -f /var/log/syslog --expire 24h
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `EDITOR` | Default text editor | `vi` |
| `PASTECTL_CONFIG` | Config file path | `~/.pastectl.yaml` |

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