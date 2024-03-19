package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Siswa struct {
	Id     int
	Name   string
	RoleId int
	jwt.RegisteredClaims
}

func WebIssueToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		cookie, err := ec.Cookie("token")
		if err != nil {
			fmt.Printf("%v test", cookie)
			return ec.JSON(http.StatusUnauthorized, "Unauthorized access")
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &Siswa{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !token.Valid {
			fmt.Println("token invalid")
			return ec.JSON(http.StatusUnauthorized, "Unauthorized access")
		}
		if err != nil {
			fmt.Printf("%v", err)
			return ec.JSON(http.StatusInternalServerError, "Error parsing token")
		}

		return next(ec)
	}
}
