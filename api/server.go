package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	//"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/raulickis/api-myapp/config"
	"github.com/raulickis/api-myapp/internal/enderecos"
	"github.com/raulickis/api-myapp/internal/usuarios"
	//"github.com/raulickis/api-myapp/tools"
	"go.elastic.co/apm/module/apmgin"

)

func Run() {

	r := SetupRoutes()
	_ = r.Run(":" + config.APP_PORT)
}

func SetupRoutes() *gin.Engine {

	router := gin.New()
	//router.Use(apmgin.Middleware(router))
	router.Use(
		apmgin.Middleware(router),
		//gin.Recovery(), // Não é necessário com o uso do Elastic APM
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/check"),
	)
	router.GET("/health/check", CheckHealth)


	//// Optional custom metrics list
	//customMetrics := []*tools.Metric{
	//	&tools.Metric{
	//		ID:				"1234",				// optional string
	//		Name:			"test_metric",			// required string
	//		Description:	"Counter test metric",	// required string
	//		Type:			"counter",			// required string
	//	},
	//	&tools.Metric{
	//		ID:				"1235",				// Identifier
	//		Name:			"test_metric_2",		// Metric Name
	//		Description:	"Summary test metric",	// Help Description
	//		Type:			"summary", // type associated with prometheus collector
	//	},
	//	// Type Options:
	//	//	counter, counter_vec, gauge, gauge_vec,
	//	//	histogram, histogram_vec, summary, summary_vec
	//}
	//
	//prom := tools.NewPrometheus("gin", customMetrics)
	//
	//// Preserve low cardinality when request have params in path
	//prom.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
	//	url := c.Request.URL.Path
	//	for _, p := range c.Params {
	//		if p.Key == "id" {
	//			url = strings.Replace(url, p.Value, ":id", 1)
	//			break
	//		}
	//	}
	//	return url
	//}
	//
	//prom.Use(router)

	//router.Any("/metrics", gin.WrapH(promhttp.Handler()))

	// API
	cadastroRouter := router.Group("/cadastro/usuario")
	{
		cadastroRouter.POST("", usuarios.InserirUsuario)
		cadastroRouter.GET("", usuarios.ListarUsuarios)
		cadastroRouter.GET("/:id", usuarios.ObterUsuario)
		cadastroRouter.PUT("/:id", usuarios.AtualizarUsuario)
		cadastroRouter.DELETE("/:id", usuarios.ExcluirUsuario)
	}

	cadastroEnderecosRouter := router.Group("/cadastro/endereco")
	{
		cadastroEnderecosRouter.POST("", enderecos.InserirEndereco)
		cadastroEnderecosRouter.GET("", enderecos.ListarEnderecos)
		cadastroEnderecosRouter.GET("/:id", enderecos.ObterEndereco)
		cadastroEnderecosRouter.PUT("/:id", enderecos.AtualizarEndereco)
		cadastroEnderecosRouter.DELETE("/:id", enderecos.ExcluirEndereco)
	}

	return router
}

func CheckHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func PrometheusHandler(ctx *gin.Context) {
	promhttp.Handler()
}
