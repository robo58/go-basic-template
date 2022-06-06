package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	models "go-oauth/data/models/users"
	persistence "go-oauth/data/repositories"
	"go-oauth/helpers/crypto"
	"go-oauth/helpers/http_error"
	"log"
	"net/http"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

// GetUserById godoc
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} users.User
// @Router /api/users/{id} [get]
// @Security Authorization Token
func GetUserById(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Param("id")
	if user, err := s.Get(id); err != nil {
		http_error.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers godoc
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param username query string false "Username"
// @Param firstname query string false "FirstName"
// @Param lastname query string false "LastName"
// @Success 200 {array} []users.User
// @Router /api/users [get]
// @Security Authorization Token
func GetUsers(c *gin.Context) {
	s := persistence.GetUserRepository()
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.Query(&q); err != nil {
		http_error.NewError(c, http.StatusNotFound, errors.New("users not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Username:  userInput.Username,
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Password:      crypto.HashAndSalt([]byte(userInput.Password)),
		Role:      models.UserRole{RoleName: userInput.Role},
	}
	if err := s.Add(&user); err != nil {
		http_error.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

func UpdateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		http_error.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		user.Username = userInput.Username
		user.LastName = userInput.LastName
		user.FirstName = userInput.FirstName
		user.Password = crypto.HashAndSalt([]byte(userInput.Password))
		user.Role = models.UserRole{RoleName: userInput.Role}
		if err := s.Update(user); err != nil {
			http_error.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

func DeleteUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		http_error.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if err := s.Delete(user); err != nil {
			http_error.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
