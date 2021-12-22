using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Writes;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace Spectator.Controllers;

[ApiController]
[Route("[controller]")]
public class UserController : ControllerBase {
	InfluxDBService _service;
	HttpContext _context;

	public UserController(InfluxDBService service) {
		_service = service;
	}
	// Get user info through JWT -> User/Info
	// Supply header: Authorization: eyJhfd8sa7fudsahfkiasdhfd89a==
	[HttpGet("Info")]
	public dynamic UserInfo() {
		var authHeader = HttpContext.Request.Headers["authorization"].FirstOrDefault();

		// TODO: actually parse jwt

		return $"[Auth Header] {authHeader}";
	}
}
