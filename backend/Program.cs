using System;
using InfluxDB.Client;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

var builder = WebApplication.CreateBuilder(args);

// Connect to the InfluxDB instance


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
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
