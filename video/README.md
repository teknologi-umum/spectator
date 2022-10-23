# Video

This module is a background worker that converts chunk of WEBM/VP9 video stream into a playable
webm (for backups) and MP4 (with H264 codec) file format.

## Structure

| Package      | Description                                            |
| ------------ | ------------------------------------------------------ |
| samples      | Sample video files for unit testing purposes           |
| logger_porto | Contains protocol buffers stub for the logger service. |
| worker_proto | Contains protocol buffers stub for the worker service. |

## Visual Studio Code development setup

- Install [Visual Studio Code](https://code.visualstudio.com/Download)
- Install [Go version 1.18.5](https://go.dev/dl/) or newer
- Install [Go extension for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- Install Ffmpeg
- Open this worker directory on your Visual Studio Code, don't open the whole Spectator monolith project.
  Go server on Visual Studio Code doesn't do well with multiple modules being opened at the same time.
- Set up secrets on `.env` file using `.env.example` as a template. Or use the default value provided
  on the `main.go` file.

## Vim development setup

- Install [Vim](https://www.vim.org/download.php) or [Neovim](https://github.com/neovim/neovim/wiki/Installing-Neovim)
- Install [Go version 1.18.5](https://go.dev/dl/) or newer
- Install Go language server
  - For Vim, install [vim-go](https://github.com/fatih/vim-go)
  - For Neovim, install [lspconfig](https://github.com/neovim/nvim-lspconfig)
- Install Ffmpeg
- Open the worker directory.
- Set up secrets on `.env` file using `.env.example` as a template. Or use the default value provided
  on the `main.go` file.

## Running the application

For development purposes:

```
go run .
```

For testing purposes:

```
go -v -cover -race -timeout=180s ./...
```

For production purposes:

```
go build .
./video

# or if you are on Windows
go build -o video.exe .
video.exe
```

## Required Environment Variables

* `INFLUX_TOKEN` - InfluxDB token to enable read/write access to the database
* `INFLUX_HOST` - InfluxDB URL
* `INFLUX_ORG` - InfluxDB organization name
* `MINIO_HOST` - Minio hostname, without "http://" prefix
* `MINIO_ACCESS_ID` - Minio access ID
* `MINIO_SECRET_KEY` - Minio secret key
* `MINIO_TOKEN` - Minio token (defaults to empty)
* `LOGGER_SERVER_ADDRESS` - Logger service address
* `LOGGER_TOKEN` - Token for the logger service
* `ENVIRONMENT` - Service runtime environment
* `PORT` - Application TCP port to listen to
