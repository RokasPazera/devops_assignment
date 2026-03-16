package handlers

import (
    "backend/database"
    "backend/models"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// GET /api/percentage
func GetPercentage(c *gin.Context) {
    var percentage models.Percentage
    
    // Get the latest percentage entry
    if err := database.DB.Order("id desc").First(&percentage).Error; err != nil {
        c.JSON(http.StatusOK, models.NewPercentageResponse(0))
        return
    }

    c.JSON(http.StatusOK, models.NewPercentageResponse(percentage.Value))
}

// POST /api/percentage/:percentage
func AddPercentage(c *gin.Context) {
    // Validate and parse the percentage parameter
    percStr := c.Param("percentage")
    perc, err := strconv.Atoi(percStr)
    if err != nil || perc < 0 || perc > 100 {
        c.JSON(http.StatusBadRequest,
            models.NewErrorResponse("Percentage must be between 0 and 100"))
        return
    }

    // Create and save the new percentage entry in the database
    percentage := models.Percentage{Value: perc}
    if err := database.DB.Create(&percentage).Error; err != nil {
        c.JSON(http.StatusInternalServerError,
            models.NewErrorResponse("Error adding percentage"))
        return
    }

    c.JSON(http.StatusCreated, models.NewPercentageResponse(perc))
}
