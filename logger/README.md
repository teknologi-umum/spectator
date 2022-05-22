# Logger

Logger provides standard logging capabilities to log data into the InfluxDB database storage.


## Visual Studio Code development setup

- Install [Visual Studio Code](https://code.visualstudio.com/Download)
- Install [Go version 1.17.6](https://go.dev/dl/) or newer
- Install [Go extension for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- Open this worker directory on your Visual Studio Code, don't open the whole Spectator monolith project.
  Go server on Visual Studio Code doesn't do well with multiple modules being opened at the same time.
- Set up secrets on `.env` file using `.env.example` as a template. Or use the default value provided
  on the `main.go` file.

## Vim development setup

- Install [Vim](https://www.vim.org/download.php) or [Neovim](https://github.com/neovim/neovim/wiki/Installing-Neovim)
- Install [Go version 1.17.6](https://go.dev/dl/) or newer
- Install Go language server
  - For Vim, install [vim-go](https://github.com/fatih/vim-go)
  - For Neovim, install [lspconfig](https://github.com/neovim/nvim-lspconfig)
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
go run test -v -cover -covermode=atomic -race -timeout=120s
```

## Required Environment Variables

* `INFLUX_URL` - InfluxDB URL
* `INFLUX_TOKEN` - InfluxDB token to enable read/write access to the database
* `INFLUX_ORG` - InfluxDB organization name
* `ACCESS_TOKEN` - Accesss token to send request to the logger service.
* `PORT` - Application TCP port to listen to