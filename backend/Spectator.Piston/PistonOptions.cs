namespace Spectator.Piston {
	public class PistonOptions {
		public string Address { get; set; } = "http://localhost:50051";
		public int MaxConcurrentExecutions { get; set; } = 2;
		public int CompileTimeout { get; set; } = 5000;
		public int RunTimeout { get; set; } = 3000;
		public int MemoryLimit { get; set; } = 200_000_000;
	}
}
