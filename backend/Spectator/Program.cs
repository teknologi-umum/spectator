using System;
using System.Linq;
using System.Threading.Tasks;
using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Writes;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddSingleton<InfluxDBService>();
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
app.MapGet("/sam-test", (InfluxDBService service) => {
	service.Write(write => {
		write.WritePoint("teknum1",
		"teknum1",
		 PointData.Measurement("testmesurement")
					.Tag("plane", "test-plane")
					.Field("value", 55)
					.Timestamp(DateTime.UtcNow, WritePrecision.Ns)
		);
	});
});

// Get JWT through authorization header, parse it
app.MapGet("/user", (HttpContext r) => {
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

public class InfluxDBService {
	private readonly string _token;

	public InfluxDBService(IConfiguration configuration) {
		// _token = configuration.GetValue<string>("InfluxDB:Token");
		_token = "NxlAfGxz20SufSzLDGHnZRYJrUA_-l4b8HoXHZrkU5U41AS_lSyDvJuuzi6QAuN_r8XxcAfGp1kCWQ3Hl4qE1w==";
	}

	public void Write(Action<WriteApi> action) {
		using var client = InfluxDBClientFactory.Create("http://localhost:8086", _token);
		using var write = client.GetWriteApi();
		action(write);
	}

	public async Task<T> QueryAsync<T>(Func<QueryApi, Task<T>> action) {
		using var client = InfluxDBClientFactory.Create("http://localhost:8086", _token);
		var query = client.GetQueryApi();
		return await action(query);
	}
}
