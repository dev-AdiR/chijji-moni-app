package bootstrap

import (
	"chijji-moni-backend-go/domain"
	"chijji-moni-backend-go/enums"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Env    *domain.Env
	Client domain.DB
	Router *gin.Engine
}

func App() *Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Client = ResgisterDb(app.Env, enums.Supabase)
	app.Router = ConfigureNetwork()

	return app
}
