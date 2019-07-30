package controller

import (
	"fmt"
	"net/http"
	"shopstore/src/redis"
	"strconv"
)

//InsertProduct -
func InsertProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	quantity, _ := strconv.Atoi(r.Form["quantity"][0])
	price, _ := strconv.ParseFloat(r.Form["price"][0], 64)

	product := redis.Product{
		Code:        r.Form["code"][0],
		Name:        r.Form["name"][0],
		Description: r.Form["description"][0],
		Image:       r.Form["image"][0],
		Quantity:    quantity,
		Price:       price,
	}

	redis.InsertProduct(product)
}

//GetAllProducts -
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
