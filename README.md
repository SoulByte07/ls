# ls-go

A minimal, from-scratch implementation of `ls` written in Go — built for learning the language, not for production.

---

## What this is

This project is a self-guided exercise in Go fundamentals: modules, CLI parsing (`flag`), file I/O (`os`, `path/filepath`), formatting (`text/tabwriter`), sorting, recursion, and terminal interaction (colors, ANSI codes). Every feature is implemented incrementally so the code stays understandable.

---

## Quick start

```bash
# Initialize the module (already done: go.mod exists)
go mod tidy

# Build
go build -o ls-go .

# Run
./ls-go
./ls-go -l -a /tmp
./ls-go -R -lh
```

---

## Features (implemented incrementally)

| Flag | Description |
|------|-------------|
| `-l` | Long listing format (permissions, size, owner, time) |
| `-a` | Show hidden files (dotfiles) |
| `-h` | Show help text |
| `-S` | Sort by file size |
| `-t` | Sort by modification time |
| `-r` | Reverse sort order |
| `-R` | Recursive directory listing |
| `-lh` | Human-readable file sizes |

Colors are enabled automatically when stdout is a terminal.

---

## Project structure

```
ls-go/
├── docs/
│   └── PLAN.md          # Phase-by-phase learning roadmap
├── main.go              # CLI entry point, flag parsing
├── list.go              # Directory listing logic (future)
├── list_test.go         # Unit tests (future)
├── LICENSE              # MIT
└── README.md            # This file
```

---

## Learning path

See [`docs/PLAN.md`](docs/PLAN.md) for the 8-phase roadmap. Each phase introduces new Go concepts and ends with a verifiable milestone (compile, test pass, or manual observation). The recommendation is to implement one phase at a time and review the code before moving on.

---

## Requirements

- Go 1.21 or later
- Linux / macOS / Windows (cross-compilation tested)

---

## License

MIT — see [`LICENSE`](LICENSE).
