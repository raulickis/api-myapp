package api

import (
	"github.com/autorei/api-myapp/config"
	"github.com/autorei/api-myapp/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func Run() {

	r := SetupRoutes()
	_ = r.Run(":" + config.APP_PORT)
}

func SetupRoutes() *gin.Engine {

	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/check"),
		gin.Recovery(),
	)
	router.GET("/health/check", CheckHealth)

	// API
	cadastroRouter := router.Group("/cadastro/usuario")
	{
		cadastroRouter.POST("", usuarios.InserirUsuario)
		cadastroRouter.GET("", usuarios.ListarUsuarios)
		cadastroRouter.GET("/:id", usuarios.ObterUsuario)
		cadastroRouter.PUT("/:id", usuarios.AtualizarUsuario)
		cadastroRouter.DELETE("/:id", usuarios.ExcluirUsuario)
	}

	return router
}

func CheckHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
