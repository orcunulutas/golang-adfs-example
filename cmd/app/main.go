package main

import (
	"log"
	"vkmanagment/cmd/internal/config"
	"vkmanagment/cmd/internal/saml"
	"vkmanagment/cmd/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config yüklenirken hata: %v", err)
	}

	samlSP, err := saml.InitializeSAMLServiceProvider(cfg)
	if err != nil {
		log.Fatalf("SAML SP başlatılırken hata: %v", err)
	}

	server.StartHTTPServer(samlSP, cfg)
}
