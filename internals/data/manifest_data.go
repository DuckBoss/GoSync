package data

type Manifest struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Directories struct {
		SourceDirectory        string   `json:"source_directory"`
		DestinationDirectories []string `json:"destination_directories"`
	}
	Settings struct {
		HashAlgorithm       string `json:"hash_algorithm"`
		ScanIntervalSeconds int    `json:"scan_interval_seconds"`
	}
}
