package routes

import (
	"chijji-moni-backend-go/bootstrap"
)

func SetupRoutes(app *bootstrap.Application) {
	publicRouter := app.Router.Group("/api")
	RegisterAuthRoutes(publicRouter.Group("/auth"), app)
	RegisterExpenseRoutes(publicRouter.Group("/expenses"), app)
	RegisterTempRoutes(publicRouter.Group("/temp"))
}
