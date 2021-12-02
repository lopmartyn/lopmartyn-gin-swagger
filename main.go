package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lopmartyn-gin-swagger/app/controller/employee"
	"lopmartyn-gin-swagger/app/controller/gender"
	"lopmartyn-gin-swagger/app/controller/user"
	"lopmartyn-gin-swagger/config"
	"lopmartyn-gin-swagger/config/database"
	"lopmartyn-gin-swagger/docs"
)

func main() {
	// https://github.com/joho/godotenv
	// get env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// init starter
	goDotENV := config.GetGoDotENV()
	// create connection
	database.CreateDatabaseConnection(goDotENV.PostgresConfig)

	// start gin web server
	gin.SetMode(goDotENV.GinMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.Title = "Golang Assignment API"
	docs.SwaggerInfo.Description = "This is API develop base on opn-boilerplate"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	version := router.Group("/v1")

	// gender
	genderRoute := version.Group("/gender")
	{
		genderRoute.GET("get/all", gender.GetAllGenders)
	}

	// user
	registerRoute := version.Group("/user")
	{
		registerRoute.POST("create", user.Registration)
		registerRoute.GET("get/all", user.GetAllUsersInfo)
	}

	// employee
	employeeRoute := version.Group("/employee")
	{
		employeeRoute.POST("create", employee.InsertEmployee)
		employeeRoute.GET("get/:employeeID", employee.GetEmployeeByID)
		employeeRoute.PUT("update", employee.UpdateEmployee)
		employeeRoute.DELETE("delete/:employeeID", employee.DeleteEmployeeByID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(goDotENV.Port); err != nil {
		panic(err)
	}
}
