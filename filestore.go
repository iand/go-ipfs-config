package config

// Filestore tracks the configuration of the filestore.
type Filestore struct {
	Path string `json:",omitempty"`
}

// FilestorePath returns the default filestore path given a configuration root
func FilestorePath(configroot string) (string, error) {
	return Path(configroot, "")
}
