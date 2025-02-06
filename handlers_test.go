package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	user := NewUserHandlers(NewInMemoryDataStore[User]())
	req := UserRegisterRequest{
		Username:        "ukasz",
		Email:           "ukasz@bulinski.com",
		Password:        "ukanio",
		PasswordConfirm: "ukanio",
	}
	var buf bytes.Buffer

	// data, err := json.Marshal(req)
	// reader := bytes.NewReader(data)

	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		t.Error("failed to encode struct to json")
	}

	s := httptest.NewServer(handler(user.handleRegister))
	resp, err := http.Post(s.URL, "application/json", &buf)
	if err != nil {
		t.Error("failed to do request")
	}

	if resp.StatusCode != 201 {
		t.Error("status code should be 201")
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	apiResponse := ApiResponse{}

	if err := json.Unmarshal(b, &apiResponse); err != nil {
		t.Error("failed to unmarshal response")
	}

	exp := "User ukasz has been successfully created activation email has been sent to ukasz@bulinski.com"

	if apiResponse.Msg != exp {
		t.Error("response msg does not match expected")

	}
}

func TestRegisterHandlerResponseRecorder(t *testing.T) {
	rr := httptest.NewRecorder()

	r := UserRegisterRequest{
		Username:        "ukasz",
		Email:           "ukasz@bulinski.com",
		Password:        "ukanio",
		PasswordConfirm: "ukanio",
	}

	data, err := json.Marshal(r)
	reader := bytes.NewReader(data)
	if err != nil {
	}

	req, err := http.NewRequest(http.MethodPost, "", reader)

	user := NewUserHandlers(NewInMemoryDataStore[User]())

	if err := user.handleRegister(rr, req); err != nil {
	}

	fmt.Println("Everything is going to be recorder here ->", rr)

}
