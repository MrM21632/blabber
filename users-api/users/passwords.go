package users

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2idParams struct {
	Memory  uint32 // Note: Minimum 19MiB recommended
	Time    uint32 // Note: 2 iterations are deemed sufficient with recommended memory configurations
	Threads uint8  // Note: 1 thread is deemed sufficient with recommended memory and iteration configurations
	Saltlen uint32 // Note: 16 bytes is recommended for salt length
	Hashlen uint32 // Note: 32 bytes is recommended for hash length
}

var (
	ErrInvalidHash  = errors.New("encoded hash is not correctly formatted")
	ErrWrongVersion = errors.New("versions of argon2id are not compatible")
)

// Generate an encoded Argon2id hash for the given password, using the provided parameters.
//
// Note: these parameters should be standardized across your application to avoid any
// issues with encoding and later validating passwords.
func GenerateHash(password string, params *Argon2idParams) (string, error) {
	// Generate random salt for password key
	salt := make([]byte, params.Saltlen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Make the result key
	hash := argon2.IDKey([]byte(password), salt, params.Time, params.Memory, params.Threads, params.Hashlen)

	// Encode and return
	hash_to_b64 := base64.RawStdEncoding.EncodeToString(hash)
	salt_to_b64 := base64.RawStdEncoding.EncodeToString(salt)
	encoded_password := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		params.Memory,
		params.Time,
		params.Threads,
		salt_to_b64,
		hash_to_b64,
	)
	return encoded_password, nil
}

// Compares the given password to the password represented by the encoded Argon2id hash.
func ComparePasswordToHash(password, encoded string) (match bool, err error) {
	// Decode provided hash
	params, salt, hash, err := DecodeHash(encoded)
	if err != nil {
		return false, err
	}

	// Encode provided password
	other_hash := argon2.IDKey([]byte(password), salt, params.Time, params.Memory, params.Threads, params.Hashlen)

	// Using subtle.ConstantTimeCompare here helps prevent timing attacks
	if subtle.ConstantTimeCompare(hash, other_hash) == 1 {
		return true, nil
	}
	return false, nil
}

// Decodes the given Argon2id hash into the original parameters, salt, and hash.
func DecodeHash(encoded string) (p *Argon2idParams, salt, hash []byte, err error) {
	// Split encoded string
	values := strings.Split(encoded, "$")
	if len(values) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	// Validate algorithm version matches
	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrWrongVersion
	}

	// Start parsing the parameters
	// First up: memory, time, and threads
	p = &Argon2idParams{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Time, &p.Threads)
	if err != nil {
		return nil, nil, nil, err
	}

	// Parse the salt and its length
	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.Saltlen = uint32(len(salt))

	// Parse the hash and its length
	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.Hashlen = uint32(len(hash))

	return p, salt, hash, nil
}
