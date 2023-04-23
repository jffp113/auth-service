package hasher

import (
	"errors"
	"fmt"
	"github.com/alexedwards/argon2id"
)

var params = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func CreateHash(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, params)

	if err != nil {
		return "", fmt.Errorf("creating password hash: %w", err)
	}

	return hash, nil
}

func VerifyHash(password, hash string) error {
	match, err := argon2id.ComparePasswordAndHash(password, hash)

	if err != nil || !match {
		return errors.New("invalid password")
	}

	return nil
}
