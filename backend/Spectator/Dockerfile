FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["backend/Spectator/Spectator.csproj", "backend/Spectator/"]
COPY ["backend/Spectator.DomainEvents/Spectator.DomainEvents.csproj", "backend/Spectator.DomainEvents/"]
COPY ["backend/Spectator.DomainModels/Spectator.DomainModels.csproj", "backend/Spectator.DomainModels/"]
COPY ["backend/Spectator.DomainServices/Spectator.DomainServices.csproj", "backend/Spectator.DomainServices/"]
COPY ["backend/Spectator.JwtAuthentication/Spectator.JwtAuthentication.csproj", "backend/Spectator.JwtAuthentication/"]
COPY ["backend/Spectator.Observables/Spectator.Observables.csproj", "backend/Spectator.Observables/"]
COPY ["backend/Spectator.Piston/Spectator.Piston.csproj", "backend/Spectator.Piston/"]
COPY ["backend/Spectator.Primitives/Spectator.Primitives.csproj", "backend/Spectator.Primitives/"]
COPY ["backend/Spectator.Repositories/Spectator.Repositories.csproj", "backend/Spectator.Repositories/"]
COPY ["backend/Spectator.RepositoryDALs/Spectator.RepositoryDALs.csproj", "backend/Spectator.RepositoryDALs/"]
COPY ["proto/Spectator.Protos/Spectator.Protos.csproj", "proto/Spectator.Protos/"]
COPY ["frontend/Frontend.csproj", "frontend/"]
RUN dotnet restore "backend/Spectator/Spectator.csproj"
COPY . .
WORKDIR "/src/backend/Spectator"
RUN dotnet build "Spectator.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Spectator.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Spectator.dll"]
