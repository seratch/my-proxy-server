# my-proxy-server

This tiny Go app is a simple CLI tool to run your own HTTP proxy server.

## Build the CLI

`go build` for running the command on your local machine:

```bash
brew update && brew upgrade go
go build
```

To build cross-platform binary files, you can use `goreleaser`:

```bash
brew install goreleaser
rm -rf dist/ && goreleaser build --snapshot
tree dist/
```

## Run the CLI

By default, the command starts `http://localhost:9000/` without basic auth.

```bash
$ ./my-proxy-server -h
Run a simple HTTP proxy server

Usage:
  my-proxy-server [flags]

Flags:
  -a, --auth              Enable proxy auth (Basic Auth)
  -h, --help              help for my-proxy-server
  -p, --password string   Basic auth password (default: pass) (default "pass")
  -P, --port string       HTTP proxy port (default: 9000) (default "9000")
  -u, --username string   Basic auth username (default: user) (default "user")
```

You can use the above flags as necessary:

```bash
./my-proxy-server -P 8080 -a -u user -p pass
```

## License

The MIT License