#if DEBUG
using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading;
using System.Threading.Tasks;
using FluentAssertions;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Spectator.DomainEvents.SessionDomain;
using Spectator.Primitives;
using Spectator.Repositories;
using Xunit;

namespace Spectator.RepositoryDALs.Tests {
	public class SessionEventRepositoryDALTests {
		private IServiceProvider ServiceProvider { get; }

		public SessionEventRepositoryDALTests() {
			var configuration = new ConfigurationBuilder()
				.AddKeyPerFile("/run/secrets", optional: true)
				.AddEnvironmentVariables("ASPNETCORE_")
				.AddUserSecrets<SessionEventRepositoryDALTests>(optional: true)
				.Build();
			var services = new ServiceCollection();
			services.AddSingleton<IConfiguration>(configuration);
			services.Configure<InfluxDbOptions>(configuration.GetSection("InfluxDbOptions"));
			services.AddHttpClient();
			services.AddRepositoryDALs();
			ServiceProvider = services.BuildServiceProvider();
		}

		[Fact]
		public async Task CanLogSessionStartedAsync() {
			var repository = ServiceProvider.GetRequiredService<ISessionEventRepository>();

			// Create event
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var newEvent = new SessionStartedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				Locale: Locale.EN
			);

			// Add event to log
			await repository.AddEventAsync(newEvent);

			// Get logged events
			var loggedEvents = new List<SessionEventBase>();
			await foreach (var loggedEvent in repository.GetAllEventsAsync(sessionId, CancellationToken.None)) {
				loggedEvents.Add(loggedEvent);
			}

			// Validate logged events
			loggedEvents.Count.Should().Be(1);
			loggedEvents[0].Should().BeOfType<SessionStartedEvent>();
			var sessionStartedEvent = (SessionStartedEvent)loggedEvents[0];
			sessionStartedEvent.SessionId.Should().Be(sessionId);
			sessionStartedEvent.Timestamp.Should().Be(timestamp);
			sessionStartedEvent.Locale.Should().Be(Locale.EN);
		}

		[Fact]
		public async Task CanLogPersonalInfoSubmittedAsync() {
			var repository = ServiceProvider.GetRequiredService<ISessionEventRepository>();

			// Create event
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var newEvent = new PersonalInfoSubmittedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				StudentNumber: "22536257326",
				YearsOfExperience: 25,
				HoursOfPractice: 12,
				FamiliarLanguages: "Java, English, Tagalog, Brainfuck",
				WalletNumber: "123456789",
				WalletType: "grabpay"
			);

			// Add event to log
			await repository.AddEventAsync(newEvent);

			// Get logged events
			var loggedEvents = new List<SessionEventBase>();
			await foreach (var loggedEvent in repository.GetAllEventsAsync(sessionId, CancellationToken.None)) {
				loggedEvents.Add(loggedEvent);
			}

			// Validate logged events
			loggedEvents.Count.Should().Be(1);
			loggedEvents[0].Should().BeOfType<PersonalInfoSubmittedEvent>();
			var personalInfoSubmittedEvent = (PersonalInfoSubmittedEvent)loggedEvents[0];
			personalInfoSubmittedEvent.SessionId.Should().Be(sessionId);
			personalInfoSubmittedEvent.Timestamp.Should().Be(timestamp);
			personalInfoSubmittedEvent.StudentNumber.Should().Be("22536257326");
			personalInfoSubmittedEvent.YearsOfExperience.Should().Be(25);
			personalInfoSubmittedEvent.HoursOfPractice.Should().Be(12);
			personalInfoSubmittedEvent.FamiliarLanguages.Should().Be("Java, English, Tagalog, Brainfuck");
			personalInfoSubmittedEvent.WalletNumber.Should().Be("123456789");
			personalInfoSubmittedEvent.WalletType.Should().Be("grabpay");
		}

		[Fact]
		public async Task CanLogBeforeExamSAMSubmittedAsync() {
			var repository = ServiceProvider.GetRequiredService<ISessionEventRepository>();

			// Create event
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var newEvent = new BeforeExamSAMSubmittedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				SelfAssessmentManikin: new SelfAssessmentManikin(3, 5)
			);

			// Add event to log
			await repository.AddEventAsync(newEvent);

			// Get logged events
			var loggedEvents = new List<SessionEventBase>();
			await foreach (var loggedEvent in repository.GetAllEventsAsync(sessionId, CancellationToken.None)) {
				loggedEvents.Add(loggedEvent);
			}

			// Validate logged events
			loggedEvents.Count.Should().Be(1);
			loggedEvents[0].Should().BeOfType<BeforeExamSAMSubmittedEvent>();
			var samSubmittedEvent = (BeforeExamSAMSubmittedEvent)loggedEvents[0];
			samSubmittedEvent.SessionId.Should().Be(sessionId);
			samSubmittedEvent.Timestamp.Should().Be(timestamp);
			samSubmittedEvent.SelfAssessmentManikin.ArousedLevel.Should().Be(3);
			samSubmittedEvent.SelfAssessmentManikin.PleasedLevel.Should().Be(5);
		}

		[Fact]
		public async Task CanLogExamStartedAsync() {
			var repository = ServiceProvider.GetRequiredService<ISessionEventRepository>();

			// Create event
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var newEvent = new ExamStartedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				QuestionNumbers: ImmutableArray.Create(2, 1, 3, 4, 5, 6),
				Deadline: timestamp.AddHours(71)
			);

			// Add event to log
			await repository.AddEventAsync(newEvent);

			// Get logged events
			var loggedEvents = new List<SessionEventBase>();
			await foreach (var loggedEvent in repository.GetAllEventsAsync(sessionId, CancellationToken.None)) {
				loggedEvents.Add(loggedEvent);
			}

			// Validate logged events
			loggedEvents.Count.Should().Be(1);
			loggedEvents[0].Should().BeOfType<ExamStartedEvent>();
			var examStartedEvent = (ExamStartedEvent)loggedEvents[0];
			examStartedEvent.SessionId.Should().Be(sessionId);
			examStartedEvent.Timestamp.Should().Be(timestamp);
			examStartedEvent.QuestionNumbers.Should().ContainInOrder(2, 1, 3, 4, 5, 6);
			examStartedEvent.Deadline.Should().BeCloseTo(timestamp.AddHours(71), TimeSpan.FromMilliseconds(1));
		}
	}
}
#endif
