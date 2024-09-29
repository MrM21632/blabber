package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/argon2"
)

type Argon2IDParams struct {
	Memory  uint32 // Minimum 19 MiB recommended
	Time    uint32 // 2 iterations deemed sufficient with recommended memory config
	Threads uint8  // 1 thread is deemed sufficient with recommended memory and time config
	SaltLen uint32 // 16 bytes is recommended for salt length
	HashLen uint32 // 32 bytes is recommended for hash length
}

var (
	ErrInvalidHash  = errors.New("encoded hash is not correctly formatted")
	ErrWrongVersion = errors.New("versions of argon2id are not compatible")
)

// GetArgon2IDConfig creates a new parameter object for the Argon2id algorithm.
func GetArgon2IDConfig() Argon2IDParams {
	var memory, time, threads, saltlen, hashlen uint64
	var err error

	if memory, err = GetenvInteger("ARGON2_MEMSIZE"); err != nil {
		log.Warn("get memory size failed: envvar ARGON2_MEMSIZE not found, using default")
		memory = 19 * 1024
	}
	if time, err = GetenvInteger("ARGON2_ITERATIONS"); err != nil {
		log.Warn("get iterations failed: envvar ARGON2_ITERATIONS not found, using default")
		time = 2
	}
	if threads, err = GetenvInteger("ARGON2_THREADS"); err != nil {
		log.Warn("get threads failed: envvar ARGON2_THREADS not found, using default")
		threads = 1
	}
	if saltlen, err = GetenvInteger("ARGON2_SALTLEN"); err != nil {
		log.Warn("get salt length failed: envvar ARGON2_SALTLEN not found, using default")
		saltlen = 16
	}
	if hashlen, err = GetenvInteger("ARGON2_HASHLEN"); err != nil {
		log.Warn("get key length failed: envvar ARGON2_HASHLEN not found, using default")
		hashlen = 16
	}

	return Argon2IDParams{
		Memory:  uint32(memory),
		Time:    uint32(time),
		Threads: uint8(threads),
		SaltLen: uint32(saltlen),
		HashLen: uint32(hashlen),
	}
}

// GenerateHash generates an encoded Argon2id hash for the given password, using the provided
// parameters.
//
// Parameters should be standardized across your application to avoid any issues with encoding and
// later validating passwords.
func GenerateHash(password string, params *Argon2IDParams) (string, error) {
	// Generate random salt for password key
	salt := make([]byte, params.SaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Make the result key
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Time,
		params.Memory,
		params.Threads,
		params.HashLen,
	)

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

// DecodeHash decodes the given Argon2id hash into the original parameters, salt, and hash.
func DecodeHash(encoded string) (p *Argon2IDParams, salt, hash []byte, err error) {
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
	p = &Argon2IDParams{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Time, &p.Threads)
	if err != nil {
		return nil, nil, nil, err
	}

	// Parse the salt and its length
	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLen = uint32(len(salt))

	// Parse the hash and its length
	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.HashLen = uint32(len(hash))

	return p, salt, hash, nil
}

// ComparePasswordToHash compares the given password to the password represented by the encoded
// Argon2id hash.
func ComparePasswordToHash(password, encoded string) (match bool, err error) {
	// Decode provided hash
	params, salt, hash, err := DecodeHash(encoded)
	if err != nil {
		return false, err
	}

	// Encode provided password
	other_hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Time,
		params.Memory,
		params.Threads,
		params.HashLen,
	)

	// Using subtle.ConstantTimeCompare here helps prevent timing attacks
	if subtle.ConstantTimeCompare(hash, other_hash) == 1 {
		return true, nil
	}
	return false, errors.New("password does not match encoding")
}
