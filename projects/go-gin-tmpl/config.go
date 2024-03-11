package main

import "os"

// Config holds the configuration parameters
type Config struct {
	ConnectionURI  string
	DBName         string
	CollectionName string
}

// NewDefaultConfig creates a new Config with default values
func NewDefaultConfig() *Config {
	return &Config{
		ConnectionURI:  "mongodb://root:pass@localhost/",
		DBName:         "blog",
		CollectionName: "articles",
	}
}

// NewConfig loads the configuration from environment variables or uses default values
func NewConfig() *Config {
	config := NewDefaultConfig()

	if connURI := os.Getenv("MONGODB_CONNECTION_URI"); connURI != "" {
		config.ConnectionURI = connURI
	}
	if dbName := os.Getenv("MONGODB_DB_NAME"); dbName != "" {
		config.DBName = dbName
	}
	if collectionName := os.Getenv("MONGODB_COLLECTION_NAME"); collectionName != "" {
		config.CollectionName = collectionName
	}

	return config
}
