#!/bin/bash

# Build dizinini tanımlayın, eğer yoksa oluşturulacak
BUILD_DIR="./build"
if [ ! -d "$BUILD_DIR" ]; then
  mkdir -p "$BUILD_DIR"
fi

# Uygulama adı
APP_NAME="vkmanagment"

# Platforma özgü build ayarları (örnek olarak Linux/amd64 için)
GOOS=linux
GOARCH=amd64

# Build komutu
echo "Building $APP_NAME for $GOOS/$GOARCH..."
GOOS=$GOOS GOARCH=$GOARCH go build -o "$BUILD_DIR/$APP_NAME" cmd/app/main.go

echo "Build completed: $BUILD_DIR/$APP_NAME"
