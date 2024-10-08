package utils

import (
	"io"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestDecodeHash(t *testing.T) {
	log.SetOutput(io.Discard)

	var tests = []struct {
		name       string
		encoded    string
		wantParams *Argon2IDParams
		wantSalt   []byte
		wantHash   []byte
		wantErr    bool
	}{
		{
			"Successful Decoding",
			"$argon2id$v=19$m=19456,t=2,p=1$OUhHc1FMODFzZ0VLZkU1NQ$RukCwGtr0IOQx0Cfg1D01yXGlMRXu8yQhR+WHznLjvw",
			&Argon2IDParams{19 * 1024, 2, 1, 16, 32},
			[]byte("9HGsQL81sgEKfE55"),
			[]byte{
				0x46, 0xe9, 0x02, 0xc0, 0x6b, 0x6b, 0xd0, 0x83,
				0x90, 0xc7, 0x40, 0x9f, 0x83, 0x50, 0xf4, 0xd7,
				0x25, 0xc6, 0x94, 0xc4, 0x57, 0xbb, 0xcc, 0x90,
				0x85, 0x1f, 0x96, 0x1f, 0x39, 0xcb, 0x8e, 0xfc,
			},
			false,
		},
		{
			"Unsuccessful Decoding",
			"$argon2id$v=19$m=abcd,t=2,p=1$OUhHc1FMODFzZ0VLZkU1NQ$RukCwGtr0IOQx0Cfg1D01yXGlMRXu8yQhR+WHznLjvw",
			nil,
			nil,
			nil,
			true,
		},
		{
			"Invalid Hash",
			"$argon2id$v=19$m=19456,t=2,p=1$foo=bar$OUhHc1FMODFzZ0VLZkU1NQ$RukCwGtr0IOQx0Cfg1D01yXGlMRXu8yQhR+WHznLjvw",
			nil,
			nil,
			nil,
			true,
		},
		{
			"Wrong Argon2id Algorithm Version",
			"$argon2id$v=20$m=19456,t=2,p=1$OUhHc1FMODFzZ0VLZkU1NQ$RukCwGtr0IOQx0Cfg1D01yXGlMRXu8yQhR+WHznLjvw",
			nil,
			nil,
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, salt, hash, err := DecodeHash(tt.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeHash() error = %v, want = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(salt, tt.wantSalt) {
				t.Errorf("DecodeHash() salt = %v, want = %v", salt, tt.wantSalt)
				return
			}
			if !reflect.DeepEqual(hash, tt.wantHash) {
				t.Errorf("DecodeHash() hash = %v, want = %v", hash, tt.wantHash)
				return
			}
			if !reflect.DeepEqual(params, tt.wantParams) {
				t.Errorf("DecodeHash() params = %v, want = %v", params, tt.wantParams)
			}
		})
	}
}

func TestComparePasswordToHash(t *testing.T) {
	var tests = []struct {
		name     string
		password string
		encoded  string
		want     bool
		wantErr  bool
	}{
		{
			"Password and Encoding Match",
			"password",
			"$argon2id$v=19$m=19,t=2,p=1$dndOSlBlcDlzT3BuUlZDYw$EUpSlQrPa/BG2Ee3MJOXpg",
			true,
			false,
		},
		{
			"Password and Encoding Do Not Match",
			"password",
			"$argon2id$v=19$m=19,t=2,p=1$QlhvM1pxaHJLR1RUaXM0bw$ERaaEm2foHNh7xLPQ3x8bg",
			false,
			true,
		},
		{
			"Invalid Hash",
			"password",
			"$argon2id$foo$v=19$m=19,t=2,p=1$dndOSlBlcDlzT3BuUlZDYw$EUpSlQrPa/BG2Ee3MJOXpg",
			false,
			true,
		},
		{
			"Wrong Argon2id Algorithm Version",
			"password",
			"$argon2id$v=18$m=19,t=2,p=1$dndOSlBlcDlzT3BuUlZDYw$EUpSlQrPa/BG2Ee3MJOXpg",
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, err := ComparePasswordToHash(tt.password, tt.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComparePasswordToHash() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if match != tt.want {
				t.Errorf("ComparePasswordToHash() = %v, want = %v", match, tt.want)
			}
		})
	}
}

func BenchmarkGenerateHashDefaultParams(b *testing.B) {
	log.SetOutput(io.Discard)

	// Default params, suits our purposes here
	params := GetArgon2IDConfig()
	password := "password"

	for i := 0; i < b.N; i++ {
		GenerateHash(password, &params)
	}
}

func BenchmarkGenerateHashCustomParams(b *testing.B) {
	log.SetOutput(io.Discard)

	// Custom params as used in the deployed app
	b.Setenv("ARGON2_MEMSIZE", "65536")
	b.Setenv("ARGON2_ITERATIONS", "3")
	b.Setenv("ARGON2_THREADS", "2")
	b.Setenv("ARGON2_SALTLEN", "16")
	b.Setenv("ARGON2_HASHLEN", "32")

	params := GetArgon2IDConfig()
	password := "password"

	for i := 0; i < b.N; i++ {
		GenerateHash(password, &params)
	}
}

func FuzzGenerateHash(f *testing.F) {
	log.SetOutput(io.Discard)

	params := GetArgon2IDConfig()
	f.Add("password")
	f.Fuzz(func(t *testing.T, s string) {
		GenerateHash(s, &params)
	})
}
