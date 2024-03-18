package controllers

import (
	"fmt"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/helpers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (controller *MainController) LoginView(c echo.Context) error {
	// set header bearer)
	//c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	return c.Render(200, "login", nil)
}

func (controller *MainController) GetTokenSiswa(c echo.Context) error {
	hour, _ := strconv.Atoi(os.Getenv("JWT_EXP_HOUR"))
	claims := &helpers.Siswa{
		Id:     1,
		Name:   "Budi",
		RoleId: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(hour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	c.SetCookie(cookie)

	return c.Redirect(302, "/siswa")
}
