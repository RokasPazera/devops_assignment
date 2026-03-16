package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "backend/database"
    "backend/handlers"
    "backend/models"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/api/goals", handlers.GetGoals)
    r.POST("/api/goals", handlers.AddGoal)
    return r
}

// Reset database before each test
func resetDB() {
    database.ConnectTest() // must use in-memory SQLite
    database.DB.Exec("DELETE FROM goals")
}

func TestAddAndGetGoal(t *testing.T) {
    resetDB()
    router := setupRouter()

    // POST a new goal
    goal := models.Goal{Name: "TestGoal", TargetPercentage: 50}
    jsonValue, _ := json.Marshal(goal)
    req, _ := http.NewRequest("POST", "/api/goals", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // GET all goals
    req, _ = http.NewRequest("GET", "/api/goals", nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var goals []models.Goal
    if err := json.Unmarshal(w.Body.Bytes(), &goals); err != nil {
        t.Fatalf("failed to unmarshal response: %v", err)
    }
    assert.Len(t, goals, 1)
    assert.Equal(t, "TestGoal", goals[0].Name)
    assert.Equal(t, 50, goals[0].TargetPercentage)
}
