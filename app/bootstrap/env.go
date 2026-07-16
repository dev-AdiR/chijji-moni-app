package bootstrap

import (
	"chijji-moni-backend-go/domain"
	"fmt"
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
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file", err)
		panic(fmt.Errorf("Error loading .env file %w", err))
	}

	return &domain.Env{
		SupabaseUrl: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_SECRET_KEY"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}
