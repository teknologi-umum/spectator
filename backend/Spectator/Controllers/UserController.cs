using System.Linq;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace Spectator.Controllers;

[ApiController]
[Route("[controller]")]
public class UserController : ControllerBase {
	private readonly InfluxDBService _service;

	public UserController(InfluxDBService service) {
		_service = service;
	}
	// Get user info through JWT -> User/Info
	// Supply header: Authorization: eyJhfd8sa7fudsahfkiasdhfd89a==
	[HttpGet("Info")]
	public string UserInfo() {
		var authHeader = HttpContext.Request.Headers["authorization"].FirstOrDefault();

		// TODO: actually parse jwt

		return $"[Auth Header] {authHeader}";
	}
}
