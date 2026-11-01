# AGENTS.md — ls (mini-ls learning repo)

## Critical context (would miss without this)

- Directory: `ls` (renamed from `ls-go`). Module in `go.mod`: `github.com/SoulByte07/ls`. Do not reference `github.com/user/ls-go`.
- User writes all code. Agent acts as tutor/mentor only. Do not rewrite `main.go` without user approval.
- `docs/PLAN.md` is the single source of truth for phases, milestones, and verification commands.
- `main.go` currently panics on zero args (`arguments[-1]` with empty slice). Any agent testing should supply an argument or note the crash.

## Verified repo state

- `go.mod`: module `github.com/SoulByte07/ls`, go `1.27.3`.
- `.gitignore`: ignores binary named `ls`. A compiled binary `ls` already exists in root (built artifact).
- No `.cursor/`, `.cursorrules`, `.github/copilot-instructions.md`, `.opencode.json`, or pre-commit config.
- No tests (`*_test.go` missing). No CI workflows. No `Makefile`.
- Single package (`main`). No sub-packages, no generated code, no migrations.

## Commands (exact, non-obvious)

- Build/run: `go build -o ls .` then `./ls <path>`. Binary named `ls` conflicts with system `ls`; use `./ls` explicitly.
- Verify phases: reference `docs/PLAN.md` checklist; do not invent phase numbers.
- Cross-compile check for Phase 8: `GOOS=windows GOARCH=amd64 go build .`

## Constraints

- No filler, no emojis. Straight technical output. Brief.
- Scope locked to target step; do not edit files outside it.
- If docs (`README.md`, `PLAN.md`) conflict with `main.go` or `go.mod`, trust executable source (`main.go` / `go.mod`).
