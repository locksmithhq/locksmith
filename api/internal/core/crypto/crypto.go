package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const prefix = "aes256gcm:"

var errMissingKey = errors.New("ENCRYPTION_KEY env var is not set")

func loadKey() ([]byte, error) {
	raw := os.Getenv("ENCRYPTION_KEY")
	if raw == "" {
		return nil, errMissingKey
	}
	key, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("crypto: invalid ENCRYPTION_KEY (must be base64): %w", err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("crypto: ENCRYPTION_KEY must decode to 32 bytes (got %d)", len(key))
	}
	return key, nil
}

// Encrypt returns aes256gcm:<base64(nonce+ciphertext)>.
func Encrypt(plaintext string) (string, error) {
	key, err := loadKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return prefix + base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt handles both encrypted (aes256gcm: prefix) and legacy plaintext values.
// Legacy values are returned as-is so existing data keeps working until next write.
func Decrypt(value string) (string, error) {
	if !strings.HasPrefix(value, prefix) {
		return value, nil
	}

	key, err := loadKey()
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(value, prefix))
	if err != nil {
		return "", fmt.Errorf("crypto: base64 decode failed: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("crypto: ciphertext too short")
	}

	plaintext, err := gcm.Open(nil, data[:nonceSize], data[nonceSize:], nil)
	if err != nil {
		return "", fmt.Errorf("crypto: decryption failed: %w", err)
	}

	return string(plaintext), nil
}
