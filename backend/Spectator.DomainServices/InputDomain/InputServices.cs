using System.Threading.Tasks;
using Spectator.DomainEvents.InputDomain;
using Spectator.Repositories;

namespace Spectator.DomainServices.InputDomain {
	public class InputServices {
		private readonly IInputEventRepository _inputEventRepository;

		public InputServices(
			IInputEventRepository inputEventRepository
		) {
			_inputEventRepository = inputEventRepository;
		}

		public Task AddInputEventAsync(InputEventBase @event) {
			return _inputEventRepository.AddEventAsync(@event);
		}
	}
}
