using System;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Writes;
using Microsoft.AspNetCore.Mvc;

namespace Spectator.Controllers;

[ApiController]
[Route("[controller]")]
public class SamTestController : ControllerBase {
	private readonly InfluxDBService _service;

	public SamTestController(InfluxDBService service) {
		_service = service;
	}

	// Write dummy data to teknum1 bucket -> /SamTest/WriteTest
	[HttpGet("WriteTest")]
	public string WriteTest() {
		_service.Write(write => {
			write.WritePoint("teknum1",
			"teknum1",
			 PointData.Measurement("testmesurement")
						.Tag("plane", "test-plane")
						.Field("value", 44)
						.Timestamp(DateTime.UtcNow, WritePrecision.Ns)
			);
		});

		return "Write OK";
	}

	// TODO: Read SAM test -> /SamTest/ReadTest
	[HttpGet("ReadTest")]
	public string ReadTEst() {
		return "Write OK";
	}
}
