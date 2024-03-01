package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWTSecretKey = os.Getenv("JWT_SECRET_KEY")

// UserClaims, JWT içinde taşınacak özel claim'ler için bir yapıdır.
type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	// Token, genellikle "Authorization" header'ında "Bearer {token}" formatında taşınır
	tokenString := r.Header.Get("Authorization")
	fmt.Fprintf(w, `{"token":"%s"}`, tokenString)
	// "Bearer " öneki kaldırılıyor
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Token'ın çözümlenmesi ve doğrulanması
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Token doğrulanamadı: %v", err)
		return
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		// Token doğrulandı
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"username":"%s","valid":true}`, claims.Username)
	} else {
		// Token geçersiz
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"valid":false}`)
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	// Token'ı doğrula
	c, err := r.Cookie("auth_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// Eğer cookie yoksa, yetkisiz erişim hatası döndür
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Auth token bulunamadı")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"valid":false}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cookie okuma hatası: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"valid":false}`)
		return
	}

	tokenString := c.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Token imza algoritmasının beklediğiniz algoritma ile eşleşip eşleşmediğini kontrol edin
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// HMAC için kullanılan secret key'i döndür
		return []byte(JWTSecretKey), nil
	})
	token.Method.Alg()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"valid":false}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"valid":true}`)
}

// CreateToken, başarılı kimlik doğrulama sonrası bir JWT token oluşturur.
func CreateToken(w http.ResponseWriter, username string) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token'ın geçerlilik süresi 24 saat olarak ayarlanır.
	claims := &UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Token oluşturulurken hata: %v", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Path:     "/",
		Secure:   true, // Güvenli bir bağlantı üzerinde çalıştığınızı varsayarsak
		SameSite: http.SameSiteStrictMode,
	})

	fmt.Fprintf(w, `{"token":"%s"}`, tokenString)
}
