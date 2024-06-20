package middleware

import (
	"Pembayaran-Spp/model"
	"Pembayaran-Spp/model/json"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	applicationNone  = "incubation-golang"
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignaturKey   = []byte("incubation-golang") // TODO: replace with secure key
)

func GenerateTokenJwt(username string, expiredAt int, roles string) (string, error) {
	loginExpDuration := time.Duration(expiredAt) * time.Minute
	myExpiresAt := time.Now().Add(loginExpDuration).Unix()
	claims := model.JWTClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationNone,
			ExpiresAt: myExpiresAt,
		},
		Username: username,
		Roles:    roles,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(jwtSignaturKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseBadRequest(c, []json.ValidationField{}, "Invalid token", "01", "01")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &model.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSignaturKey, nil
		})

		if err != nil {
			json.NewResponseBadRequest(c, []json.ValidationField{}, "Invalid token", "01", "01")
			c.Abort()
			return
		}

		if !token.Valid {
			json.NewResponseBadRequest(c, []json.ValidationField{}, "Forbidden", "03", "03")
			c.Abort()
			return
		}

		// validation roles
		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if role == claims.Roles {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			json.NewResponseBadRequest(c, []json.ValidationField{}, "Forbidden", "03", "03")
			c.Abort()
			return
		}

		c.Next()
	}
}
