package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"gome/api/models"
	"gome/api/responses"
	"gome/config"
	"net/http"
	"time"
)

// GenerateJWT creates a new token to the client
func GenerateJWT(user models.User) (string, error)  {
	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer: "Zhongkang Shen",
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.SECRETKEY)
}

// ExtractToken retrieves the token from headers ans Query
func ExtractToken(w http.ResponseWriter, r *http.Request) *jwt.Token  {
	token, err := request.ParseFromRequest(
		r,
		request.OAuth2Extractor,
		func(t *jwt.Token) (interface{}, error) {
			return config.SECRETKEY, nil
		},
		request.WithClaims(&models.Claim{}),
		)

	if err != nil {
		code := http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				err = errors.New("your token has expired")
				responses.ERROR(w, code, err)
				return nil
			case jwt.ValidationErrorClaimsInvalid:
				err = errors.New("the signature is invalid")
				responses.ERROR(w, code ,err)
				return nil
			default:
				responses.ERROR(w, code, err)
				return nil
			}
		}
	}
	return token
}