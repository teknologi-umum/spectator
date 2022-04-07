using System.Threading.Tasks;
using Spectator.Protos.Input;

namespace Spectator.Protos.HubInterfaces {
	public interface IEventHub {
		Task<EventReply> MouseClickAsync(MouseClickRequest request);
		Task<EventReply> MouseMoveAsync(MouseMoveRequest request);
		Task<EventReply> MouseScrollAsync(MouseScrollRequest request);
		Task<EventReply> KeystrokeAsync(KeystrokeRequest request);
	}
}
