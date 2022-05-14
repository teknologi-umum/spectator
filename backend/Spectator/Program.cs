﻿using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Reflection;
using System.Text.Json;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using RG.ProtobufConverters.Json;
using Spectator.DomainServices;
using Spectator.DomainServices.ExamReportDomain;
using Spectator.Hubs;
using Spectator.JwtAuthentication;
using Spectator.LoggerClient;
using Spectator.Observables;
using Spectator.Piston;
using Spectator.PoormansAuth;
using Spectator.RepositoryDALs;
using Spectator.WorkerClient;

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
	services.Configure<PistonOptions>(builder.Configuration.GetSection("PistonOptions"));
	services.Configure<LoggerOptions>(builder.Configuration.GetSection("LoggerOptions"));
	services.Configure<WorkerOptions>(builder.Configuration.GetSection("WorkerOptions"));
	services.Configure<ExamReportOptions>(builder.Configuration.GetSection("ExamReportOptions"));

	// Add application layers 
	services.AddHttpClient();
	services.AddMemoryCache();
	services.AddRepositoryDALs();
	services.AddPistonClient();
	services.AddLoggerClient();
	services.AddWorkerClient();
	services.AddDomainServices();
	services.AddObservables();

	// Add MVC
	services.AddControllersWithViews();

	// Add Swagger
	services.AddEndpointsApiExplorer();
	services.AddSwaggerGen(options => {
		options.SwaggerDoc("v1", new Microsoft.OpenApi.Models.OpenApiInfo { Title = "Spectator SignalR API v1", Version = "v1" });
		options.DocumentFilter<SignalRSwaggerGen.SignalRSwaggerGen>(new List<Assembly> { typeof(SessionHub).Assembly });
	});

	// Add SignalR
	services.AddSignalR(hubOptions => {
		hubOptions.EnableDetailedErrors = true;
	}).AddJsonProtocol(options => {
		options.PayloadSerializerOptions = ProtobufJsonConverter.Options;
		options.PayloadSerializerOptions.PropertyNamingPolicy = JsonNamingPolicy.CamelCase;
	});

	// Add authentication & authorization
	services.AddJwtBearerAuthentication();
	services.AddJwtBearerAuthorization();
	services.AddPoormansAuth();

	// Add Cors Policy
	services.AddCors(options => options.AddPolicy("AllowAll", builder => {
		// TODO(elianiva): replace this with proper CORS policy, ATM this is being used to make it *just works*
		builder.WithOrigins("http://localhost:3000").AllowAnyMethod().AllowAnyHeader().AllowCredentials();
	}));
});

// Build app
var app = builder.Build();

await app.Services.MigrateDatabaseAsync();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
} else {
	app.UseHsts();
}

// Redirect HTTP traffic
// app.UseHttpsRedirection();

// Middlewares
app.UseRouting();
app.UseAuthentication();
app.UseAuthorization();
app.UseCors("AllowAll");

// Map Frontend static files
app.UseDefaultFiles();
app.UseStaticFiles();

// Map Controllers
app.UseEndpoints(endpoints => {
	endpoints.MapControllerRoute(
		name: "default",
		pattern: "{controller}/{action=Index}/{id?}"
	);
});

// Map SignalR Hubs
app.MapHub<SessionHub>("/hubs/session");
app.MapHub<EventHub>("/hubs/event");

// Run app
app.Run();

public static class Application {
	public static Task MigrateDatabaseAsync(this IServiceProvider services) {
		return services.GetRequiredService<InfluxDbInitializer>().InitializeAsync();
	}
}
