using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
using Microsoft.Extensions.Logging;
using SignalRSwaggerGen.Attributes;
using SignalRSwaggerGen.Enums;
using Spectator.DomainEvents.InputDomain;
using Spectator.DomainServices.InputDomain;
using Spectator.PoormansAuth;
using Spectator.Primitives;
using Spectator.Protos.HubInterfaces;
using Spectator.Protos.Input;

namespace Spectator.Hubs {
	[SignalRHub(autoDiscover: AutoDiscover.MethodsAndParams)]
	public class EventHub : Hub<IEventHub>, IEventHub {
		private readonly PoormansAuthentication _poormansAuthentication;
		private readonly InputServices _inputServices;
		private readonly ILogger<EventHub> _logger;

		public EventHub(
			PoormansAuthentication poormansAuthentication,
			InputServices inputServices,
			ILogger<EventHub> logger
		) {
			_poormansAuthentication = poormansAuthentication;
			_inputServices = inputServices;
			_logger = logger;
		}

		public async Task<EventReply> LogMouseUpAsync(MouseClickInfo mouseClickInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseClickInfo.AccessToken);

			// Send event
			var @event = new MouseUpEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseClickInfo.QuestionNumber,
				X: mouseClickInfo.X,
				Y: mouseClickInfo.Y,
				Button: (MouseButton)mouseClickInfo.Button
			);
			await _inputServices.AddInputEventAsync(@event);

			_logger.LogDebug("MouseUpEvent: {@event}", @event);

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogMouseDownAsync(MouseClickInfo mouseClickInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseClickInfo.AccessToken);

			// Send event
			var @event = new MouseDownEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseClickInfo.QuestionNumber,
				X: mouseClickInfo.X,
				Y: mouseClickInfo.Y,
				Button: (MouseButton)mouseClickInfo.Button
			);
			await _inputServices.AddInputEventAsync(@event);

			_logger.LogDebug("MouseDownEvent: {@event}", @event);

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogMouseMovedAsync(MouseMoveInfo mouseMoveInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseMoveInfo.AccessToken);

			// Send event
			var @event = new MouseMovedEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseMoveInfo.QuestionNumber,
				X: mouseMoveInfo.X,
				Y: mouseMoveInfo.Y,
				Direction: (MouseDirection)mouseMoveInfo.Direction
			);
			await _inputServices.AddInputEventAsync(@event);

			_logger.LogDebug("MouseMovedEvent: {@event}", @event);

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogMouseScrolledAsync(MouseScrollInfo mouseScrollInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseScrollInfo.AccessToken);

			// Send event
			var @event = new MouseScrolledEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseScrollInfo.QuestionNumber,
				X: mouseScrollInfo.X,
				Y: mouseScrollInfo.Y,
				Delta: mouseScrollInfo.Delta
			);
			await _inputServices.AddInputEventAsync(@event);

			_logger.LogDebug("MouseScrolledEvent: {@event}", @event);

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogKeystrokeAsync(KeystrokeInfo keystrokeInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(keystrokeInfo.AccessToken);

			// Send event
			var @event = new KeystrokeEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: keystrokeInfo.QuestionNumber,
				KeyChar: keystrokeInfo.KeyChar,
				Shift: keystrokeInfo.Shift,
				Alt: keystrokeInfo.Alt,
				Control: keystrokeInfo.Control,
				Meta: keystrokeInfo.Meta,
				UnrelatedKey: keystrokeInfo.UnrelatedKey
			);
			await _inputServices.AddInputEventAsync(@event);

			_logger.LogDebug("KeystrokeEvent: {@event}", @event);

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogWindowSizedAsync(WindowSizeInfo windowSizeInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(windowSizeInfo.AccessToken);

			// Send event
			var @event = new WindowSizedEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: windowSizeInfo.QuestionNumber,
				Width: windowSizeInfo.Width,
				Height: windowSizeInfo.Height
			);
			await _inputServices.AddInputEventAsync(@event);

			_logger.LogDebug("WindowSizedEvent: {@event}", @event);

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}
	}
}
