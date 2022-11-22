package main

// @title User API documentation
// @version 1.0.0
// @host localhost:3000
// @BasePath /user

import (
	"github.com/RustedBl4de/cross-cloud/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/RustedBl4de/cross-cloud/cmd/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	db.Init(dbUrl)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	r.Run(port)
}
