package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/juanvaleriand/go-test/common"
	"github.com/juanvaleriand/go-test/models"
	"github.com/juanvaleriand/go-test/utils"
	"github.com/juanvaleriand/go-test/validators"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type ApiResponse struct {
	Data []User `json:"data"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://reqres.in/api/users?per_page=12", nil)

	if err != nil {
		utils.Respond(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		utils.Respond(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, err.Error())
		return
	}

	var user models.User
	for _, data := range apiResponse.Data {
		var exists bool

		models.DB.Model(&user).Select("count(*) > 0").Where("email = ?", data.Email).Find(&exists)

		if exists {
			continue
		}

		user := &models.User{
			Email:      data.Email,
			First_name: data.FirstName,
			Last_name:  data.LastName,
			Avatar:     data.Avatar,
		}

		models.DB.Create(user)
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	var totalRecords int64
	models.DB.Model(&user).Count(&totalRecords)

	offset := (page - 1) * limit
	var data []models.User

	models.DB.Order("id DESC").Offset(offset).Limit(limit).Find(&data)

	pagination := common.Pagination{
		Page:  page,
		Limit: limit,
		Total: totalRecords,
		Data:  data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pagination)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.Respond(w, http.StatusNotFound, "User not found!")
		return
	}

	data := common.ShowData{
		Data: user,
	}

	json.NewEncoder(w).Encode(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input validators.UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.Respond(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	_, errEmail := mail.ParseAddress(input.Email)

	if errEmail != nil {
		utils.Respond(w, http.StatusBadRequest, "Email is not valid")
		return
	}

	var users models.User

	if isEmailExisting := models.DB.Model(&users).Where("email = ?", input.Email).First(&users).Error; isEmailExisting == nil {
		utils.Respond(w, http.StatusBadRequest, "Email already exists!")
		return
	}

	user := &models.User{
		Email:      input.Email,
		First_name: input.First_name,
		Last_name:  input.Last_name,
		Avatar:     input.Avatar,
	}

	models.DB.Create(user)

	utils.Respond(w, http.StatusCreated, "Data created successfully!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.Respond(w, http.StatusNotFound, "User not found!")
		return
	}

	var input validators.UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.Respond(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	_, errEmail := mail.ParseAddress(input.Email)

	if errEmail != nil {
		utils.Respond(w, http.StatusBadRequest, "Email is not valid")
		return
	}

	var findUserByEmail models.User

	models.DB.Where("email = ?", input.Email).First(&findUserByEmail)

	if user.ID != findUserByEmail.ID && input.Email == findUserByEmail.Email {
		utils.Respond(w, http.StatusBadRequest, "Email already exists!")
		return
	}

	user.Email = input.Email
	user.First_name = input.First_name
	user.Last_name = input.Last_name
	user.Avatar = input.Avatar

	models.DB.Save(&user)

	utils.Respond(w, http.StatusOK, "Data updated successfully!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		utils.Respond(w, http.StatusNotFound, "User not found")
		return
	}

	if err := user.Delete(models.DB); err != nil {
		utils.Respond(w, http.StatusInternalServerError, "Test")
		return
	}

	utils.Respond(w, http.StatusOK, "Data deleted successfully!")
}
