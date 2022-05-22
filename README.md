# Spectator

Brief explanation of the project. TBC.

## Structure

| Codebase  | Description                                                                 |
| --------- | --------------------------------------------------------------------------- |
| frontend  | React SPA frontend                                                          |
| backend   | ASP.NET Core core backend API                                               |
| rce       | Node.js service for  code execution engine inspired by [Piston][piston-url] |
| logger    | Go service for any logging                                                  |
| worker    | Go service for post-data processing                                         |
| dummy     | Python sample data generation for InfluxDB                                  |
| assertion | Coding test assertion scripts                                               |
| proto     | Shared protocol buffers contracts                                           |

## Development setup

This development setup is for running the whole application on a single machine.
For running each component, see the respective README for each directory.

Prerequisites:
- Docker
- Docker Compose

To run just the InfluxDB:
```sh
docker-compose --file docker-compose.influx.yml up -d
```

To run everything on the development mode:
```sh
docker-compose --file docker-compose.dev.yml up -d
```

To run apps that is production ready:
```sh
docker-compose --file docker-compose.prod.yml up -d
```

To drop the container, simply use:
```sh
docker-compose --file <which file you ran before> down
```

Or if you just want to stop it:
```sh
docker-compose --file <which file you ran before> stop
```

## Code of Conduct

Please read [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md) for details on our code of conduct.

## License

TBD.

[piston-url]: https://github.com/engineer-man/piston