# Worker

This module is a worker, providing function and endpoints to aggregate and calculate data, as such:
- Generating a fun fact about a user after they had finished the coding test.
- Generating a few files for a certain user containing their information (keystrokes, mouse activities, personal
  information, etc.).

## Structure

| Package | Description |
| ------- | ----------- |
| common  | Provides constants and enum for measurement names and common types. |
| file    | Provides data aggregation and file generation functionalities. |
| funfact | Provides quick calculation for generating WPM, deletion rate, and test attempts and store it as a projection. |
| logger | Provides a helper function to communicate with the logger gRPC server. |
| logger_porto | Contains protocol buffers stub for the logger service. |
| worker_proto | Contains protocol buffers stub for the worker service. |

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
go -v -cover -race ./...

# or for specific packages
go -v -cover worker/funfact
go -v -cover worker/file
```

For production purposes:

```
go build .
./worker

# or if you are on Windows
worker.exe
```