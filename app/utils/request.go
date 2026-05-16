package utils

import (
	"chijji-moni-backend-go/domain"

	"github.com/gin-gonic/gin"
)

func BindJson(c *gin.Context) (*domain.LoginRequest, error) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
