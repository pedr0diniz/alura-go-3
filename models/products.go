package models

import "store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetAllProducts() []Product {

	// products := []Product{
	// 	{Name: "T-Shirt", Description: "Blue, pretty", Price: 39.0, Amount: 5},
	// 	{Name: "Sneakers", Description: "Comfortable", Price: 89.0, Amount: 3},
	// 	{Name: "Headphones", Description: "Really good", Price: 59.0, Amount: 2},
	// 	{Name: "New product", Description: "Really cool", Price: 1.99, Amount: 1},
	// }

	// connects to the db
	db := db.DatabaseConnection()

	// selects all products from the db table
	selectAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	tempProduct := Product{}
	products := []Product{}

	// checks all products individually
	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		// assigns the value for each column to a variable in the for loop
		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		// fills the temp object with the scanned values
		tempProduct.Id = id
		tempProduct.Name = name
		tempProduct.Description = description
		tempProduct.Price = price
		tempProduct.Amount = amount

		// adds the temp object to the slice
		products = append(products, tempProduct)
	}
	defer db.Close()
	return products
}

func CreateProduct(name string, description string, price float64, amount int) {
	db := db.DatabaseConnection()

	insertProduct, err := db.Prepare("insert into products(name, description, price, amount) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DatabaseConnection()

	deleteProduct, err := db.Prepare("delete from products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.DatabaseConnection()

	productFromDb, err := db.Query("select * from products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productFromDb.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productFromDb.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Amount = amount
	}

	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name string, description string, price float64, amount int) {
	db := db.DatabaseConnection()
	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
}
