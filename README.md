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

youtube-cli does not ship with shared Google credentials. Each user creates their own Google Cloud app — this keeps your channel data private and under your control. The setup takes about 10 minutes and is a one-time process.

### Step 1 — Create a Google Cloud project

1. Go to [console.cloud.google.com](https://console.cloud.google.com) and sign in with the Google account that owns your YouTube channel
2. Click the project dropdown at the top → **New Project**
3. Give it any name (e.g. `youtube-cli`), click **Create**
4. Make sure the new project is selected in the dropdown

### Step 2 — Enable the YouTube APIs

1. In the left sidebar, go to **APIs & Services → Library**
2. Search for **YouTube Data API v3**, click it, then click **Enable**
3. Go back to the Library, search for **YouTube Analytics API**, click it, then click **Enable**

### Step 3 — Configure the OAuth consent screen

This is the screen users see when authorizing the app. Since this is a personal tool, you only need minimal configuration.

1. Go to **APIs & Services → OAuth consent screen**
2. Select **External**, click **Create**
3. Fill in the required fields:
   - App name: anything (e.g. `youtube-cli`)
   - User support email: your email
   - Developer contact email: your email
4. Click **Save and Continue** through the Scopes step (no changes needed)
5. On the Test Users step, click **Add users** and add your Google account email
6. Click **Save and Continue**, then **Back to Dashboard**

> The app stays in Testing mode permanently. This means only the Google accounts you add as test users can authorize it — nobody else can.

### Step 4 — Create OAuth credentials

1. Go to **APIs & Services → Credentials**
2. Click **Create Credentials → OAuth client ID**
3. Application type: **Desktop app**
4. Name: anything (e.g. `youtube-cli`)
5. Click **Create**
6. In the dialog that appears, click **Download JSON**

### Step 5 — Place the credentials file

Move the downloaded JSON file to the youtube-cli config directory:

| Platform | Path |
|----------|------|
| macOS | `~/Library/Application Support/youtube-cli/credentials.json` |
| Linux | `~/.config/youtube-cli/credentials.json` |
| Windows | `%AppData%\youtube-cli\credentials.json` |

On macOS/Linux you can run:

```sh
# macOS
mkdir -p ~/Library/Application\ Support/youtube-cli
mv ~/Downloads/client_secret_*.json ~/Library/Application\ Support/youtube-cli/credentials.json

# Linux
mkdir -p ~/.config/youtube-cli
mv ~/Downloads/client_secret_*.json ~/.config/youtube-cli/credentials.json
```

### Step 6 — Authenticate

```sh
yt auth login
```

This opens a browser window. Sign in with the same Google account that owns your YouTube channel. If you manage multiple channels (e.g. a Brand Account), you will be prompted to choose which channel to authorize during the sign-in flow.

Once complete, your token is stored locally and automatically refreshed — you only need to log in once.

Verify it worked:

```sh
yt auth status
```

## Usage

### Auth

```sh
yt auth login     # authenticate with Google
yt auth status    # show login state and token expiry
yt auth logout    # revoke and remove credentials
```

### Upload

Upload a video file with metadata:

```sh
yt upload video.mp4 \
  --title "My video title" \
  --description "Full description here" \
  --tags cooking,food,recipe \
  --thumbnail thumbnail.jpg \
  --privacy private
```

Upload and schedule for future publishing:

```sh
yt upload video.mp4 \
  --title "My video" \
  --schedule 2026-05-01T10:00:00Z
```

> Scheduled videos are automatically set to private until the publish time. The `--schedule` flag uses RFC3339 format (UTC).

Available flags:

| Flag | Default | Description |
|------|---------|-------------|
| `--title` | | Video title |
| `--description` | | Video description |
| `--tags` | | Comma-separated tags |
| `--thumbnail` | | Path to thumbnail image (JPG/PNG) |
| `--privacy` | `private` | `public`, `private`, or `unlisted` |
| `--schedule` | | Publish at this time (RFC3339) |

### Publish

Publish a video that is currently private or scheduled:

```sh
yt publish <video-id>                                    # publish immediately
yt publish <video-id> --schedule 2026-05-01T10:00:00Z   # reschedule
```

The video ID can be found in the YouTube URL: `youtube.com/watch?v=<video-id>`

### Analytics

View channel-level stats:

```sh
yt analytics channel               # last 28 days (default)
yt analytics channel --period 7d
yt analytics channel --period 90d
```

View stats for a specific video:

```sh
yt analytics video <video-id>
yt analytics video <video-id> --period 90d
```

Period format is `<N>d` where N is the number of days.

### Playlists

```sh
yt playlist list                              # list all playlists
yt playlist add <playlist-id> <video-id>     # add a video to a playlist
```

Playlist IDs are visible in the YouTube Studio URL when viewing a playlist.

## Multiple channels

If you manage more than one YouTube channel (e.g. a personal channel and a Brand Account), run `yt auth logout` and then `yt auth login` again to switch channels. During the Google sign-in flow, YouTube will prompt you to select which channel to authorize.

## Notes

- `yt` and `youtube` are identical — use whichever you prefer
- All credentials are stored locally on your machine and never shared
- The Google app stays in Testing mode — only you can authorize it
- youtube-cli does not keep a copy of your video files — the originals stay on your machine
