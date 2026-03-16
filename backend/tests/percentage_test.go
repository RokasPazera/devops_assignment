package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "backend/database"
    "backend/handlers"
    "backend/models"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func setupPercentageRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/api/percentage", handlers.GetPercentage)
    r.POST("/api/percentage/:percentage", handlers.AddPercentage)
    return r
}

func resetPercentageDB() {
    database.ConnectTest()
    database.DB.Exec("DELETE FROM percentages")
}

func TestAddAndGetPercentage(t *testing.T) {
    resetPercentageDB()
    router := setupPercentageRouter()

    // POST percentage
    req, _ := http.NewRequest("POST", "/api/percentage/42", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)

    // GET latest percentage
    req, _ = http.NewRequest("GET", "/api/percentage", nil)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)

    var resp models.PercentageResponse
    if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
        t.Fatalf("failed to unmarshal response: %v", err)
    }
    assert.Equal(t, 42, resp.Percentage)
}

// Optional: test invalid percentage
func TestAddInvalidPercentage(t *testing.T) {
    resetPercentageDB()
    router := setupPercentageRouter()

    req, _ := http.NewRequest("POST", "/api/percentage/150", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
}
