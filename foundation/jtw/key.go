package jtw

func isValidPubKeyFormat(k any) error {
	return nil
}

func isValidPrivKeyFormat(k any) error {
	return nil
}

/**
Name	"alg" Parameter Values	Signing Key Type	Verification Key Type
HMAC signing method2	HS256,HS384,HS512	[]byte	[]byte
RSA signing method3	RS256,RS384,RS512	*rsa.PrivateKey	*rsa.PublicKey
ECDSA signing method4	ES256,ES384,ES512	*ecdsa.PrivateKey	*ecdsa.PublicKey
RSA-PSS signing method5	PS256,PS384,PS512	*rsa.PrivateKey	*rsa.PublicKey
EdDSA signing method6	Ed25519	ed25519.PrivateKey	ed25519.PublicKey
*/
