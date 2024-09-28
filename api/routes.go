package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tokha04/swe-farmer-market-system/db/sqlc"
)

func UserRoutes(incomingRoutes *gin.Engine, q *db.Queries) {
	incomingRoutes.POST("/signup", Signup(q))
	incomingRoutes.POST("/login", Login(q))
}

func FarmerRoutes(incomingRoutes *gin.Engine, q *db.Queries) {
	incomingRoutes.POST("/farms", CreateFarm(q))
	incomingRoutes.GET("/farms/:id", GetFarm(q))
	incomingRoutes.GET("/farms", ListFarms(q))
	incomingRoutes.PATCH("/farms/:id", UpdateFarm(q))
	incomingRoutes.DELETE("/farms/:id", DeleteFarm(q))
}
