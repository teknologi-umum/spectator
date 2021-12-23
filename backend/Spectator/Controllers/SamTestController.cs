using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Writes;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace Spectator.Controllers;

[ApiController]
[Route("[controller]")]
public class SamTestController : ControllerBase {
	InfluxDBService _service;

	public SamTestController(InfluxDBService service) {
		_service = service;
	}

    // Write dummy data to teknum1 bucket -> /SamTest/WriteTest
	[HttpGet("WriteTest")]
	public dynamic WriteTest() {
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
	public dynamic ReadTEst() {
        return "Write OK";
	}

    
}
