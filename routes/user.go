package routes

import (
	"server/controllers"
	"server/models"
)

// Get a UserController instance
var uc = controllers.NewUserController(models.GetSession())

func init() {
	R.POST("/user", uc.CreateUser)

	R.GET("/user/:id", uc.GetUser)

	R.DELETE("/user/:id", uc.RemoveUser)

	R.PUT("/user/avatar/:id", uc.UpdateAvatar)
}
