using System.Threading.Tasks;
using Spectator.Protos.Input;

namespace Spectator.Protos.HubInterfaces {
	public interface IEventHub {
		Task<EventReply> LogMouseClickedAsync(MouseClickInfo mouseClickInfo);
		Task<EventReply> LogMouseMovedAsync(MouseMoveInfo mouseMoveInfo);
		Task<EventReply> LogMouseScrolledAsync(MouseScrollInfo mouseScrollInfo);
		Task<EventReply> LogKeystrokeAsync(KeystrokeInfo keystrokeInfo);
		Task<EventReply> LogWindowSizedAsync(WindowSizeInfo windowSizeInfo);
	}
}
