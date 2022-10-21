package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"goxhactive-final-project/app/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Jwt struct {
}

func GetScretKey() []byte {

	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	key := os.Getenv("SECRET_KEY")
	str := fmt.Sprintf("%v", key)
	jwtKey := []byte(str)

	return jwtKey
}

func GenerateToken(user models.Users) (models.Token, error) {
	var err error

	claims := jwt.MapClaims{}
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	jwt := models.Token{}

	jwt.AccessToken, err = token.SignedString(GetScretKey())
	if err != nil {
		return jwt, err
	}

	return CreateRefreshToken(jwt)
}

func ValidateToken(signedToken string) (interface{}, error) {

	errResponse := errors.New("Signed")
	token, _ := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil

}

func ValidateRefreshToken(accessToken string, refreshToken string) (models.Users, error) {

	sha1 := sha1.New()

	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	user := models.Users{}

	salt := string(sha1.Sum(nil))[0:16]

	block, err := aes.NewCipher([]byte(salt))

	fmt.Println(block, "Block AES chipher")

	if err != nil {
		fmt.Println(err, "Error get Block AES chipher")
		return user, err
	}

	gcm, err := cipher.NewGCM(block)

	fmt.Println(gcm, "Decode string token")

	if err != nil {
		fmt.Println(err, "Error gcm")
		return user, err
	}

	data, err := base64.URLEncoding.DecodeString(refreshToken)

	fmt.Println(data, "Decode string token")

	if err != nil {
		fmt.Println(err, "Error decode string token")
		return user, err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)

	fmt.Println(string(plain), "Plain text")

	if err != nil {
		return user, err
	}

	if string(plain) != accessToken {
		return user, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}

	parser := jwt.Parser{}

	token, _, err := parser.ParseUnverified(accessToken, claims)

	if err != nil {
		return user, err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return user, errors.New("invalid token")
	}

	return user, nil
}

func CreateRefreshToken(token models.Token) (models.Token, error) {

	sha1 := sha1.New()

	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]

	block, err := aes.NewCipher([]byte(salt))

	if err != nil {

		fmt.Println(err.Error())

		return token, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {

		return token, err

	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)

	if err != nil {

		return token, err

	}

	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))

	return token, nil
}
