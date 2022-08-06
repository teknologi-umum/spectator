using System;
using System.Threading;
using System.Threading.Tasks;
using System.Security.Authentication;
using Microsoft.AspNetCore.Mvc;
using Spectator.DTO;
using Spectator.DomainServices.ExamReportDomain;

namespace Spectator.Controllers {
	[ApiController]
	public class AdminController : ControllerBase {
		private readonly ExamReportServices _examReportServices;

		public AdminController(
			ExamReportServices examReportServices
		) {
			_examReportServices = examReportServices;
		}

		[HttpPost]
		[Route("/admin/login")]
		public IActionResult Login([FromBody] LoginRequest request) {
			if (request.Password == null) return BadRequest(new { Message = "Password is required" });

			try {
				var session = _examReportServices.Login(request.Password);
				return Ok(session);
			} catch (AuthenticationException e) {
				return Unauthorized(new { Message = e.Message });
			} catch (ArgumentNullException e) {
				return BadRequest(new { Message = e.Message });
			}
		}

		[HttpPost]
		[Route("/admin/logout")]
		public IActionResult Logout([FromBody] LogoutRequest request) {
			if (request.SessionId == null) return BadRequest(new { Message = "SessionId is required" });

			if (!Guid.TryParse(request.SessionId, out var sessionId)) {
				return BadRequest(new { Message = "Invalid session id" });
			}

			_examReportServices.Logout(sessionId);
			return Ok();
		}

		[HttpPost]
		[Route("/admin/files")]
		public async Task<IActionResult> FilesAsync([FromBody] FilesRequest request, CancellationToken cancellationToken) {
			if (request.SessionId == null) return BadRequest(new { Message = "SessionId is required" });

			if (!Guid.TryParse(request.SessionId, out var sessionId)) {
				return BadRequest(new { Message = "Invalid session id" });
			}

			try {
				var files = await _examReportServices.GetFilesAsync(sessionId, cancellationToken);
				return Ok(files);
			} catch (UnauthorizedAccessException) {
				return Unauthorized();
			}
		}

		[HttpPost]
		[Route("/admin/retrigger")]
		public async Task<IActionResult> RetriggerAsync([FromBody] RetriggerRequest request, CancellationToken cancellationToken) {
			if (request.SessionId == null) return BadRequest(new { Message = "SessionId is required" });
			if (request.ExamSessionId == null) return BadRequest(new { Message = "ExamSessionId is required" });

			if (!Guid.TryParse(request.SessionId, out var adminSessionId)) {
				return BadRequest(new { Message = "Invalid admin session id" });
			}

			if (!Guid.TryParse(request.ExamSessionId, out var examSessionId)) {
				return BadRequest(new { Message = "Invalid exam session id" });
			}

			try {
				await _examReportServices.RetriggerResultAsync(adminSessionId, examSessionId, cancellationToken);
				return Ok();
			} catch (UnauthorizedAccessException) {
				return Unauthorized();
			}
		}
	}
}
