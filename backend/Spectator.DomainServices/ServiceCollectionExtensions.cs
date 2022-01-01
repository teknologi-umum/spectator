﻿using Microsoft.Extensions.DependencyInjection;
using Spectator.DomainServices.PistonDomain;
using Spectator.DomainServices.SessionDomain;

namespace Spectator.DomainServices {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddDomainServices(this IServiceCollection services) {
			services.AddTransient<SessionServices>();
			services.AddTransient<SubmissionServices>();
			return services;
		}
	}
}
