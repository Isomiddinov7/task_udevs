package api

import (
	_ "task_udevs/api/docs"
	"task_udevs/api/handler"
	"task_udevs/config"
	"task_udevs/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	r.Use(customCORSMiddleware())

	h := handler.NewHandler(cfg, strg)

	r.POST("/login-user", h.AuthUser)
	r.POST("/login-curier", h.AuthCurier)
	r.Use(MaxAllowed(5000))

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetByIDProduct)
	r.GET("/product", h.GetListProduct)
	// r.PUT("/product/:id", h.UpdateProduct)
	// r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/addition-product", h.CreateAdditionProduct)
	r.GET("/addition-product/:id", h.GetAdditionProductByID)

	r.GET("/history-user/:id", h.DeserializeUser(), h.GetByIDHistoryUser)
	r.GET("/history-user", h.DeserializeUser(), h.GetListHistoryUser)

	r.POST("/history-curier", h.DeserializeCurier(), h.CreateHistoryCurier)
	r.GET("/history-curier/:id", h.DeserializeCurier(), h.GetByIDHistoryCurier)
	r.GET("history-curier", h.DeserializeCurier(), h.GetListHistoryCurier)

	r.POST("/order", h.DeserializeUser(), h.CreateOrder)
	r.GET("/order/:id", h.DeserializeCurier(), h.GetByIDOrder)
	r.GET("/order/user/:id", h.DeserializeUser(), h.GetByIDOrder)
	r.GET("/order", h.DeserializeCurier(), h.GetListOrder)
	r.PUT("/order/:id", h.DeserializeCurier(), h.UpdateOrder)

	r.POST("/cart", h.DeserializeUser(), h.CreateCart)
	r.GET("/cart/:id", h.DeserializeUser(), h.GetByIDCart)
	r.DELETE("/cart/:id", h.DeserializeUser(), h.DeleteCart)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}
