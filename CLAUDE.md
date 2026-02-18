# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Metsaman is a Go CLI application (`ikosaedri.fi/metsaman`) using Go 1.24 with no external dependencies. It uses Nix flakes with gomod2nix for reproducible builds and development environments.

## Build & Development Commands

```bash
# Build
nix build

# Run all checks (tests + linting)
nix flake check

# Run tests directly
go test -v ./...

# Run linting directly
golangci-lint run

# Enter development shell (provides go + gomod2nix)
nix develop

# Update Go module hashes for Nix after changing go.mod
gomod2nix
```

Direnv is configured (`.envrc`) to auto-load the flake dev shell.

## Architecture

Single-package Go application (`main` package) built via Nix flakes:

- **`flake.nix`** — Defines build package, dev shell, checks (go-test, go-lint), and a Nix overlay for importing `metsaman` as a package in other flakes
- **`default.nix`** — `buildGoApplication` definition (pname, version, module source)
- **`shell.nix`** — Dev shell with `mkGoEnv` and `gomod2nix`
- **`gomod2nix.toml`** — Nix-compatible Go dependency lockfile (currently empty, no deps)

## Conventions

- Commit messages use conventional commits style (`feat:`, `fix:`)
