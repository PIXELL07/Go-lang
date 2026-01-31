package config

import "os"

var JWT_SECRET string

func LoadEnv() {
	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		JWT_SECRET = "supersecretkey"
	}
}
