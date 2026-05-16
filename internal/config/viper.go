package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var config StaticConfig
var Debug = false

// load is constructor of public config
func load() *viper.Viper {
	// Init public config
	cfg := viper.New()

	// Set config type
	cfg.SetConfigType("yaml")

	// Set config path
	cfg.AddConfigPath("./")

	// Set config file
	cfg.SetConfigName("config")

	// Read the config file
	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal("fatal error happened while reading config file:", err)
	}

	// Is debug mode?
	if _, err := os.Stat("config_debug.yaml"); err == nil {
		// Init config file
		cfg.SetConfigName("config_debug")

		// Set debug status
		Debug = true

		// Read the debug config file
		if err = cfg.ReadInConfig(); err != nil {
			log.Fatal("fatal error happened while reading debug config file:", err)
		}
	}

	// Unmarshal config
	if err := cfg.Unmarshal(&config); err != nil {
		log.Fatal("fatal error happened while marshalling config:", err)
		return nil
	}

	return cfg
}

// Init loads public config
func Init() {
	load()
}

// Get returns static config
func Get() StaticConfig {
	return config
}
