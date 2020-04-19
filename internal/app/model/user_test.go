package model_test

import (
	"github.com/isratmir/restapi/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotNil(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u  		func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		}, {
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		}, {
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		}, {
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		}, {
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "pass"

				return u
			},
			isValid: false,
		}, {
			name: "with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted pass"

				return u
			},
			isValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.NoError(t, tt.u().Validate())
			} else {
				assert.Error(t, tt.u().Validate())
			}
		})
	}
}