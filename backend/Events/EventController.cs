using System;
using System.Threading;
using System.Threading.Tasks;
using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Writes;
using Microsoft.AspNetCore.Mvc;

namespace Spectator.Events; 

[ApiController]
public class EventHandler : Controller {
	private readonly InfluxDBClient _db;
	public EventHandler(InfluxDBClient db) {
		_db = db;
	}

	public async Task AddAsync(CancellationToken cancellationToken, Event e) {
		var builder = PointData.Measurement("event")
			.Tag("type", e.Type.ToString())
			.Tag("user", e.User)
			.Field("value", e.Value)
			.Timestamp(e.Date, WritePrecision.Ns);
		var writeApi = _db.GetWriteApiAsync();
		await writeApi.WritePointAsync(builder, cancellationToken);
	}
	
	// GET
	public IActionResult Index() {
		return NotFound();
	}
}
