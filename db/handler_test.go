// handler_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *HandlerTestSuite) SetupSuite() {
	// Initialize your router or server
	suite.router = gin.Default()

	// Set up your routes
	suite.router.GET("/users/:id", GetUserHandler)
}

func (suite *HandlerTestSuite) TestGetUserHandler() {
	// Create a request to pass to our handler
	req, _ := http.NewRequest("GET", "/users/1", nil)
	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	// Check the response code
	assert.Equal(suite.T(), http.StatusOK, rec.Code)

	// Check the response body
	expected := `{"id":1,"name":"John Doe"}`
	assert.JSONEq(suite.T(), expected, rec.Body.String())
}

func (suite *HandlerTestSuite) TearDownSuite() {
	// Clean up resources, if any
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
