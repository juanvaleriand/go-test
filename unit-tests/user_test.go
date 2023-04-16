package unit_tests

import (
	"net/http"
	"testing"

	"github.com/juanvaleriand/go-test/controllers"
	"github.com/juanvaleriand/go-test/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	writer := controllers.MakeRequest("GET", "/api/v1/users?page=1&limit=10", nil, false)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestFindUser(t *testing.T) {
	writer := controllers.MakeRequest("GET", "/api/v1/user/1", nil, false)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestAddUser(t *testing.T) {
	payload := models.User{
		Email:      "juandelima@email.com",
		First_name: "Juan",
		Last_name:  "Delima",
		Avatar:     "https://www.anakui.com/wp-content/uploads/2023/01/karakter-anime-ter-cool-1.webp",
	}

	writer := controllers.MakeRequest("POST", "/api/v1/user", payload, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestUpdateUser(t *testing.T) {
	payload := models.User{
		Email:      "george.bluth@reqres.in",
		First_name: "George",
		Last_name:  "Bluth Update From Unit Testing",
		Avatar:     "https://reqres.in/img/faces/1-image.jpg",
	}

	writer := controllers.MakeRequest("PUT", "/api/v1/user/1", payload, false)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestDeleteUser(t *testing.T) {
	writer := controllers.MakeRequest("DELETE", "/api/v1/user/1", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}
