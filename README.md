# snake-go
### An implementation of the game snake written in Golang utilizing Ebitengine.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue)](https://golang.org)
[![Ebiten Version](https://img.shields.io/badge/ebiten-%3E%3Dv2.8-green)](https://ebiten.org/)
[![License](https://img.shields.io/badge/license-MIT-lightgrey)](/LICENSE)

![snake-go-demo](https://github.com/user-attachments/assets/da30ecd6-d955-4b55-a6ac-1036c252c192)

### Features:
- [x] WASM Compilation
- [x] Customizable Settings in `settings.yaml`
- [x] Hardware Acceleration

## Prerequisites

Before building make sure you have **Go** installed. 
Install from [https://go.dev/dl](https://go.dev/dl).

To verify that Go is installed:
```bash
go version
```

## Getting the Code

Clone the repo and enter it's directory:
```bash
git clone https://github.com/somewhat9/snake-go.git
cd snake-go
```

## Build

Fetch dependencies:
```bash
go mod tidy
```

Compile:
```bash
go build -o bin/snake-go ./cmd/snake-go
```

## Run

Run directly (without binary):
```bash
go run ./cmd/snake-go/main.go
```

Run from the built binary:
```bash
./bin/snake-go
```

