package entities

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Adam", "adam@gmail.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Adam", user.Name)
	assert.Equal(t, "adam@gmail.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Adam", "adam@gmail.com", "1234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "1234", user.Password)
}