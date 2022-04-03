# RCE - Remote Code Execution Engine

This module is a code execution engine, providing endpoints to safely and securely execute
arbitrary code from the list of supoprted languages. Current supported languages are:
C, C++, Java, Javascript, PHP, and Python. See [packages](./packages/) directory for
details.

## Development

You will need:
- Node.js v16.14 (see [.nvmrc file](./.nvmrc))
- Docker
- Chunky storage (at least 2.5GB to build the Dockerfile)

## Running with Docker

```sh
docker build -t spectator/rce .

docker run -p 50051:50051 spectator/rce
```
