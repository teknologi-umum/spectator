using System;

namespace Spectator.Piston.Exceptions; 

public class CannotParseStdoutException : ArgumentException {
	public string Stdout { get; }
	public string Stderr { get; }

	CannotParseStdoutException(string message, string stdout, string stderr) : base(message) {
		Stderr = stderr;
		Stdout = stdout;
	}
}
