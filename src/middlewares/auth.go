package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strconv"
	"strings"
	"time"
)

const secretKey = "toto"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	payload := token.Claims.(*ClaimsWithScope)
	isAmbassador := strings.Contains(c.Path(), "/api/ambassador")

	if (payload.Scope == "admin" && isAmbassador) || (payload.Scope == "ambassador" && !isAmbassador) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}

func GenerateJWT(userId uint, scope string) (string, error) {
	claims := &ClaimsWithScope{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	claims.Subject = strconv.Itoa(int(userId))
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims.Scope = scope
	return token.SignedString([]byte(secretKey))
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	payload := token.Claims.(*ClaimsWithScope)
	id, _ := strconv.Atoi(payload.Subject)
	return uint(id), nil
}
