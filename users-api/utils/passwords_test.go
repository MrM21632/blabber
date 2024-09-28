package utils

import "testing"

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
			false,
		},
		{
			"Invalid Hash Error",
			"password",
			"$argon2id$foo$v=19$m=19,t=2,p=1$dndOSlBlcDlzT3BuUlZDYw$EUpSlQrPa/BG2Ee3MJOXpg",
			false,
			true,
		},
		{
			"Invalid Argon2id Algorithm Version",
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

func BenchmarkGenerateHash(b *testing.B) {
	// Default params, suits our purposes here
	params := GetArgon2IDConfig()
	password := "password"

	for i := 0; i < b.N; i++ {
		GenerateHash(password, &params)
	}
}

func BenchmarkGenerateHashCustomParams(b *testing.B) {
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
