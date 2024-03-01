package config

import "os"

type Config struct {
	MetadataURL string
	SessionCert string
	SessionKey  string
	ServerKey   string
	ServerCert  string
	ServerURL   string
	ListenAddr  string
}

func LoadConfig() (*Config, error) {
	return &Config{
		MetadataURL: os.Getenv("METADATA_URL"),
		SessionCert: os.Getenv("SESSION_CERT"),
		SessionKey:  os.Getenv("SESSION_KEY"),
		ServerKey:   os.Getenv("SERVER_KEY"),
		ServerCert:  os.Getenv("SERVER_CERT"),
		ServerURL:   os.Getenv("SERVER_URL"),
		ListenAddr:  os.Getenv("LISTEN_ADDR"),
	}, nil
}
