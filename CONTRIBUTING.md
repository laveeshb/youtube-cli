# Contributing

## Development setup

Requires Go 1.21+.

```sh
git clone https://github.com/laveeshb/youtube-cli.git
cd youtube-cli
go mod download
make build
```

Binaries are output to `bin/yt` and `bin/youtube`.

## Project structure

```
cmd/
  yt/main.go          # entry point for 'yt' binary
  youtube/main.go     # entry point for 'youtube' binary
internal/
  auth/               # OAuth2 flow, token storage and refresh
  api/                # YouTube Data API and Analytics API wrappers
  config/             # config directory and file path resolution
pkg/
  root.go             # root Cobra command, shared by both binaries
  auth.go             # auth subcommands
  upload.go           # upload command
  publish.go          # publish command
  analytics.go        # analytics subcommands
  playlist.go         # playlist subcommands
  client.go           # shared API client helper
```

## Workflow

- Branch off `main` for all changes
- Open a PR — direct commits to `main` are not allowed
- PRs are merged via squash

## Running tests

```sh
make test
```

Integration tests (require real credentials) are skipped by default:

```sh
YOUTUBE_TEST_TOKEN=1 make test
```
