package main

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	"github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var sc = securecookie.New([]byte("very-secret"), []byte("a-lot-secret-yay"))

func main() {

	const CookieName = "data"

	e := echo.New()

	e.GET("/index", func(ctx echo.Context) error {
		data, err := getCookie(ctx, CookieName)
		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
			return err
		}

		if data == nil {
			data = M{"message": "Hello", "ID": gubrak.RandomString(32)}

			err = setCookie(ctx, CookieName, data)
			if err != nil {
				return err
			}
		}

		return ctx.JSON(http.StatusOK, data)

	})

	e.Logger.Fatal(e.Start(":9000"))

}

func setCookie(c echo.Context, name string, data M) error {
	encoded, err := sc.Encode(name, data)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    encoded,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(c.Response(), cookie)

	return nil
}

func getCookie(c echo.Context, name string) (M, error) {

	cookie, err := c.Request().Cookie(name)

	if err == nil {
		data := M{}
		if err = sc.Decode(name, cookie.Value, &data); err == nil {
			return data, nil
		}
	}

	return nil, err
}
