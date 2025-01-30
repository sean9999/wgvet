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
$ wgvet . 
```

### How to extend

```sh
$ git clone github.com/sean9999/wgvet
```

Edit or add files in `./pkg/*.go`

Import those in `./main.go`

```sh
$ make install
```

