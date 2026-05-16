package bootstrap

import "github.com/gin-gonic/gin"

var r = gin.Default()

func ConfigureNetwork() *gin.Engine {
	configureProxies()

	return r
}

func configureProxies() {
	r.SetTrustedProxies([]string{"127.0.0.1"})
}
