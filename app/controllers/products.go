package controllers

import (
	"log"
	"net/http"

	"github.com/netqyq/deer-api/app/models"
	"github.com/revel/revel"
)

// you must wirite this way.
// you can't use
//
// type Products struct {
// 	GorpController
// }
//
type Products struct {
	App
}

func (c Products) Create() revel.Result {
	name := c.Params.Get("name")
	price := c.Params.Get("price")
	code := c.Params.Get("code")

	newProduct := models.Product{Name: name, Price: price, Code: code}
	err := c.Txn.Insert(&newProduct)

	checkErr(err, "insert err")
	log.Println(c.Txn)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(newProduct)
}

func (c Products) Show(id int) revel.Result {
	product, err := c.Txn.Select(models.Product{}, `select * from Product where id = ?`, id)
	if err != nil {
		log.Println(err)
	}
	if len(product) == 0 {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("not found")
	}
	return c.RenderJSON(product)
}

func (c Products) Index() revel.Result {
	// var products []*models.Product
	products, err := c.Txn.Select(models.Product{}, `select * from Product`)
	checkErr(err, "index select err")
	return c.RenderJSON(products)
}

func (c Products) Update(id int) revel.Result {
	name := c.Params.Get("name")
	price := c.Params.Get("price")
	code := c.Params.Get("code")
	newProduct := &models.Product{Id: id, Name: name, Price: price, Code: code}

	_, err := c.Txn.Update(newProduct)
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(err)
	}
	return c.RenderJSON("update success")
}

func (c Products) Destroy(id int) revel.Result {
	newProduct := &models.Product{Id: id}
	_, err := c.Txn.Delete(newProduct)
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(err)
	}
	return c.RenderJSON("delete success")
}
