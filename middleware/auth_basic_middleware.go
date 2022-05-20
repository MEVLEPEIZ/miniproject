package middleware

import (
 "github.com/MEVLEPEIZ/Mini-Project/config"
 "github.com/MEVLEPEIZ/Mini-Project/models"

 "github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
 var user models.User
 err := config.DB.Where("email = ? AND password = ? ", username, password).First(&user).Error
 if err != nil {
  return false, err
 }
 return true, nil
}