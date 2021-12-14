FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["Spectator.csproj", "./"]
RUN dotnet restore "Spectator.csproj"
COPY . .
WORKDIR "/src/"
RUN dotnet build "Spectator.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Spectator.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Spectator.dll"]
