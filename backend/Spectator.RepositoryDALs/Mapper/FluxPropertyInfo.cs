using System.Reflection;

namespace Spectator.RepositoryDALs.Mapper {
	internal record FluxPropertyInfo(
		string FluxFieldName,
		PropertyInfo PropertyInfo
	);
}
