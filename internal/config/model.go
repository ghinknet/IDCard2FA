package config

type StaticConfig struct {
	// Config of system log
	Log struct {
		File struct {
			All string `mapstructure:"all"`
			Err string `mapstructure:"err"`
		} `mapstructure:"file"`
		MaxSize    int  `mapstructure:"maxSize"`
		MaxBackups int  `mapstructure:"maxBackups"`
		MaxAge     int  `mapstructure:"maxAge"`
		Compress   bool `mapstructure:"compress"`
	} `mapstructure:"log"`
	// Credentials for Ghink OpenAPI platform
	GhinkOpenAPI map[string]struct {
		SecretID  string `mapstructure:"secretID"`
		SecretKey string `mapstructure:"secretKey"`
	} `mapstructure:"ghinkOpenAPI"`
}
