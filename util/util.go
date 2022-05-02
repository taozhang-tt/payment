package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func SignPKCS1v15WithDerKey(content, key []byte, hash crypto.Hash) ([]byte, error) {
	privKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKCS1PrivateKey(%v) with error(%v)", key, err)
	}
	privateKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("privKey(%v) assert to *rsa.PrivateKey fail", privKey)
	}
	h := hash.New()
	h.Write(content)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
}

func VerifyPKCS1v15WithDerKey(content, sig, key []byte, hash crypto.Hash) error {
	pub, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return err
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("VerifyPKCS1v15: pubKey(%v) assert *rsa.PublicKey fail", pub)
	}

	h := hash.New()
	h.Write(content)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(publicKey, hash, hashed, sig)
}

func SignPKCS1v15WithPemKey(src, key []byte, hash crypto.Hash) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("pem.Decode(%v) fail", key)
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKCS1PrivateKey(%v) with error(%v)", block.Bytes, err)
	}
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, pri, hash, hashed)
}
