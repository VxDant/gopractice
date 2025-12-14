# gopractice

A small practice repository to learn Go fundamentals.

## About
This repo contains exercises and small projects for learning Go (Golang). It is intentionally unstructured today — Day 1 — and will evolve as you progress.

## Goals
- Learn Go syntax and core language features
- Practice data structures and algorithms in Go
- Build small projects to apply concepts
- Explore standard library and common packages
- Improve code organization and testing habits

## Getting started
1. Install Go: https://golang.org/doc/install  
2. Set up GOPATH / Go modules (recommended: use modules).  
3. Create a module at the repo root:
    ```
    go mod init github.com/yourusername/gopractice
    ```
4. Run programs:
    ```
    go run ./path/to/package
    ```
5. Build:
    ```
    go build ./...
    ```

## Suggested project layout
- cmd/        — entry points (one folder per small program)
- pkg/        — reusable packages
- internal/   — private packages
- examples/   — short example programs
- data/       — sample data files
- tests/      — integration tests or sandbox tests
- README.md

## Day 1 (current)
- Create this README
- Initialize module
- Add a simple "hello world" in `cmd/hello/main.go`
- Explore `fmt`, `errors`, and basic types

## Learning plan (short-term)
- Week 1: basics — types, control flow, functions, packages, modules
- Week 2: data structures — slices, maps, structs, interfaces
- Week 3: concurrency — goroutines, channels, sync primitives
- Week 4: testing and tooling — go test, benchmarking, vet, lint

## Resources
- The Go Programming Language: https://golang.org/doc/
- Effective Go: https://golang.org/doc/effective_go.html
- Go by Example: https://gobyexample.com

## Contributing
- Keep examples small and focused
- Include README or comments for each example
- Add tests where applicable

