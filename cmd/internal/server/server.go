package server

import (
	"net/http"
	"vkmanagment/cmd/internal/config"
	"vkmanagment/cmd/internal/handlers"

	"github.com/crewjam/saml/samlsp"
)

func StartHTTPServer(samlSP *samlsp.Middleware, cfg *config.Config) {
	http.Handle("/saml/", samlSP)
	http.Handle("/", samlSP.RequireAccount(http.HandlerFunc(handlers.Hello)))
	http.Handle("/verify", http.HandlerFunc(handlers.ValidateToken))

	err := http.ListenAndServeTLS(cfg.ListenAddr, cfg.ServerCert, cfg.ServerKey, nil)
	if err != nil {
		panic("HTTP sunucusu başlatılırken hata: " + err.Error())
	}
}
