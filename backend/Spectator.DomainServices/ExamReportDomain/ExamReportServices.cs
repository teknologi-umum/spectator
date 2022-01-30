using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Caching.Memory;
using Spectator.DomainModels.ExamReportDoman;
using Spectator.DomainServices.MemoryCache;
using Spectator.JwtAuthentication;

namespace Spectator.DomainServices.ExamReportDomain {
	public class ExamReportServices {
		private readonly JwtAuthenticationServices _jwtAuthenticationService;
		private readonly ResultCache<IReadOnlyDictionary<string, AdministratorUser>> _cache;
		private readonly ExamReportOptions _examReportOptions;
		public ExamReportServices(
			JwtAuthenticationServices jwtAuthenticationServices,
			IMemoryCache memoryCache,
			ExamReportOptions options
		) {
			_jwtAuthenticationService = jwtAuthenticationServices;
			ResultCache.Initialize(out _cache, memoryCache);
			_examReportOptions = options;
		}

		public string Login(string username, string password) {
			if (username == null || password == null) {
				throw new ArgumentNullException("username and/or password should not be empty");
			}

			if (username != _examReportOptions.Username && password != _examReportOptions.Password) {
				throw new Exception("username and/or password do not match");
			}

			// TODO: validate whether the current user is logged in or not

			var sessionId = Guid.NewGuid();
			var tokenPayload = _jwtAuthenticationService.CreatePayload(sessionId);
			var jwt = _jwtAuthenticationService.EncodeToken(tokenPayload);

			// TODO: create a dictionary of sessionId as key and AdministratorUser as value
			// then insert it into the MemoryCache.
			return jwt;
		}

		public string Logout(string jwt) {
			var decodedToken = _jwtAuthenticationService.DecodeToken(jwt);
			if (decodedToken.Expires < DateTime.UtcNow) {
				throw new Exception("Token already expired");
			}

			// TODO: remove the data from the MemoryCache.
			return "";
		}

		public async Task<List<File>> GetFilesAsync(CancellationToken cancellationToken) {
			// TODO: acquire the list of session id from InfluxDB directly
			// TODO: make a gRPC client call to the worker service to acquire the files list
		}
	}
}
