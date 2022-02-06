package common

const (
	// BucketInputEvents is the bucket name for storing
	// keystroke events, window events, and mouse events.
	BucketInputEvents = "input_events"
	// BucketSessionEvents is the bucket name for storing
	// the session events, including their personal information.
	BucketSessionEvents = "session_events"
	// BucketFileEvents is the bucket name for storing
	// the file events, most importantly the URL to the MinIO storage.
	BucketFileEvents = "file_results"
	// BucketInputStatistics is the bucket name for storing
	// the input statistics, including their personal information.
	BucketInputStatisticEvents = "input_statistics"
)
