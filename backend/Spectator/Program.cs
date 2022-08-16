using System;
using System.Threading.Tasks;
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
using Spectator.VideoClient;
using Spectator;

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
	services.Configure<VideoOptions>(builder.Configuration.GetSection("VideoOptions"));
	services.Configure<ExamReportOptions>(builder.Configuration.GetSection("ExamReportOptions"));
	services.Configure<MinioOptions>(builder.Configuration.GetSection("MinioOptions"));

	// Add application layers
	services.AddHttpClient();
	services.AddMemoryCache();
	services.AddRepositoryDALs();
	services.AddPistonClient();
	services.AddLoggerClient();
	services.AddWorkerClient();
	services.AddVideoClient();
	services.AddDomainServices();
	services.AddObservables();

	// Add MVC
	services.AddControllersWithViews();

	// Add Swagger
	services.AddEndpointsApiExplorer();
	services.AddSwaggerGen(options => {
		options.SwaggerDoc("v1", new Microsoft.OpenApi.Models.OpenApiInfo { Title = "Spectator SignalR API v1", Version = "v1" });
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
		builder.SetIsOriginAllowed(origin => true).AllowCredentials().AllowAnyMethod().AllowAnyHeader();
	}));

	// Add SPA static files
	services.AddSpaStaticFiles(options => {
		options.RootPath = "wwwroot";
	});
});

// Build app
var app = builder.Build();

await app.Services.MigrateDatabaseAsync();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}
// else {
// 	app.UseHsts();
// }

// // Redirect HTTP traffic on production to HTTPS
// if (app.Environment.IsProduction()) {
// 	app.UseHttpsRedirection();
// }

// Middlewares
app.UseRouting();
app.UseAuthentication();
app.UseAuthorization();
app.UseCors("AllowAll");

// Map Frontend static files
app.UseDefaultFiles();
app.UseStaticFiles();
app.UseSpaStaticFiles();

// Map Controllers
app.UseEndpoints(endpoints => {
	endpoints.MapControllerRoute(
		name: "default",
		pattern: "{controller}/{action=Index}/{id?}"
	);
});

// Host SPA application
app.UseSpa(spa => {
	spa.Options.SourcePath = "../../frontend";

	if (app.Environment.IsDevelopment()) {
		// use proxy instead of react development server since we're not using create-react-app
		// this requires running `npm run dev` manually from the frontend folder
		// see: https://github.com/dotnet/aspnetcore/issues/33466#issuecomment-859487783
		spa.UseProxyToSpaDevelopmentServer("http://localhost:3000");
	}
});

// Map SignalR Hubs
app.MapHub<SessionHub>("/hubs/session");
app.MapHub<EventHub>("/hubs/event");

// Run app
await app.RunAsync();

namespace Spectator {
	public static class Application {
		public static Task MigrateDatabaseAsync(this IServiceProvider services) {
			return services.GetRequiredService<InfluxDbInitializer>().InitializeAsync();
		}
	}
}
