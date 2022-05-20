package controllers

import (
	"net/http"
	"strconv"

	"github.com/MEVLEPEIZ/Mini-Project/config"
	"github.com/MEVLEPEIZ/Mini-Project/database"
	"github.com/MEVLEPEIZ/Mini-Project/middleware"
	"github.com/MEVLEPEIZ/Mini-Project/models"

	"github.com/labstack/echo/v4"
)

// GetUsersController get all users
func GetUsersController(c echo.Context) error {

	users, err := database.GetAllUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"code":     200,
		"users":    users,
	})

}

// GetUserController get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "bad request",
			"code":     400,
		})
	}

	user, e := database.GetUsersByID(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"code":     200,
		"data":     user,
	})

}

// DeleteUserController delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "bad request",
			"code":     400,
		})
	}

	_, e := database.GetUsersByID(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&models.User{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}

// UpdateUserController update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "bad request",
			"code":     400,
		})
	}

	data := models.User{}
	err = c.Bind(&data)
	if err != nil {
		return err
	}

	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.User{
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
	}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": err.Error(),
			"code":     400,
		})
	}

	user, _ := database.GetUsersByID(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user with id " + strconv.Itoa(id),
		"code":     200,
		"data":     user,
	})
}

// CreateUserController create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"code":     200,
		"user":     user,
	})
}
func LoginsUserController(c echo.Context) error {

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if err := config.DB.Where("email = ? AND password = ? ", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "Failed Login",
			"error":    err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Nama)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "Failed Login",
			"error":    err.Error(),
		})
	}
		usersResponse := models.UsersResponse(user.ID, user.Nama:"Eunike Yolanda", user.Email:"eunikeshn45@gmail.com", user.Kondisi_kulit:"kusam", token:)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"code":     200,
		"user":     usersResponse,
	})
}
func LoginsUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success Login",
		"users":    users,
	})
}
func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := database.GetDetailUsers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}
