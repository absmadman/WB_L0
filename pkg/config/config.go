package config

type Config struct {
	DatabaseName     string
	DatabaseLogin    string
	DatabasePassword string
	ClientId         string
	ClusterId        string
	CacheSize        int
}

func NewConfig() *Config {
	return &Config{
		DatabaseName:     "postgres",
		DatabaseLogin:    "postgres",
		DatabasePassword: "postgres",
		ClusterId:        "test-cluster",
		ClientId:         "receiver",
		CacheSize:        256,
	}
}
