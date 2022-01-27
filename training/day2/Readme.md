
# Troubleshooting

## 1. goroot not found

<strong>GOPATH</strong> is discussed in the `cmd/go` documentation:

The GOPATH environment variable lists places to look for Go code. On Unix, the value is a colon-separated string. On Windows, the value is a semicolon-separated string. On Plan 9, the value is a list.

GOPATH must be set to get, build and install packages outside the standard Go tree.

GOROOT is discussed in the installation instructions:

The Go binary distributions assume they will be installed in /usr/local/go (or c:\Go under Windows), but it is possible to install the Go tools to a different location. In this case you must set the GOROOT environment variable to point to the directory in which it was installed.

For example, if you installed Go to your home directory you should add the following commands to $HOME/.profile:

```bash
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
```
Note: GOROOT must be set only when installing to a custom location.

## 2. go: go.mod file not found in current directory or any parent directory; see 'go help modules'

First make sure that your GO111MODULE value is set to "auto". You can check it from:

```bash
go env
```
if it is not set to "auto", then run:

```bash
go env -w GO111MODULE=auto
```

go to your work directory in terminal and run:

```bash
go mod init main
go mod tidy
```