package routes

import (
	"chijji-moni-backend-go/bootstrap"
	"chijji-moni-backend-go/domain"
	"chijji-moni-backend-go/repo"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthData struct {
	UserId   float64 `json:"userId"`
	Username string  `json:"username"`
}

var expenseRepo domain.ExpenseRepo

func RegisterExpenseRoutes(rg *gin.RouterGroup, appContext *bootstrap.Application) {
	expenseRepo = &repo.ExpenseRepo{
		Client: appContext.Client,
	}
	rg.Use(ExtractAuthDataMiddleware())
	rg.GET("/month/:year/:month", getExpensesForCurrentMonth)
}

func ExtractAuthDataMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func getExpensesForCurrentMonth(c *gin.Context) {

	year, month := c.Params.ByName("year"), c.Params.ByName("month")

	userIdVal, exists := c.Get("userId")
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

	userId, ok := userIdVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userId type"})
		return
	}
	fmt.Println("user", userId)

	data, err := expenseRepo.Fetch(userId)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
