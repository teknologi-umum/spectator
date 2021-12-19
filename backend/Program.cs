using System;
using System.Linq;
using System.Text.Json;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

WebApplicationBuilder builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

WebApplication app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}

// SAM test data
app.MapGet("/sam-test", () => {
	return "SAM test endpoint";
});

// Get JWT through authorization header, parse it
app.MapGet("/user", async (HttpContext r) => {
	try {
		// TODO: parse JWT to user information
		var jwToken = r.Request.Headers["authorization"].FirstOrDefault();

		return Results.Ok(jwToken);
	} catch (Exception e) {
		return Results.Problem("Parsing JWT error");
	}
});



app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
