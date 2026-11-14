# sysinfo

Simple tool to collect system information: CPU count, hostname, OS and
architecture, temp dir, Go version, current user, and non-loopback network
interfaces with their addresses.

## Install

```sh
go install github.com/mikemadden42/sysinfo@latest
```

Or download a prebuilt binary from the
[releases page](https://github.com/mikemadden42/sysinfo/releases).

## Usage

```sh
sysinfo          # human-readable output
sysinfo -json    # JSON output
```

Example:

```
CPUs: 8
Hostname: myhost
OS: linux
OS arch: amd64
Temp dir: /tmp
Go version: go1.25.0
User: alice
UID: 1000
GID: 1000
eth0 00:11:22:33:44:55
eth0 -> 192.168.1.10/24
```

## Development

```sh
make check    # golangci-lint
make build    # snapshot build via goreleaser
go test ./...
```
