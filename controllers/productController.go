package controllers

import (
	"html/template"
	"log"
	"net/http"
	"store/models"
	"strconv"
)

// encapsulates and renders all HTML templates from the templates folder.
var temp = template.Must(template.ParseGlob("templates/*.html"))

// GoLang requests by default need a responseWriter with a header map and a request object.
func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		priceConversion, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Price conversion error")
		}

		amountConversion, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Amount conversion error")
		}

		models.CreateProduct(r.FormValue("name"), r.FormValue("description"), priceConversion, amountConversion)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idConversion, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Amount conversion error")
		}

		priceConversion, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Price conversion error")
		}

		amountConversion, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Amount conversion error")
		}

		models.UpdateProduct(idConversion, r.FormValue("name"), r.FormValue("description"), priceConversion, amountConversion)
	}
	http.Redirect(w, r, "/", 301)
}
