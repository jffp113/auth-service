package keystore

import "errors"

type PrivateKey any
type PublicKey any

type KeyPair struct {
	PrivateKey
	PublicKey
}

type Lookuper interface {
	Lookup(kid string) (KeyPair, error)
}

//type Storer interface {
//	Store(Kid string, priv PrivateKey, pub PublicKey)
//}

type ImmutableKeystore struct {
	keys map[string]KeyPair
}

func New(keys ...NewKey) *ImmutableKeystore {
	km := make(map[string]KeyPair)

	for _, k := range keys {
		km[k.Kid] = KeyPair{
			PrivateKey: k.PrivKey,
			PublicKey:  k.PubKey,
		}
	}

	//KeyPair()
	return &ImmutableKeystore{
		keys: km,
	}
}

func (ks *ImmutableKeystore) Lookup(kid string) (KeyPair, error) {
	if k, ok := ks.keys[kid]; ok {
		return k, nil
	}

	return KeyPair{}, errors.New("key not found")
}
