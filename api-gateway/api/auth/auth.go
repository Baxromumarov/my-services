package auth

import (
	"time"

	"github.com/baxromumarov/my-services/api-gateway/pkg/logger"
	"github.com/dgrijalva/jwt-go"
)

type JwtHandler struct {
	Sub        string
	Iss        string
	Exp        string
	Iat        string
	Aud        []string
	Role       string
	Token      string
	SigningKey string
	Log        logger.Logger
}

// GenerateAuthJWT ...
func (jwtHandler *JwtHandler) GenerateJwt() (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	
	claims["sub"] = jwtHandler.Sub
	claims["iss"] = jwtHandler.Iss
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role
	claims["aud"] = jwtHandler.Aud
	
	access, err = accessToken.SignedString([]byte(jwtHandler.SigningKey))
	if err !=nil {
		jwtHandler.Log.Error("error generating access token",logger.Error(err))
		return
	}

	refresh,err = refreshToken.SignedString([]byte(jwtHandler.SigningKey))
	if err !=nil {
		jwtHandler.Log.Error("error generating refresh token",logger.Error(err))
		return
	}
	
	return access, refresh, nil

}

//ExtractClaims ...
func(jwtHandler *JwtHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err error
	)
	token, err = jwt.Parse(jwtHandler.Token,func(t *jwt.Token)(interface{},error){
		return []byte(jwtHandler.SigningKey),nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		jwtHandler.Log.Error("invalid jwt token")
		return nil, err
	}
	return claims, nil
}
