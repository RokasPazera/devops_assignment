package handlers

import (
    "backend/database"
    "backend/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    "log"
)

// GET /api/goals
func GetGoals(c *gin.Context) {
    var goals []models.Goal
    
    // Fetch goals ordered by target percentage
    if err := database.DB.Order("target_percentage").Find(&goals).Error; err != nil {
        c.JSON(http.StatusInternalServerError,
            models.NewErrorResponse("Unknown server error while fetching goals"))
        return
    }
    c.JSON(http.StatusOK, goals)
}

// POST /api/goals
func AddGoal(c *gin.Context) {
    var goal models.Goal
    
    // Try to bind the JSON body to the Goal struct
    if err := c.ShouldBindJSON(&goal); err != nil {
        c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid JSON"))
        return
    }

    // INTENTIONAL VULNERABILITY: raw SQL string concatenation. THIS IS SOMETHING YOU HAVE TO FIX!!!
    query := "INSERT INTO goals (name, target_percentage) VALUES ('" +
        goal.Name + "', " + strconv.Itoa(goal.TargetPercentage) + ")"

    if _, err := database.SQLDB.Exec(query); err != nil {
        log.Printf("DB ERROR: %v\nQuery: %s\n", err, query)
        c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Database error"))
        return
    }

    c.JSON(http.StatusCreated, goal)
}
