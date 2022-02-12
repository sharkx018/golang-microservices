package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUser_NO(t *testing.T) {
	//Initialization:

	//Execution:
	user, err  := GetUser(0)

	//Validation:
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "we are expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found",err.Code)
	assert.EqualValues(t,"User 0 not found", err.Message)

}

func TestGetUserValidUser(t *testing.T) {

	user, err := GetUser(123)

	assert.Nil(t, err, "We are not expecting any error")
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Mukul", user.FirstName)
	assert.EqualValues(t, "Verma", user.LastName)
	assert.EqualValues(t, "aryanv018@gmail.com", user.Email)

}