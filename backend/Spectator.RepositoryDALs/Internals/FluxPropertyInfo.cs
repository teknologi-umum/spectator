using System.Reflection;

namespace Spectator.RepositoryDALs.Internals {
	internal record FluxPropertyInfo(
		string FluxFieldName,
		PropertyInfo PropertyInfo
	);
}
