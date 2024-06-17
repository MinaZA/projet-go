package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	Quantity    int
	Active      bool
}

func addProduct() {
	var title, description string
	var price float64
	var quantity int

	fmt.Print("Enter title: ")
	fmt.Scan(&title)
	fmt.Print("Enter description: ")
	fmt.Scan(&description)
	fmt.Print("Enter price: ")
	fmt.Scan(&price)
	fmt.Print("Enter quantity: ")
	fmt.Scan(&quantity)

	_, err := db.Exec(`INSERT INTO products (title, description, price, quantity) VALUES (?, ?, ?, ?)`, title, description, price, quantity)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product added successfully!")
}

func viewProducts() {
	rows, err := db.Query(`SELECT id, title, description, price, quantity, active FROM products WHERE active = TRUE`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Quantity, &product.Active)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Title: %s, Description: %s, Price: %.2f, Quantity: %d\n", product.ID, product.Title, product.Description, product.Price, product.Quantity)
	}
}

func modifyProduct() {
	var id int
	fmt.Print("Enter product ID to modify: ")
	fmt.Scan(&id)

	var title, description string
	var price float64
	var quantity int

	fmt.Print("Enter new title: ")
	fmt.Scan(&title)
	fmt.Print("Enter new description: ")
	fmt.Scan(&description)
	fmt.Print("Enter new price: ")
	fmt.Scan(&price)
	fmt.Print("Enter new quantity: ")
	fmt.Scan(&quantity)

	_, err := db.Exec(`UPDATE products SET title = ?, description = ?, price = ?, quantity = ? WHERE id = ?`, title, description, price, quantity, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product modified successfully!")
}

func deleteProduct() {
	var id int
	fmt.Print("Enter product ID to delete: ")
	fmt.Scan(&id)

	_, err := db.Exec(`UPDATE products SET active = FALSE WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product deleted successfully!")
}

func exportProductsToCSV() {
	rows, err := db.Query(`SELECT id, title, description, price, quantity, active FROM products`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Description", "Price", "Quantity", "Active"})

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Quantity, &product.Active)
		if err != nil {
			log.Fatal(err)
		}
		writer.Write([]string{fmt.Sprintf("%d", product.ID), product.Title, product.Description, fmt.Sprintf("%.2f", product.Price), fmt.Sprintf("%d", product.Quantity), fmt.Sprintf("%t", product.Active)})
	}
	fmt.Println("Products exported to CSV successfully!")
}
