using System;

namespace Spectator.Piston.Exceptions; 

public class CannotParseStdoutException : Exception {
	public string Stdout { get; }
	public string Stderr { get; }

	public CannotParseStdoutException(string stdout, string stderr) {
		Stderr = stderr;
		Stdout = stdout;
	}
}
