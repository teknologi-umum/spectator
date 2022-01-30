using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Spectator.DomainModels.ExamReportDoman {
	public record AdministratorUser(
		string JwtToken,
		DateTime ExpiredAt,
		DateTime CreatedAt
		) {
		public static AdministratorUser Apply(
			string jwtToken,
			DateTime expiredAt,
			DateTime createdAt
		) {
			return new(
				JwtToken: jwtToken,
				ExpiredAt: expiredAt,
				CreatedAt: createdAt);
		}
	}
}
