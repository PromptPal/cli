# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

PromptPal CLI is a Go-based command-line tool that generates type definitions for prompt management. It connects to a PromptPal server and generates strongly-typed code for Go and TypeScript applications.

## Essential Commands

### Development
```bash
# Run tests with coverage
go test -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run a specific test
go test -v ./commands/internal -run TestGeneratorName

# Build the CLI
go build -o promptpal main.go

# Build with version info
go build -ldflags "-s -w -X github.com/PromptPal/cli/main.GitCommit=v1.0.0" -o promptpal main.go
```

### Release Process
```bash
# Local release build (requires GITHUB_TOKEN and Apple tokens for macOS)
goreleaser --clean
```

## Architecture Overview

### Core Components

1. **CLI Framework**: Built on `urfave/cli/v2`, entry point in `main.go`
2. **Commands** (`commands/`):
   - `init.go`: Creates `promptpal.yml` configuration file
   - `gen.go`: Orchestrates code generation from server
3. **Code Generators** (`commands/internal/`):
   - `generator.go`: Base generator interface and shared logic
   - `generator_go.go`: Go code generation using templates
   - `generator_typescript.go`: TypeScript code generation
   - Template files (`*.tpl`) for code generation

### Key Design Patterns

1. **Configuration-Driven**: All operations center around `promptpal.yml`:
   ```yaml
   input:
       http:
           url: "server-url"
           token: "token-or-@env.VARIABLE"
   output:
       go_types:
           prefix: PP
           output: "./types.g.go"
           package_name: "pkg"
   ```

2. **Template-Based Generation**: Uses Go templates to generate code:
   - `go.tpl`: Template for Go constants and structs
   - `typescript.tpl`: Template for TypeScript enums and types

3. **Server Communication**: HTTP client (`resty`) fetches prompt schemas from PromptPal server

### Testing Strategy

Tests use `testify/suite` for organization. Each generator has comprehensive tests validating:
- Template rendering
- Type generation correctness
- Edge cases and error handling

When adding new features, follow the existing test patterns in `*_test.go` files.
