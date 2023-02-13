package controllers

import (
	"ambassador/src/database"
	"ambassador/src/middlewares"
	"ambassador/src/models"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	user := models.User{
		FirstName:    data["firstname"],
		LastName:     data["lastname"],
		Email:        data["email"],
		IsAmbassador: strings.Contains(c.Path(), "/api/ambassador"),
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	isAmbassador := strings.Contains(c.Path(), "/api/ambassador")
	var scope string
	if isAmbassador {
		scope = "ambassador"
	} else {
		scope = "admin"
	}

	if !isAmbassador && user.IsAmbassador {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, err := middlewares.GenerateJWT(user.Id, scope)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	// The cookie is HTTPOnly because its only purpose is to be sent from the frontend to the backend.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// The only way to remove a cookie is to set it in the past.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	id, _ := middlewares.GetUserId(c)

	user := models.User{
		FirstName: data["firstname"],
		LastName:  data["lastname"],
		Email:     data["email"],
	}
	user.Id = id
	database.DB.Model(&user).Updates(&user)
	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	id, _ := middlewares.GetUserId(c)
	user := models.User{}
	user.Id = id
	user.SetPassword(data["password"])
	database.DB.Model(&user).Updates(&user)
	return c.JSON(user)
}
