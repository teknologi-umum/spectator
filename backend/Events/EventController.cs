using System.Threading;
using System.Threading.Tasks;
using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Writes;

namespace Spectator.Events; 

public class EventController {
	private readonly InfluxDBClient _db;
	public EventController(InfluxDBClient db) {
		_db = db;
	}
	
	public async Task AddAsync(EventBase e, CancellationToken cancellationToken) {
		var builder = PointData.Measurement("event")
			.Tag("type", e.Type.ToString())
			.Tag("user", e.User.StudentNumber)
			.Field("value", e.Value)
			.Timestamp(e.Date, WritePrecision.Ns);
		var writeApi = _db.GetWriteApiAsync();
		await writeApi.WritePointAsync(builder, cancellationToken);
	}
}
