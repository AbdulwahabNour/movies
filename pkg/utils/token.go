package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type IDTokenClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
type RefreshTokenClaims struct {
	UID string
	jwt.StandardClaims
}
type RefreshToken struct {
	Token string
	ID    string
	ExpIn time.Duration
}
type Token struct {
	Plaintext string
	Hash      string
}

func GenerateIDToken(claims IDTokenClaims, privateKey string, exp time.Duration) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(exp).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", fmt.Errorf("could not create a id token: %w", err)
	}
	return tokenString, nil
}

func ValidateIDToken(tokenReq string, publicKey string) (*IDTokenClaims, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}
	claims := IDTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenReq, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid ID token")
	}
	claims, ok := token.Claims.(IDTokenClaims)
	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")
	}

	return &claims, nil

}

func GenerateRefreshToken(claims RefreshTokenClaims, key string, exp time.Duration) (*RefreshToken, error) {

	claims.IssuedAt = time.Now().Unix()
	expiresAt := time.Now().Add(exp)
	claims.ExpiresAt = expiresAt.Unix()
	tokenID, err := uuid.NewRandom()

	if err != nil {
		log.Println("Failed to generate refresh token ID")
		return nil, err
	}
	claims.Id = tokenID.String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("could not create a refresh token: %w", err)
	}

	return &RefreshToken{
		Token: tokenString,
		ID:    tokenID.String(),
		ExpIn: expiresAt.Sub(time.Now()),
	}, nil

}
func ValidateRefreshToken(tokenString string, key string) (*RefreshTokenClaims, error) {
	claims := RefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(RefreshTokenClaims)
	if !ok {
		return nil, fmt.Errorf("couldn't parse claims")
	}
	return &claims, nil
}

// func ExtractJWTFromRequest(r *http.Request, key string) (*Claims, error) {

// 	tokenString, err := GetBearerToken(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims := &Claims{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return nil, nil
// 	})

// 	if err != nil {
// 		if errors.Is(err, jwt.ErrSignatureInvalid) {
// 			return nil, errors.New("invalid token signature")
// 		}
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	return claims, nil
// }

func GetBearerToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("not authorized")
	}

	authHeaderParts := strings.Split(auth, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "token" {
		return "", errors.New("not authorized")
	}

	return authHeaderParts[1], nil
}

func Byte(n uint) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandToken() (*Token, error) {
	rand, err := Byte(32)
	if err != nil {
		return nil, fmt.Errorf("something happened during generate random byte, error: %w ", err)
	}
	plainText := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(rand)
	hash := Hash(plainText)
	return &Token{
		Plaintext: plainText,
		Hash:      hash,
	}, nil

}
func Hash(data string) string {
	h := sha256.Sum256([]byte(data))
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(h[:])
}
