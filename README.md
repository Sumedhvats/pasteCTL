# pasteCTL

A CLI tool to create, view, and manage code pastes from your terminal.

```bash
# Share a file instantly
pastectl create -f main.go

# View a paste with syntax highlighting
pastectl get abc12
```

## Install

```bash
go install github.com/Sumedhvats/pasteCTL@latest
```

Or download a binary from [Releases](https://github.com/Sumedhvats/pasteCTL/releases).

## Quick Start

pasteCTL works out of the box with the hosted service at [paste.sumedh.app](https://paste.sumedh.app).

```bash
# Create a paste from a file (language auto-detected)
pastectl create -f script.py

# Create a paste in your editor
pastectl create

# View a paste (syntax highlighted)
pastectl get abc12

# Get raw content (useful for piping)
pastectl get abc12 --raw > downloaded.py

# Edit an existing paste
pastectl update abc12
```

## Commands

| Command | Description |
|---|---|
| `pastectl create` | Create a paste from a file (`-f`) or your editor |
| `pastectl get <id>` | View a paste with syntax-highlighted output |
| `pastectl update <id>` | Edit an existing paste in your editor |
| `pastectl config set <key> <value>` | Set configuration (backend/frontend URL) |
| `pastectl version` | Print CLI version |

## Create Options

```bash
pastectl create [flags]
```

| Flag | Short | Default | Description |
|---|---|---|---|
| `--file` | `-f` | | Path to file |
| `--language` | `-l` | auto | Override language detection |
| `--expire` | `-e` | `1h` | Expiry: `1h`, `24h`, `7d`, or `never` |

## Get Options

```bash
pastectl get <id> [flags]
```

| Flag | Default | Description |
|---|---|---|
| `--raw` | `false` | Output raw content only (for piping) |
| `--no-color` | `false` | Disable syntax highlighting |

## Language Detection

When using `--file`, the language is auto-detected from the file extension:

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

Use `--language` to override: `pastectl create -f query.txt -l sql`

## Configuration

Config is stored in `~/.config/pastectl/config.yaml`.

```bash
# Point to a self-hosted backend
pastectl config set backend_url http://localhost:8080
pastectl config set frontend_url http://localhost:3000/paste
```

**Defaults:**
- `backend_url`: `https://api.paste.sumedh.app`
- `frontend_url`: `https://paste.sumedh.app/paste`

## Editor

pasteCTL uses your `$EDITOR` environment variable (defaults to `vim` on Unix, `notepad` on Windows).

```bash
export EDITOR=nano  # or vim, code --wait, etc.
```

## Build from Source

```bash
git clone https://github.com/Sumedhvats/pasteCTL.git
cd pasteCTL
go build -o pastectl .
```

## Related

- [pasteCTL Web](https://github.com/Sumedhvats/pasteCTL_web) — Full-stack web app (Next.js + Go)
- [paste.sumedh.app](https://paste.sumedh.app) — Live instance