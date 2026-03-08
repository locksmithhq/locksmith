package domain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func GenerateRSAKeys() error {
	var err error
	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	PublicKey = &PrivateKey.PublicKey
	return nil
}

func LoadRSAKeys(privateKeyPath, publicKeyPath string) error {
	// Load Private Key
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return errors.New("failed to parse PEM block containing the private key")
	}
	PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	// Load Public Key
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}
	block, _ = pem.Decode(publicKeyBytes)
	if block == nil {
		return errors.New("failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var ok bool
	PublicKey, ok = pub.(*rsa.PublicKey)
	if !ok {
		return errors.New("not an RSA public key")
	}

	return nil
}
