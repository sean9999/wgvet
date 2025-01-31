# WG Vet

Is a tool that does what `go vet` does, plus a few custom checks. It is useful for specific needs in a go dev shop. It's easy to extend.

## Getting Started

### Install

```sh
$ go install github.com/sean9999/wgvet@latest
```

### Run

```sh
$ cd /my/repo
$ wgvet . # or
$ wgvet ./...
```

## Github CI

You can have this run on every push with `.github/`:

```yaml
name: wgvet
run-name: vetting with wgvet
on: [push, workflow_dispatch]
jobs:
  vet:
    runs-on: ubuntu-latest
    steps:

      - name: Setup Go 1.23.x
        uses: actions/setup-go@v5
        with:
          go-version: '1.23.x'

      - name: checkout repo
        uses: actions/checkout@v4

      - name: insatll wgvet
        run: go install github.com/sean9999/wgvet@latest

      - name: vet
        run: wgvet ./...
```

### How to extend

```sh
$ git clone github.com/sean9999/wgvet
```

Edit or add exported symbols in `./pkg/*.go`

Import those in `./main.go`

```sh
$ make install
```

[![Go Reference](https://pkg.go.dev/badge/github.com/sean9999/wgvet.svg)](https://pkg.go.dev/github.com/sean9999/wgvet)