using System;
using FluentAssertions;

namespace Spectator.Piston.Tests.Utilities {
	public static class ObjectShouldExtension {
		public static AndConstraint<T> Should<T>(this T obj, Action<T> constraints) {
			constraints.Invoke(obj);
			return new AndConstraint<T>(obj);
		}
	}
}
