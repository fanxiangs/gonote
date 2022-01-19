# [how to wirte go code](https://golang.google.cn/doc/code)

## Introduction

introduces the [go tool](https://golang.google.cn/cmd/go/),teh standard way to fetch, build and install Go module,
packages and commands.

## first program

### creat a go mod file

```shell
mkdir hello
cd hello
go mod init example/user/hello
# go: creating new go.mod: module example/user/hello
```

### go install

Create a file named [hello.go](hello/hello.go).

```shell
go install example/user/hello
```

This command producing an executable binary. The install directory is controlled by GOPATH and
GOBIN [environment variables](https://golang.google.cn/cmd/go/#hdr-Environment_variables). If GOBIN is set, binaries are
install to that directory. If GOPATH is set, binaries are install to the bin subdirectory of the first directory in the
GOPATH list. Otherwise, binaries are install to the bin subdirectory of the default GOPATH($HOME/go)
You can use the go env command to portably set the default value.

```shell
go env -w GOBIN=/somewhere/else/bin
```

To unset a variable previously set by go env -w, use go env -u:

```shell
go env -u GOBIN
```

For convenience, go command accept paths relative to the working directory, and default to package in the current
working directory if no path given. So the following commands are all equivalent:

```shell
go install example/user/hello
go install .
go install
```

For added convenience, we'll add the install directory to our PATH to make running binaries easy:

```shell
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
hello
```