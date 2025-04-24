package testing

import (
	"DummyMultifinance/domain"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestE2E_UserFlow(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	// Simulating a new user registration
	newUser := &domain.User{Username: "testuser", Password: "password"}
	url := os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + "/register"

	reqBody, _ := json.Marshal(newUser)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Simulate login for the created user
	loginReq := map[string]string{"username": "testuser", "password": "password"}
	loginReqBody, _ := json.Marshal(loginReq)
	loginReq2, err := http.NewRequest("POST", os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT")+"/login", bytes.NewBuffer(loginReqBody))
	if err != nil {
		t.Fatalf("Error creating login request: %v", err)
	}

	resp2, err := client.Do(loginReq2)
	if err != nil {
		t.Fatalf("Error sending login request: %v", err)
	}
	defer resp2.Body.Close()

	assert.Equal(t, http.StatusOK, resp2.StatusCode)
}
