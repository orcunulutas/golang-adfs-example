package saml

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/url"
	"vkmanagment/cmd/internal/config"

	"github.com/crewjam/saml/samlsp"
)

func InitializeSAMLServiceProvider(cfg *config.Config) (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair(cfg.SessionCert, cfg.SessionKey)
	if err != nil {
		return nil, err
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, err
	}

	rootURL, err := url.Parse(cfg.ServerURL)
	if err != nil {
		return nil, err
	}

	metadataURL := mustParseURL(cfg.MetadataURL)
	ctx := context.Background()
	client := http.DefaultClient
	idpMetadata, err := samlsp.FetchMetadata(ctx, client, *metadataURL)
	if err != nil {
		log.Fatalf("IDP metadata yüklenirken hata: %v", err)
	}
	samlSP, err := samlsp.New(samlsp.Options{
		URL:               *rootURL,
		Key:               keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:       keyPair.Leaf,
		IDPMetadata:       idpMetadata,
		AllowIDPInitiated: true,
	})
	if err != nil {
		return nil, err
	}

	return samlSP, nil
}

func mustParseURL(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic("URL parse hatası: " + err.Error())
	}
	return parsedURL
}
