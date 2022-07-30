using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
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

		public EventHub(
			PoormansAuthentication poormansAuthentication,
			InputServices inputServices
		) {
			_poormansAuthentication = poormansAuthentication;
			_inputServices = inputServices;
		}

		public async Task<EventReply> LogMouseUpAsync(MouseClickInfo mouseClickInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseClickInfo.AccessToken);

			// Send event
			await _inputServices.AddInputEventAsync(new MouseUpEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseClickInfo.QuestionNumber,
				X: mouseClickInfo.X,
				Y: mouseClickInfo.Y,
				Button: (MouseButton)mouseClickInfo.Button
			));

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogMouseDownAsync(MouseClickInfo mouseClickInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseClickInfo.AccessToken);

			// Send event
			await _inputServices.AddInputEventAsync(new MouseDownEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseClickInfo.QuestionNumber,
				X: mouseClickInfo.X,
				Y: mouseClickInfo.Y,
				Button: (MouseButton)mouseClickInfo.Button
			));

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogMouseMovedAsync(MouseMoveInfo mouseMoveInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseMoveInfo.AccessToken);

			// Send event
			await _inputServices.AddInputEventAsync(new MouseMovedEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseMoveInfo.QuestionNumber,
				X: mouseMoveInfo.X,
				Y: mouseMoveInfo.Y,
				Direction: (MouseDirection)mouseMoveInfo.Direction
			));

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogMouseScrolledAsync(MouseScrollInfo mouseScrollInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(mouseScrollInfo.AccessToken);

			// Send event
			await _inputServices.AddInputEventAsync(new MouseScrolledEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: mouseScrollInfo.QuestionNumber,
				X: mouseScrollInfo.X,
				Y: mouseScrollInfo.Y,
				Delta: mouseScrollInfo.Delta
			));

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogKeystrokeAsync(KeystrokeInfo keystrokeInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(keystrokeInfo.AccessToken);

			// Send event
			await _inputServices.AddInputEventAsync(new KeystrokeEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: keystrokeInfo.QuestionNumber,
				KeyChar: keystrokeInfo.KeyChar,
				Shift: keystrokeInfo.Shift,
				Alt: keystrokeInfo.Alt,
				Control: keystrokeInfo.Control,
				Meta: keystrokeInfo.Meta,
				UnrelatedKey: keystrokeInfo.UnrelatedKey
			));

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}

		public async Task<EventReply> LogWindowSizedAsync(WindowSizeInfo windowSizeInfo) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(windowSizeInfo.AccessToken);

			// Send event
			await _inputServices.AddInputEventAsync(new WindowSizedEvent(
				SessionId: session.Id,
				Timestamp: DateTimeOffset.UtcNow,
				QuestionNumber: windowSizeInfo.QuestionNumber,
				Width: windowSizeInfo.Width,
				Height: windowSizeInfo.Height
			));

			// Reply OK
			return new EventReply {
				Message = "OK"
			};
		}
	}
}
