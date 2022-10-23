# Spectator backend

The backend is event driven, domain driven, and object oriented.
It hosts SignalR endpoints to interact with frontend and logs every event to InfluxDB.

## Structure

| Project                     | Description                                                              |
| --------------------------- | ------------------------------------------------------------------------ |
| Spectator                   | ASP.NET Core web server                                                  |
| Spectator.Primitives        | Contains enums and primitive types to be shared across all .NET projects |
| Spectator.DomainModels      | Contains domain models which represent current state of entities         |
| Spectator.DomainEvents      | Contains domain events which represent state mutation of entities        |
| Spectator.DomainServices    | Contains business logic of each domain                                   |
| Spectator.Observables       | Reactive in-memory stores for storing state of entities                  |
| Spectator.Repositories      | Contains repository interfaces                                           |
| Spectator.RepositoryDALs    | Contains repository DAL implementations                                  |
| Spectator.JwtAuthentication | Contains authentication and authorization logic                          |
| Spectator.Piston            | Contains piston client implementation                                    |

## Visual Studio development setup

- Get [Visual Studio 2022 Community](https://visualstudio.microsoft.com/)
- Install `ASP.NET and web development` workload
- Open `Spectator.sln` in the root directory
- Set up user secrets by right clicking on `Solution 'Spectator'`/`backend`/`Spectator` then select `Manage User Secrets`

```json
{
  "PistonOptions:BaseUrl": "https://your-piston-url.com/",
  "InfluxDbOptions:Url": "http://your-influxdb-url.com/",
  "InfluxDbOptions:Token": "yourinfluxdbtoken"
}
```

## Jetbrains Rider development setup

- Install [Jetbrains Rider](https://www.jetbrains.com/rider/)
- Open `Spectator.sln` in the root directory
- Go to `File` > `Settings...` > `Plugins`, then find and install `.NET Core User Secrets`
- Set up user secrets by right clicking on `Solution 'Spectator'`/`backend`/`Spectator` then select `Tools` > `Initialize User Secrets`

```json
{
  "PistonOptions:BaseUrl": "https://your-piston-url.com/",
  "InfluxDbOptions:Url": "http://your-influxdb-url.com/",
  "InfluxDbOptions:Token": "yourinfluxdbtoken"
}
```

## Visual Studio Code development setup

- Assuming you've already had Visual Studio Code installed
- Get [.NET 6 SDK](https://dotnet.microsoft.com/en-us/download/visual-studio-sdks)
- Open the root folder, then open `Program.cs` and let Visual Studio Code download all the required extensions
- Set up user secrets:

```
cd backend/Spectator
dotnet user-secrets set "PistonOptions:BaseUrl" "https://your-piston-url.com/"
dotnet user-secrets set "InfluxDbOptions:Url" "http://your-influxdb-url.com/"
dotnet user-secrets set "InfluxDbOptions:Token" "yourinfluxdbtoken"
```