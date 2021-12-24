using InfluxDB.Client;
using Microsoft.AspNetCore.Mvc;

namespace Spectator.Events;

[ApiController]
[Route("/events/")]
public class EventHandlerController : Controller {
	private readonly InfluxDBClient _db;
	public EventHandlerController(InfluxDBClient db) {
		_db = db;
	}

	[HttpGet]
	public IActionResult Index() {
		return NotFound();
	}
}
