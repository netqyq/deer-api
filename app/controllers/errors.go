package controllers

import (
	"errors"
	"log"
)

var (
	errAuthHeaderNotFound = errors.New("authorization header not found")
	errInvalidTokenFormat = errors.New("token format is invalid")
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}
