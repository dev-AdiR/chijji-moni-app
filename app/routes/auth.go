package routes

import (
	"chijji-moni-backend-go/bootstrap"
	"chijji-moni-backend-go/domain"
	"chijji-moni-backend-go/helper"
	"chijji-moni-backend-go/repo"
	"chijji-moni-backend-go/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appContext *bootstrap.Application
	userRepo   *repo.UserRepo
)

func RegisterAuthRoutes(rg *gin.RouterGroup, appCtx *bootstrap.Application) {
	appContext = appCtx
	userRepo = &repo.UserRepo{
		Client: appContext.Client,
	}
	rg.POST("/login", login)
	rg.POST("/signup", signUp)
}

func login(c *gin.Context) {

	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	username, password := req.Username, req.Password

	if helper.ValidateUserLoginRequest(username, password) != nil {
		c.String(http.StatusBadRequest, fmt.Sprintln(domain.ErrInvalidRequest))
		return
	}

	result, err := userRepo.Fetch(username)

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintln(err))
	}

	err = utils.ComparePasswordHash([]byte(result.Password), []byte(password))

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintln("Invalid password"))
		return
	}

	tokenString, err := utils.GenerateToken(appContext.Env, result)

	c.JSON(200, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       result.Id,
			"username": result.Username,
		},
	})
}

func signUp(c *gin.Context) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	username, password := req.Username, req.Password

	hashedPassword, err := utils.HashPassowrd(password)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	go func() {
		userRepo.Create(username, string(hashedPassword))
	}()

	c.JSON(http.StatusOK, fmt.Sprintln("Success"))
}
