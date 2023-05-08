package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dvnggak/miniProject/model"
	"github.com/dvnggak/miniProject/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAdmin(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new request and recorder
	req := httptest.NewRequest(http.MethodGet, "/admins", nil)
	rec := httptest.NewRecorder()

	// Create a new Echo context
	c := e.NewContext(req, rec)

	// Create a new instance of the controller
	ctrl := &Controller{}

	// Call the controller function
	err := ctrl.GetAdmin(c)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert the status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert the response body
	expected := `{"message":"success get all admins","users":[]}`
	assert.Equal(t, expected, rec.Body.String())
}

func TestController_CreateAdmin(t *testing.T) {
	// Initialize the echo context
	e := echo.New()
	admin := model.Admin{Username: "John", Password: "john"}
	reqBody, err := json.Marshal(admin)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/admins/", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Initialize the Controller
	m := &Controller{}

	// Call the CreateAdmin function
	err = m.CreateAdmin(c)

	// Check for any errors
	if err != nil {
		t.Errorf("CreateAdmin returned an error: %v", err)
	}

	// Check the status code of the response
	if rec.Code != http.StatusOK {
		t.Errorf("CreateAdmin returned an unexpected status code: %d", rec.Code)
	}

	// Check the body of the response
	var respBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &respBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	expectedMsg := "success to create admin"
	if respBody["message"] != expectedMsg {
		t.Errorf("CreateAdmin returned an unexpected message: %v", respBody["message"])
	}

	// Clean up test data
	err = service.GetAdminRepository().DeleteAdmin(admin.ID)
	if err != nil {
		t.Errorf("Failed to delete admin: %v", err)
	}
}