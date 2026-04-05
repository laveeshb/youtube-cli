# youtube-cli

A cross-platform command-line tool to manage your YouTube channel from the terminal. Works as both `yt` and `youtube`.

## Installation

### From source

Requires Go 1.21+.

```sh
git clone https://github.com/laveeshb/youtube-cli.git
cd youtube-cli
make install
```

This installs both `yt` and `youtube` binaries to your `$GOPATH/bin`.

## Setup

Before using youtube-cli, you need to create your own Google Cloud credentials. This tool does not ship with shared credentials — each user brings their own.

### 1. Create a Google Cloud project

1. Go to [console.cloud.google.com](https://console.cloud.google.com)
2. Create a new project
3. Enable the following APIs:
   - YouTube Data API v3
   - YouTube Analytics API

### 2. Create OAuth credentials

1. Go to **APIs & Services → Credentials**
2. Configure the **OAuth consent screen** first (External, Testing mode)
3. Add your Google account as a test user under the **Audience** tab
4. Back in Credentials, click **Create Credentials → OAuth client ID**
5. Application type: **Desktop app**
6. Download the JSON file

### 3. Place credentials

Move the downloaded JSON to:

| Platform | Path |
|----------|------|
| macOS | `~/Library/Application Support/youtube-cli/credentials.json` |
| Linux | `~/.config/youtube-cli/credentials.json` |
| Windows | `%AppData%\youtube-cli\credentials.json` |

### 4. Authenticate

```sh
yt auth login
```

This opens a browser for Google OAuth. Once complete, your token is stored locally and auto-refreshed.

## Usage

### Auth

```sh
yt auth login     # authenticate with Google
yt auth status    # show login state and token expiry
yt auth logout    # revoke and remove credentials
```

### Upload

```sh
yt upload video.mp4 \
  --title "My video" \
  --description "Description here" \
  --tags cooking,food \
  --thumbnail thumb.jpg \
  --privacy private

# Schedule for future publishing
yt upload video.mp4 --title "My video" --schedule 2026-05-01T10:00:00Z
```

### Publish

```sh
yt publish <video-id>                          # publish immediately
yt publish <video-id> --schedule 2026-05-01T10:00:00Z  # schedule
```

### Analytics

```sh
yt analytics channel              # channel stats (default: last 28 days)
yt analytics channel --period 90d
yt analytics video <video-id>
```

### Playlists

```sh
yt playlist list
yt playlist add <playlist-id> <video-id>
```

## Notes

- `yt` and `youtube` are identical — use whichever you prefer
- Scheduled videos must use `--privacy private` (YouTube requirement)
- The app stays in Google's "Testing" mode — only you can authorize it
