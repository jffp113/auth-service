package jtw

import (
	"com.cross-join.crossviewer.authservice/foundation/keystore"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

var defaultSignMethod jwt.SigningMethod = jwt.SigningMethodHS256

type Validator interface {
	Validate()
}

type Signer interface {
	Sign(claims Claims, exp time.Time, nbf time.Time) (string, error)
}

type JwtWrapper struct {
	ks  keystore.Lookuper
	alg jwt.SigningMethod
}

func New(ks keystore.Lookuper) *JwtWrapper {
	return &JwtWrapper{
		ks:  ks,
		alg: defaultSignMethod,
	}
}

type Claims map[string]interface{}

type Configs struct {
	Exp time.Time
	Nbf time.Time
}

func (j *JwtWrapper) Sign(claims Claims, kid string, cfg Configs) (string, error) {
	privKey, _ := j.ks.Lookup(kid)

	if err := isValidPrivKeyFormat(privKey); err != nil {
		return "", errors.New("invalid key")
	}

	claims["kid"] = kid
	claims["exp"] = cfg.Exp
	claims["nbf"] = cfg.Nbf

	t := jwt.NewWithClaims(defaultSignMethod, jwt.MapClaims(claims))

	signed, err := t.SignedString(privKey)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return signed, nil
}

func (j *JwtWrapper) Validate(tokenStr string) (Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if j.alg != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid := fmt.Sprintf("%v", token.Header["kid"])
		k, err := j.ks.Lookup(kid)

		if err != nil {
			return nil, fmt.Errorf("looking for key with kid %v: %w", kid, err)
		}

		if err := isValidPubKeyFormat(k); err != nil {
			return nil, fmt.Errorf("validating public key: %w", err)
		}

		return k.PublicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Claims(claims), nil
	} else {
		return nil, errors.New("invalid claims")
	}
}
