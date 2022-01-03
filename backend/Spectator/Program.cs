﻿using System;
using System.Collections.Generic;
using System.Reflection;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using RG.ProtobufConverters.Json;
using Spectator.DomainServices;
using Spectator.Hubs;
using Spectator.JwtAuthentication;
using Spectator.Observables;
using Spectator.Piston;
using Spectator.RepositoryDALs;

var builder = WebApplication.CreateBuilder(args);

// Add configuration providers
builder.Configuration
	.AddJsonFile("appsettings.json", optional: true, reloadOnChange: true)
	.AddJsonFile($"appsettings.{builder.Environment.EnvironmentName}.json", optional: true, reloadOnChange: true)
	.AddKeyPerFile("/run/secrets", optional: true, reloadOnChange: true)
	.AddEnvironmentVariables("ASPNETCORE_")
	.AddUserSecrets<SessionHub>(optional: true, reloadOnChange: true);

// Add services
builder.Services.Setup(services => {
	// Configure options
	services.Configure<InfluxDbOptions>(builder.Configuration.GetSection("InfluxDbOptions"));

	// Add application layers
	services.AddHttpClient();
	services.AddRepositoryDALs();
	services.AddPistonClient();
	services.AddDomainServices();
	services.AddObservables();

	// Add MVC
	services.AddControllers();

	// Add Swagger
	services.AddEndpointsApiExplorer();
	services.AddSwaggerGen(options => {
		options.SwaggerDoc("v1", new Microsoft.OpenApi.Models.OpenApiInfo { Title = "Spectator SignalR API v1", Version = "v1" });
		options.DocumentFilter<SignalRSwaggerGen.SignalRSwaggerGen>(new List<Assembly> { typeof(SessionHub).Assembly });
	});

	// Add SignalR
	services.AddSignalR().AddJsonProtocol(options => options.PayloadSerializerOptions = ProtobufJsonConverter.Options);

	// Add authentication & authorization
	services.AddJwtBearerAuthentication();
	services.AddJwtBearerAuthorization();
});

// Build app
var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseAuthorization();
app.UseStaticFiles();

app.MapControllers();

app.MapHub<SessionHub>("/session");

// Run app
app.Run();
