using System.IdentityModel.Tokens.Jwt;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Spectator.JwtAuthentication.RequirementHandlers;
using Spectator.JwtAuthentication.Requirements;

namespace Spectator.JwtAuthentication {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddJwtBearerAuthentication(this IServiceCollection services) {
			services.AddSingleton<JwtSecurityTokenHandler>();
			services.AddSingleton<JwtAuthenticationSettings>();
			services.AddSingleton<IPostConfigureOptions<JwtBearerOptions>, PostConfigureJwtBearerOptions>();
			services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
				.AddJwtBearer();
			services.AddTransient<JwtAuthenticationServices>();
			return services;
		}

		public static IServiceCollection AddJwtBearerAuthorization(this IServiceCollection services) {
			services.AddTransient<IAuthorizationHandler, RegisteredRequirementHandler>();
			services.AddTransient<IAuthorizationHandler, SAMRequirementHandler>();
			services.AddTransient<IAuthorizationHandler, ExamRequirementHandler>();
			services.AddAuthorization(options => {
				options.DefaultPolicy = new AuthorizationPolicyBuilder()
					.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
					.RequireAuthenticatedUser()
					.Build();
				options.AddPolicy(
					name: AuthPolicy.ANONYMOUS,
					configurePolicy: builder => builder
						.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
						.RequireAuthenticatedUser()
						.AddRequirements(new RegisteredRequirement(isRegistered: false))
						.Build()
				);
				options.AddPolicy(
					name: AuthPolicy.REGISTERED,
					configurePolicy: builder => builder
						.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
						.RequireAuthenticatedUser()
						.AddRequirements(new RegisteredRequirement(isRegistered: true))
						.AddRequirements(new SAMRequirement(beforeTest: false))
						.Build()
				);
				options.AddPolicy(
					name: AuthPolicy.READY_TO_TAKE_EXAM,
					configurePolicy: builder => builder
						.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
						.RequireAuthenticatedUser()
						.AddRequirements(new SAMRequirement(beforeTest: true, afterTest: false))
						.AddRequirements(new ExamRequirement(started: false))
						.Build()
				);
				options.AddPolicy(
					name: AuthPolicy.TAKING_EXAM,
					configurePolicy: builder => builder
						.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
						.RequireAuthenticatedUser()
						.AddRequirements(new ExamRequirement(started: true, ended: false))
						.Build()
				);
				options.AddPolicy(
					name: AuthPolicy.HAS_TAKEN_EXAM,
					configurePolicy: builder => builder
						.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
						.RequireAuthenticatedUser()
						.AddRequirements(new ExamRequirement(ended: true))
						.AddRequirements(new SAMRequirement(afterTest: false))
						.Build()
				);
				options.AddPolicy(
					name: AuthPolicy.COMPLETED,
					configurePolicy: builder => builder
						.AddAuthenticationSchemes(JwtBearerDefaults.AuthenticationScheme)
						.RequireAuthenticatedUser()
						.AddRequirements(new SAMRequirement(afterTest: true))
						.Build()
				);
			});
			return services;
		}
	}
}
