
using System;
using System.Linq;
using InfluxDB.Client;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddScoped<InfluxDBClient>(_ => {
	var token = System.Environment.GetEnvironmentVariable("INFLUX_TOKEN") ??
	            throw new ArgumentNullException($"INFLUX_TOKEN must not be empty");
	var dbUrl = System.Environment.GetEnvironmentVariable("INFLUX_URL") ??
				throw new ArgumentNullException($"INFLUX_URL must not be empty");
	return InfluxDBClientFactory.Create(dbUrl, token);
});
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}

// SAM test data
app.MapGet("/sam-test", () => "SAM test endpoint");

// Get JWT through authorization header, parse it
app.MapGet("/user", async (HttpContext r) => {
	try {
		// TODO: parse JWT to user information
		var jwToken = r.Request.Headers["authorization"].FirstOrDefault();

		return Results.Ok(jwToken);
	} catch (Exception) {
		return Results.Problem("Parsing JWT error");
	}
});



app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
