package controllers

import (
	"errors"
	"fmt"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/helpers"
	"github.com/DaffaAhmadSM/CBTapp/internal/http/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func (ctr *MainController) LoginView(c echo.Context) error {
	// set header bearer)
	//c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	return c.Render(200, "login", nil)
}

type Siswa struct {
	Id   int
	Name string
}

func (ctr *MainController) GetTokenSiswa(c echo.Context) error {
	var user models.User
	ctr.Db.Statement.RaiseErrorOnNotFound = true
	data := ctr.Db.Where("name = ? AND password = ?", c.FormValue("username"), c.FormValue("password")).First(&user).Error

	if errors.Is(gorm.ErrRecordNotFound, data) {
		return c.Render(200, "invalid_credential_err", nil)
	}
	claims := &helpers.Siswa{
		Id:     1,
		Name:   "Budi",
		RoleId: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	http.SetCookie(c.Response(), &http.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().Header().Set("HX-Redirect", "/home/siswa")
	return c.JSON(http.StatusOK, "Login success")
}
