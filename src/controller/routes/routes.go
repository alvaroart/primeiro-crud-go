package routes

import (
	controller "github.com/alvaroart/primeiro-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByID/:userID", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.GET("/updateUser/userId", controller.UpdateUserByID)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
