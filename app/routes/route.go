package routes

import (
	"chijji-moni-backend-go/bootstrap"
	"chijji-moni-backend-go/domain"
	"chijji-moni-backend-go/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *bootstrap.Application) {
	publicRouter := app.Router.Group("/api")
	RegisterAuthRoutes(publicRouter.Group("/auth"), app)

	protectedRouter := app.Router.Group("/api")
	protectedRouter.Use(VerifyToken(app.Env))
	RegisterExpenseRoutes(protectedRouter.Group("/expenses"), app)
	RegisterTempRoutes(protectedRouter.Group("/temp"))
}

func VerifyToken(env *domain.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token == "" {
			log.Println("Auth token missing")
			c.String(http.StatusUnauthorized, fmt.Sprintln("Missing auth token"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")
		fmt.Println(tokenString)
		claims, err := utils.VerifyToken(env, tokenString)

		if err != nil {
			fmt.Println(err)
			c.String(http.StatusUnauthorized, fmt.Sprintln(err))
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Set("username", claims.Username)
		c.Next()
	}
}
