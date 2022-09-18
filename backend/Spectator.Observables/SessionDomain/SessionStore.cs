using System;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.SessionDomain;
using Spectator.Observables.Redux;

namespace Spectator.Observables.SessionDomain {
	public record SessionStore(SessionBase InitialValue) : Store<SessionBase, SessionEventBase>(
		reducer: (state, @event) => state switch {
			null => @event switch {
				SessionStartedEvent e => AnonymousSession.From(e),
				_ => throw new InvalidOperationException("Session already exists")
			},
			AnonymousSession s => @event switch {
				PersonalInfoSubmittedEvent e => s.Apply(e),
				LocaleSetEvent e => s.Apply(e),
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
				TestAcceptedEvent e => s.Apply(e),
				TestRejectedEvent e => s.Apply(e),
				LocaleSetEvent e => s.Apply(e),
				_ => throw new InvalidProgramException("Unhandled event")
			},
			_ => throw new InvalidProgramException("Unhandled event")
		},
		initialValue: InitialValue
	) {
		public SessionStore(SessionStartedEvent initialEvent) : this(AnonymousSession.From(initialEvent)) { }
	}
}
