package bootstrap

import (
	"chijji-moni-backend-go/database"
	"chijji-moni-backend-go/domain"
	"chijji-moni-backend-go/enums"
	"fmt"

	"github.com/supabase-community/supabase-go"
)

var DbMapper = map[int]domain.DBFactory{
	int(enums.Supabase): registerSupabase,
}

func ResgisterDb(env *domain.Env, dbType enums.DbName) domain.DB {

	var RegisterDbType = DbMapper[int(enums.Supabase)]

	return RegisterDbType(env)
}

// #region Private functions
func registerSupabase(env *domain.Env) domain.DB {
	client, err := supabase.NewClient(env.SupabaseUrl, env.SupabaseKey, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("Failed to initialize the client", err)
	}

	return &database.SupabaseDb{
		Client: client,
	}
}

// func registerSomeOtherDb(env *domain.Env) domain.DB {}

// #endregion
