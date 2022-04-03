using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
using SignalRSwaggerGen.Attributes;
using SignalRSwaggerGen.Enums;
using Spectator.PoormansAuth;
using Spectator.DomainModels.SessionDomain;
using Spectator.Protos.HubInterfaces;
using Spectator.Protos.Events;

namespace Spectator.Hubs {
	[SignalRHub(autoDiscover: AutoDiscover.MethodsAndArgs)]
	public class EventHub : Hub<IEventHub> {
		private readonly PoormansAuthentication _poormansAuthentication;
		private readonly IServiceProvider _serviceProvider;

		public EventHub(
			PoormansAuthentication poormansAuthentication,
			IServiceProvider serviceProvider
		) {
			_poormansAuthentication = poormansAuthentication;
			_serviceProvider = serviceProvider;
		}

		public async Task<EventReply> MouseClickAsync(MouseClickRequest request) {
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// TODO: process the request properly
			Console.WriteLine("MOUSE CLICK", request);

			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> MouseMoveAsync(MouseMoveRequest request) {
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// TODO: process the request properly
			Console.WriteLine("MOUSE MOVE", request);

			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> MouseScrollAsync(MouseScrollRequest request) {
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// TODO: process the request properly
			Console.WriteLine("MOUSE SCROLL", request);

			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> KeystrokeAsync(KeystrokeRequest request) {
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// TODO: process the request properly
			Console.WriteLine("KEYSTROKE", request);

			return new EventReply {
				Message = "OK"
			};
		}
	}
}
