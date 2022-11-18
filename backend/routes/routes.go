package routes

import (
	"github.com/MilliGoshant/merci/backend/controllers"
	"github.com/MilliGoshant/merci/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	routes := gin.Default()

	api := routes.Group("/api")

	{
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategory)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)

		api.POST("/products", controllers.CreateProduct)
		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProduct)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)

		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/token", controllers.Token)
		api.GET("/auth/me", middlewares.Auth(), controllers.GetUser)
		api.PUT("/auth/me", middlewares.Auth(), controllers.UpdateUser)
		api.DELETE("/auth/me", controllers.Logout)

		api.POST("/admin", controllers.CreateAdmin)
		api.GET("/admin/:id", controllers.GetAdmin)
		api.GET("/admin", controllers.GetAdmins)
		api.PUT("/admin/:id", controllers.UpdateAdmin)
		api.DELETE("/admin/:id", controllers.DeleteAdmin)

	}

	return routes
}
