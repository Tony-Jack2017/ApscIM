package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const expired = 24 // hour
const secret = "apsc_im"

type ApscClaims struct {
	AccountID   int16  `json:"account_id"`
	AccountRole string `json:"account_role"`
	jwt.RegisteredClaims
}

func GenerateToken(accountID int16, accountRole string) (string, error) {
	claims := &ApscClaims{
		AccountID:   accountID,
		AccountRole: accountRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr string) (*ApscClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &ApscClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	} else {
		if claims, ok := token.Claims.(*ApscClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, err
		}
	}
}
