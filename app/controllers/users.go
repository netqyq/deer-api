package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/netqyq/deer-api/app/models"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

type Users struct {
	GorpController
}

var hmacSecret = []byte{97, 48, 97, 50, 97, 98, 105, 49, 99, 102, 83, 53, 57, 98, 52, 54, 97, 102, 99, 12, 12, 13, 56, 34, 23, 16, 78, 67, 54, 34, 32, 21}

// Register create a user and returns token to client.
// params: email, password
// result: token with user.id stores in `sub` field.
func (c Users) Register() revel.Result {
	// create user use, email, password
	// return token to user
	email := c.Params.Get("email")
	password := c.Params.Get("password")

	//
	if email == "" || password == "" {
		// this is not json
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("params is not valid.")
	}

	// check if the email have already exists in DB
	results, err := Dbm.Select(models.User{}, `select * from User where Email = ?`, email)
	if err != nil {
		log.Println(err)
	}
	var users []*models.User
	for _, r := range results {
		u := r.(*models.User)
		users = append(users, u)
	}
	if users != nil {
		c.Response.Status = http.StatusConflict
		return c.RenderJSON("user have already exists.")
	}

	// Crete user struct
	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)

	token := encodeToken(email)

	newUser := &models.User{0, "Demo User", email, password, bcryptPassword, []byte(token)}

	// Validate user struct
	newUser.Validate(c.Validation)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors[0].Message)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("bad email address.")
	}

	// Save user info to DB
	if err := Dbm.Insert(newUser); err != nil {
		panic(err)
	}

	msg := make(map[string]string)
	msg["email"] = email
	msg["result"] = "user created"
	msg["token"] = token
	return c.RenderJSON(msg)
}

// Login authticate via email and password, if the user is valid,
// returns the token to client.
func (c Users) Login() revel.Result {
	log.Println("login")
	email := c.Params.Get("email")
	password := c.Params.Get("password")

	user, err := getUser(email)
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON("invalid email or password")
	}

	// get token
	tokenString := encodeToken(email)

	msg := make(map[string]string)
	msg["result"] = "login success"
	msg["token"] = tokenString
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(msg)
}

func getUser(email string) (*models.User, error) {
	log.Println("before select", email)
	users, err := Dbm.Select(models.User{}, `select * from User where Email = ?`, email)
	log.Println("after select 1")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("after select 2")
	if len(users) == 0 {
		return nil, errors.New("user not found")
	}
	return users[0].(*models.User), nil
}

func encodeToken(email string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	fmt.Println(tokenString, err)

	return tokenString
}

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println("email and nbf:", claims["email"], claims["nbf"])
	} else {
		log.Println(err)
		return nil, err
	}
	return claims, nil
	// return claims["email"].(string), claims["nbf"].(string)
}
