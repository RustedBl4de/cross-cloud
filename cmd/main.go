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

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} model.User
// @Failure 404 {object} object
// @Router / [get]
func GetUser(c *gin.Context) {
	return
}

// CreateUser ... Create User
// @Summary Create new user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body model.User true "User Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router / [post]
func CreateUser(c *gin.Context) {
	return
}

// GetUserByID ... Get the user by id
// @Summary Get one user
// @Description get user by ID
// @Tags Users
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400,404 {object} object
// @Router /{id} [get]
func GetUserByID(c *gin.Context) {
	return
}

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	db.Init(dbUrl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(port)
}
