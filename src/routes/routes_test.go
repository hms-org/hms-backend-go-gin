package routes_test

import (
	"hms-backend-go-gin/src/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := routes.SetupRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	expectedBody := `{"message":"Hello World"}`
	assert.JSONEq(t, expectedBody, resp.Body.String())
}
