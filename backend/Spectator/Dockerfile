FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build

WORKDIR /node
RUN apt-get update && apt-get upgrade -y && apt-get install -y curl tar
RUN curl -LO https://nodejs.org/dist/v16.15.0/node-v16.15.0-linux-x64.tar.gz && \
    tar -zxf node-v16.15.0-linux-x64.tar.gz && \
    cd node-v16.15.0-linux-x64 && \
    mkdir --parents --verbose /usr/bin /usr/include /usr/lib /usr/share/doc /usr/share/man && \
    mv -v bin/* /usr/bin/ && \
    mv -v include/* /usr/include/ && \
    mv -v lib/* /usr/lib/ && \
    mv -v share/doc/* /usr/share/doc/ && \
    mv -v share/man/* /usr/share/man/  && \
    mv -v share/* /usr/share/

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
