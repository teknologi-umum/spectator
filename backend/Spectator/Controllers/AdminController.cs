using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Spectator.DTO;


namespace Spectator.Controllers {
	public class AdminController : Controller {
		// POST /login
		[HttpPost]
		[Route("/login")]
		public IActionResult Login([FromBody] LoginRequest request) {
			// TODO
			return Ok();
		}

		// POST /logout
		[HttpPost]
		[Route("/logout")]
		public IActionResult Logout([FromBody] LogoutRequest request) {
			// TODO
			return Ok();
		}

		// GET /files
		[HttpGet]
		[Route("/files")]
		public async Task<IActionResult> FilesAsync([FromHeader] string Authentication) {
			// TODO: fetch session id list from influxdb direcly
			// TODO: send a grpc client request to the worker service to fetch ListFile data per each session id
			// TODO: profit. no, seriously, display the acquired results.
			return Ok();
		}
	}
}
