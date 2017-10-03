package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}
