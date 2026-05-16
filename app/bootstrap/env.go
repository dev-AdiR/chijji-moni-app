package bootstrap

import (
	"chijji-moni-backend-go/domain"
	"os"

	"github.com/joho/godotenv"
)

// func NewEnv() *domain.Env {
// 	var env domain.Env

// 	// Try .env only in local dev
// 	_ = godotenv.Load()

// 	viper.AutomaticEnv()

// 	if err := viper.Unmarshal(&env); err != nil {
// 		log.Fatal(err)
// 	}

// 	return &env
// }

func NewEnv() *domain.Env {
	_ = godotenv.Load()

	return &domain.Env{
		SupabaseUrl: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_SECRET_KEY"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}
