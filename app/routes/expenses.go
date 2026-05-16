package routes

import (
	"chijji-moni-backend-go/bootstrap"
	"chijji-moni-backend-go/repo"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthData struct {
	UserId   float64 `json:"userId"`
	Username string  `json:"username"`
}

func RegisterExpenseRoutes(rg *gin.RouterGroup, appContext *bootstrap.Application) {
	userRepo = &repo.UserRepo{
		Client: appContext.Client,
	}

	rg.Use(ExtractAuthDataMiddleware())
	rg.GET("/month/:year/:month", getExpensesForCurrentMonth)
}

func ExtractAuthDataMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("X-User-Id")
		username := c.GetHeader("X-Username")

		c.Set("userId", userId)
		c.Set("username", username)

		c.Next()
	}
}

func getExpensesForCurrentMonth(c *gin.Context) {

	year, month := c.Params.ByName("year"), c.Params.ByName("month")

	_, exists := c.Get("userId")
	if !exists {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		// handle error
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		// handle error
	}

	firstDay := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	_ = firstDay.AddDate(0, 1, -1)

	// data, err := userRepo.Client.Select()
}
