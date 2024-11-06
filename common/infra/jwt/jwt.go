package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken  = errors.New("error.invalid_token")
	ErrGenerateToken = errors.New("error.generate_token")
)

type JWT struct {
	JwtSecret []byte
}

type Claims struct {
	UserID             string `json:"userID"`
	Exp                int64  `json:"exp"`
	PlanExpirationDate int64  `json:"planExpirationDate"`
	Role               string `json:"role"`
}

func (j *JWT) CreateToken(claims Claims) (string, error) {
	claims = Claims{
		UserID:             claims.UserID,
		Exp:                time.Now().Add(15 * time.Minute).Unix(),
		PlanExpirationDate: claims.PlanExpirationDate,
		Role:               claims.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":             claims.UserID, // Reivindicação para o ID do usuário
		"exp":                claims.Exp,    // Reivindicação para a data de expiração (15 minutos a partir de agora)
		"planExpirationDate": claims.PlanExpirationDate,
		"role":               claims.Role,
	})

	tokenString, err := token.SignedString(j.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) CreateRefreshToken(claims Claims) (string, error) {
	claims = Claims{
		UserID: claims.UserID,
		Exp:    time.Now().Add(time.Hour * 24 * 7).Unix(), // Expira em 7 dias
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": claims.UserID, // Reivindicação para o ID do usuário
		"exp":    claims.Exp,    // Reivindicação para a data de expiração (7 dias a partir de agora)
	})

	tokenString, err := token.SignedString(j.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) VerifyToken(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.JwtSecret, nil
	})

	if err != nil {
		return Claims{}, err
	}

	if !token.Valid {
		return Claims{}, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Claims{}, fmt.Errorf("invalid claims")
	}

	return Claims{
		UserID:             claims["userID"].(string),
		Exp:                int64(claims["exp"].(float64)),
		PlanExpirationDate: int64(claims["planExpirationDate"].(float64)),
		Role:               claims["role"].(string),
	}, nil
}
