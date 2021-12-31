using System;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.SessionDomain;
using Spectator.Observables.Redux;

namespace Spectator.Observables.SessionDomain {
	public record SessionStore() : Store<SessionBase?, SessionEventBase>(
		reducer: (state, @event) => state switch {
			null => @event switch {
				SessionStartedEvent e => AnonymousSession.From(e),
				_ => throw new InvalidOperationException("Session already exists")
			},
			AnonymousSession s => @event switch {
				PersonalInfoSubmittedEvent e => s.Apply(e),
				_ => throw new InvalidOperationException("Personal info hasn't been submitted")
			},
			RegisteredSession s => @event switch {
				BeforeExamSAMSubmittedEvent e => s.Apply(e),
				AfterExamSAMSubmittedEvent e => s.Apply(e),
				ExamStartedEvent e => s.Apply(e),
				ExamEndedEvent e => s.Apply(e),
				ExamIDEReloadedEvent e => s,
				DeadlinePassedEvent e => s.Apply(e),
				ExamForfeitedEvent e => s.Apply(e),
				SolutionAcceptedEvent e => s.Apply(e),
				SolutionRejectedEvent e => s.Apply(e),
				_ => throw new NotImplementedException()
			},
			_ => throw new NotImplementedException()
		},
		initialValue: null
	);
}
