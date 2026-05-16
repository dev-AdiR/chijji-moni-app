package domain

type Env struct {
	SupabaseUrl string `mapstructure:"SUPABASE_URL" validate:"required"`
	SupabaseKey string `mapstructure:"SUPABASE_SECRET_KEY" validate:"required"`
	JwtSecret   string `mapstructure:"JWT_SECRET" validate:"required"`
}
