package model

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("829cabbdf7f9dd30420eedfa6a2b58518e23169affa8825f0ba7d47df5e733b17e84f87bf0cb774a8e3b17853b7366f08b44cd705b6e7776640e8e05208bbb8e4e64a8639857d1970ad5301b5e61a6adcbc6d7dcfdea44f7ec2ad0a2b16cd3f13ab172c1e8fcfc65abeb3f7bede3be713fe20c073fc21dc25f089b01c6705bb82bea8ee4804ad18499afcdb947a23467b78b0a86db45484e738c3cd806ad873e9962601a2ab5ca06b4201b9aff0de45e6c2105173e730e3ee81675b68a74b806683d43ac5a4e2cedd9c149bb3337808db6fc05e708f757b106625260216a549578b07ab54493a5ea9ae74510c49b643718f9cac0487ff8a9bee56214bf715371")

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	TeamId   int64  `json:"teamId"`
	jwt.RegisteredClaims
}

func GenerateJWT(id int64, username string, teamId int64) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Id:       id,
		Username: username,
		TeamId:   teamId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, fmt.Errorf("invalid signature")
		}
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
