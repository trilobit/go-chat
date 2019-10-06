package main

import (
	"testing"

	"github.com/trilobit/go-chat/test"
	"github.com/trilobit/go-chat/test/repositories"
)

func TestAccountService_Register_OK(t *testing.T) {
	service, err := test.InitService()
	if err != nil {
		t.Errorf("TestAccountService_Register failed")
	}

	user, err := service.Register(repositories.TestEmail, repositories.TestPassword)
	if err != nil {
		t.Errorf("TestAccountService_Register failed: %v", err)
	}

	if user.Email != repositories.TestEmail {
		t.Errorf("TestAccountService_Register failed: incorrect result data %v", user)
	}
}

func TestAccountService_Authorize_OK(t *testing.T) {
	service, err := test.InitService()
	if err != nil {
		t.Errorf("TestAccountService_Authorize failed")
	}

	user, err := service.Authorize(repositories.TestEmail, repositories.TestPassword)
	if err != nil {
		t.Errorf("TestAccountService_Authorize failed: %v", err)
	}

	if user.Email != repositories.TestEmail {
		t.Errorf("TestAccountService_Authorize failed: incorrect result data %v", user)
	}
}
