package routes

import (
    "backend/handlers"
    "github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
    r.GET("/api/goals", handlers.GetGoals)
    r.POST("/api/goals", handlers.AddGoal)
    r.GET("/api/percentage", handlers.GetPercentage)
    r.POST("/api/percentage/:percentage", handlers.AddPercentage)
}
