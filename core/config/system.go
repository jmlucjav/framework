package config

// ClusterConfig stores cluster settings
type ClusterConfig struct {
	Name  string   `config:"name"`
	Seeds []string `config:"seeds"`
}

// NetworkConfig stores network settings
type NetworkConfig struct {
	Host string `config:"host"`

	APIBinding     string `config:"api_bind"`
	HTTPBinding    string `config:"http_bind"`
	ClusterBinding string `config:"cluster_bind"`
}

// NodeConfig stores node settings
type NodeConfig struct {
	Name string `config:"name"`
}

// PathConfig stores path settings
type PathConfig struct {
	Data string `config:"data"`
	Log  string `config:"logs"`
	Cert string `config:"certs"`
}

// SystemConfig is a high priority config, init from the environment or startup, can't be changed on the fly, need to restart to make config apply
type SystemConfig struct {
	ConfigFile string

	ClusterConfig ClusterConfig `config:"cluster"`

	NetworkConfig NetworkConfig `config:"network"`

	NodeConfig NodeConfig `config:"node"`

	PathConfig PathConfig `config:"path"`

	CookieSecret string `config:"cookie_secret"`

	AllowMultiInstance bool `config:"allow_multi_instance"`
	MaxNumOfInstance   int  `config:"max_num_of_instances"`
	TLSEnabled         bool `config:"tls_enabled"`
}
