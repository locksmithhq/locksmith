package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/crypto/argon2"
)

const (
	memory      uint32 = 64 * 1024
	iterations  uint32 = 3
	parallelism uint8  = 2
	saltLength         = 16
	keyLength   uint32 = 32
)

var (
	dummyHashOnce sync.Once
	dummyHash     string
)

// PerformDummyVerify runs a full argon2id verification against a dummy hash to
// equalize timing when an account is not found, preventing user enumeration via
// response timing (CWE-204).
func PerformDummyVerify(password string) {
	dummyHashOnce.Do(func() {
		h, err := Hash("_locksmith_dummy_noop_password_")
		if err != nil {
			panic("hasher: failed to compute dummy hash: " + err.Error())
		}
		dummyHash = h
	})
	Verify(password, dummyHash) //nolint:errcheck
}

func Hash(password string) (string, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		iterations,
		parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encoded, nil
}

func Verify(password, encoded string) (bool, error) {
	if !strings.HasPrefix(encoded, "$argon2id$") {
		return false, errors.New("unsupported hash format")
	}

	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, err
	}
	if version != argon2.Version {
		return false, errors.New("incompatible argon2 version")
	}

	var mem, iter uint32
	var par uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &iter, &par); err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	comparisonHash := argon2.IDKey([]byte(password), salt, iter, mem, par, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}
