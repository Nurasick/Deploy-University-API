package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"university/helpers/jwt"

	"github.com/labstack/echo/v4"
)

// JWTAuth is a middleware function that validates JWT tokens in incoming requests
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization") //getting Authorization header
		if auth == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing Authorization header"}) //checking if header is present
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			auth = "Bearer " + auth
		}

		parts := strings.Split(auth, " ")            //splitting the header to get the token
		if len(parts) != 2 || parts[0] != "Bearer" { //validating the format of the header
			fmt.Println("Invalid Authorization header format:", auth)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format"})
		}
		claims, err := jwt.ValidateToken(parts[1]) //validating the token
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token: " + err.Error()})
		}

		c.Set("userID", claims.UserID) //setting user ID in context
		c.Set("roleID", claims.RoleID)
		return next(c)
	}
}
