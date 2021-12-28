# Piston

This is a fork of [Piston](https://github.com/engineer-man/piston) with a very specific modifications
tailored to meet the needs of the Spectator project. Including auto installation for some of the
language that Spectator needs for its' coding test.

For the general purpose of Piston, please visit the original repository.

## Available runtimes

`GET /api/v2/runtimes`

Find available and installed runtimes that are supported to execute code.

## Execute code

`POST /api/v2/execute`

Request body schema:
```json5
{
    "language": "string (required)",
    "version": "string (required)",
    "files": [
        {
            "name": "string (required)", // file name
            "content": "string (required)" // code
        }
    ],
    "compile_timeout": 10000, // optional, defaults to 10 seconds
    "run_timeout": 3000, // optional, defaults to 3 seconds
    "compile_memory_limit": -1, // optional, defaults to no limit
    "run_memory_limit": -1, // optional, defaults to no limit
}
```