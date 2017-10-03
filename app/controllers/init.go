package controllers

import (
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

func AddLog(c *revel.Controller) revel.Result {
	log.Println("InterceptFunc Test.")
	return nil
}

// Authenticate is and method will be called before any authenticate needed action.
// In order to valid the user.
func Authenticate(c *revel.Controller) revel.Result {
	log.Println("Authenticate!")
	log.Println(c.Params)

	tokenString, err := getTokenString(c)
	if err != nil {
		log.Println("get token string failed")
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("get token string failed")
	}

	var claims jwt.MapClaims
	claims, err = decodeToken(tokenString)
	if err != nil {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}
	log.Println("claims decode:", claims)
	log.Println(claims["email"])
	email, found := claims["email"]
	if !found {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("email not found in db")
	}
	log.Println("email found:", email)
	_, err = getUser(email.(string))
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusUnauthorized
		return c.RenderJSON("auth failed")
	}
	log.Println("auth token success")
	return nil
}

func getTokenString(c *revel.Controller) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println(errAuthHeaderNotFound)
		return "", errAuthHeaderNotFound
	}

	tokenSlice := strings.Split(authHeader, " ")
	if len(tokenSlice) != 2 {
		return "", errInvalidTokenFormat
	}
	tokenString = tokenSlice[1]
	return tokenString, nil

}

func init() {
	revel.OnAppStart(InitDB1)
	// test
	revel.InterceptFunc(AddLog, revel.BEFORE, &App{})

	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptFunc(Authenticate, revel.BEFORE, &App{})
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
