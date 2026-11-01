# Project Name: mini-ls — A Learning Implementation Plan

> **Goal**: Build a functional `ls` clone in Go from scratch to learn Go fundamentals through hands-on coding.
> **Stack**: Go · **Runtime**: Go 1.21+ · **Module**: `github.com/user/ls-go`
>
> **Overall Progress**: `░░░░░░░░░░░░░░░░░░░░ 0% (0/8 Phases) (0hr/12hr Time)` : by effort

---

## Phase 1: Scaffolding & Hello CLI  `+5%` `0.5hr`

- [x] **1.1** Initialize Go module
  - Run `go mod init github.com/user/ls-go` in project root
  - File: `go.mod`
  - **Done when**: `go.mod` exists with correct module name

- [x] **1.2** Create `main.go` with minimal CLI entry point
  - File: `main.go`
  - Key details: `package main`, `func main()`, `fmt.Println("hello")`
  - **Done when**: `go run .` prints "hello"

- [x] **1.3** Verify build and run workflow
  - Run `go build -o ls-go .` and `./ls-go`
  - **Done when**: Binary builds and runs successfully

---

## Phase 2: Read Directory & List Entries  `+15%` `1hr`

- [x] **2.1** Import `os` and use `os.ReadDir(".")`
  - File: `main.go`
  - Key details: Handle error returned by `ReadDir`, iterate `[]DirEntry`, print `entry.Name()`
  - **Done when**: Running `./ls-go` lists files in current directory (one per line)

- [!] **2.2** Accept directory path as positional argument
  - Use `os.Args[1:]` to get target directory (default to `"."`)
  - File: `main.go`
  - **Done when**: `./ls-go /some/path` lists that directory

- [ ] **2.3** Basic error handling for non-existent paths
  - Check `os.IsNotExist(err)` and print user-friendly message
  - **Done when**: Invalid path shows clean error, exits non-zero

---

## Phase 3: CLI Flags (`-l`, `-a`, `-h`)  `+15%` `1.5hr`

- [ ] **3.1** Define flag variables using `flag` package
  - `flag.BoolVar(&showLong, "l", false, "long format")`
  - `flag.BoolVar(&showAll, "a", false, "show hidden files")`
  - `flag.BoolVar(&showHelp, "h", false, "help")`
  - File: `main.go`
  - **Done when**: Flags parse without panic

- [ ] **3.2** Implement `-h` / `--help` usage output
  - Print flag descriptions and exit
  - **Done when**: `./ls-go -h` shows help text

- [ ] **3.3** Filter hidden files unless `-a` is set
  - Skip entries where `strings.HasPrefix(name, ".")` unless `showAll`
  - **Done when**: `./ls-go` hides dotfiles; `./ls-go -a` shows them

- [ ] **3.4** Parse flags before positional args (`flag.Parse()`)
  - Remaining `flag.Args()` are directory paths
  - **Done when**: `./ls-go -l /tmp` works correctly

---

## Phase 4: Long Format (`-l`) — File Metadata  `+25%` `2.5hr`

- [ ] **4.1** For each entry, call `entry.Info()` to get `fs.FileInfo`
  - File: `main.go` (or new `list.go` for separation)
  - Key details: `FileInfo` provides `Mode()`, `Size()`, `ModTime()`, `Name()`, `IsDir()`
  - **Done when**: Can access metadata for each file

- [ ] **4.2** Format file mode as `drwxr-xr-x` string
  - Use `Mode().String()` or custom formatting for `rwx` triplets
  - **Done when**: Output shows permission string like `ls -l`

- [ ] **4.3** Get owner/group names via `user` package
  - `user.LookupId(strconv.Itoa(int(stat.Sys().(*syscall.Stat_t).Uid)))`
  - Requires `import "os/user"` and `import "syscall"`
  - **Done when**: Owner and group names appear (or UID/GID fallback)

- [ ] **4.4** Format modification time as `Mon Jan 2 15:04`
  - `modTime.Format("Jan 2 15:04")`
  - **Done when**: Timestamp matches `ls -l` style

- [ ] **4.5** Assemble long-format line: `mode links owner group size time name`
  - Use `fmt.Printf` or `tabwriter` for alignment
  - **Done when**: `./ls-go -l` output visually matches `ls -l`

---

## Phase 5: Pretty Columns & Alignment  `+15%` `1.5hr`

- [ ] **5.1** Import `text/tabwriter` for aligned output
  - Initialize `tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)`
  - **Done when**: Columns align regardless of filename length

- [ ] **5.2** Implement columnar output for default (non-`-l`) mode
  - Calculate terminal width or use fixed columns
  - Print names separated by `\t`, flush writer
  - **Done when**: `./ls-go` shows multi-column output

- [ ] **5.3** Handle long filenames gracefully (truncate or wrap)
  - **Done when**: No broken layout on wide names

---

## Phase 6: Sorting Options  `+10%` `1hr`

- [ ] **6.1** Collect all entries into a slice before printing
  - File: `list.go` — `type Entry struct { Name string; Info fs.FileInfo }`
  - **Done when**: Can sort slice in memory

- [ ] **6.2** Implement sort-by-name (default), sort-by-size (`-S`), sort-by-time (`-t`)
  - Use `sort.Slice(entries, func(i, j int) bool { ... })`
  - **Done when**: `./ls-go -l -S` sorts by size descending

- [ ] **6.3** Add `-r` flag for reverse order
  - **Done when**: `./ls-go -lr` reverses name sort

---

## Phase 7: Recursive Listing (`-R`)  `+10%` `1.5hr`

- [ ] **7.1** Add `-R` / `--recursive` flag
  - File: `main.go`
  - **Done when**: Flag parses

- [ ] **7.2** Implement recursive walk using `filepath.WalkDir`
  - For each directory, print header (`path:`), then entries
  - Recurse into subdirectories
  - **Done when**: `./ls-go -R` shows tree-like output

- [ ] **7.3** Respect `-a` and `-l` in recursive mode
  - **Done when**: `./ls-go -R -l -a` works correctly

---

## Phase 8: Polish & Quality  `+5%` `1hr`

- [ ] **8.1** Add color output for file types (directories, executables, symlinks)
  - Use ANSI escape codes; disable if `!isatty`
  - **Done when**: `./ls-go` shows colors in terminal

- [ ] **8.2** Human-readable sizes (`-h` flag for sizes: 1K, 234M, 2G)
  - Implement `formatSize(bytes int64) string`
  - **Done when**: `./ls-go -lh` shows `4.0K` instead of `4096`

- [ ] **8.3** Add basic unit tests for core functions
  - File: `list_test.go` — test sorting, formatting, filtering
  - **Done when**: `go test ./...` passes

- [ ] **8.4** Cross-compile check: `GOOS=windows GOARCH=amd64 go build`
  - **Done when**: Build succeeds for other platforms

---

## Stretch Goals (Post-MVP)

- [ ] **S1** JSON output (`--json`) for scripting
- [ ] **S2** Git status indicators (like `exa` / `lsd`)
- [ ] **S3** Tree view (`--tree`) with ASCII connectors
- [ ] **S4** Configuration file (`~/.config/ls-go/config.toml`)
- [ ] **S5** Benchmark and optimize for large directories (100k+ files)

---

## Verification Checklist

```bash
# Phase 1
go mod init github.com/user/ls-go
go run .                          # → "hello"
go build -o ls-go . && ./ls-go    # → "hello"

# Phase 2
./ls-go                           # → lists current dir
./ls-go /tmp                      # → lists /tmp
./ls-go /nonexistent              # → clean error

# Phase 3
./ls-go -h                        # → help text
./ls-go -a                        # → shows dotfiles
./ls-go -l                        # → long format (later phases)

# Phase 4
./ls-go -l                        # → permissions, owner, size, time, name

# Phase 5
./ls-go                           # → aligned columns

# Phase 6
./ls-go -l -S                     # → sorted by size
./ls-go -l -t                     # → sorted by time
./ls-go -lr                       # → reverse name sort

# Phase 7
./ls-go -R                        # → recursive
./ls-go -R -l                     # → recursive long

# Phase 8
./ls-go                           # → colored output
./ls-go -lh                       # → human sizes
go test ./...                     # → all tests pass
GOOS=windows go build .           # → cross-compiles
```

---

## Learning Milestones (for Mentor Tracking)

| Milestone | Go Concepts Mastered |
|-----------|---------------------|
| Phase 1 | Module system, `main`, build/run cycle |
| Phase 2 | `os` package, `DirEntry`, error handling, `os.Args` |
| Phase 3 | `flag` package, boolean flags, positional args |
| Phase 4 | `fs.FileInfo`, `syscall`, `user` package, time formatting |
| Phase 5 | `tabwriter`, string formatting, stdout control |
| Phase 6 | `sort.Slice`, custom comparators, closures |
| Phase 7 | Recursion, `filepath.WalkDir`, tree traversal |
| Phase 8 | ANSI codes, terminal detection, testing, cross-compilation |

---

**Next Step**: Start **Phase 1.1** — run `go mod init github.com/user/ls-go` and create `main.go`. I'll review your code after each task.
