#!/bin/bash

# Ortam değişkenlerini ayarla
export METADATA_URL="https://idpserver/federationmetadata/2007-06/FederationMetadata.xml"
export SESSION_CERT="./sessioncert"
export SESSION_KEY="./sessionkey"
export SERVER_KEY="./serverkey"
export SERVER_CERT="./servercert"
export SERVER_URL="https://appUrl:8000"
export LISTEN_ADDR="0.0.0.0:8000"

# Uygulamayı çalıştır
./build/vkmanagment
