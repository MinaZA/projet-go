// package controllers

// import (
// 	"fmt"
// 	"shopapp/database"
// )

// type ClientController struct {
// 	db *database.DB
// }

// func NewClientController(db *database.DB) *ClientController {
// 	return &ClientController{db: db}
// }

// // Implémentez les méthodes pour ajouter, afficher, modifier et exporter les clients
// // Exemple :
// func (cc *ClientController) AddClient() {
// 	// Demandez les détails du client à l'utilisateur
// 	// Insérez les informations dans la table "clients"
// }

// package main

// import (
// 	"database/sql"
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type Client struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Phone     string
// 	Address   string
// 	Email     string
// }

// func addClient() {
// 	var firstName, lastName, phone, address, email string

// 	fmt.Print("Enter first name: ")
// 	fmt.Scan(&firstName)
// 	fmt.Print("Enter last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Print("Enter phone: ")
// 	fmt.Scan(&phone)
// 	fmt.Print("Enter address: ")
// 	fmt.Scan(&address)
// 	fmt.Print("Enter email: ")
// 	fmt.Scan(&email)

// 	_, err := db.Exec(`INSERT INTO clients (first_name, last_name, phone, address, email) VALUES (?, ?, ?, ?, ?)`, firstName, lastName, phone, address, email)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Client added successfully!")
// }

// func viewClients() {
// 	rows, err := db.Query(`SELECT id, first_name, last_name, phone, address, email FROM clients`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var client Client
// 		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("ID: %d, First Name: %s, Last Name: %s, Phone: %s, Address: %s, Email: %s\n", client.ID, client.FirstName, client.LastName, client.Phone, client.Address, client.Email)
// 	}
// }

// func modifyClient() {
// 	var id int
// 	fmt.Print("Enter client ID to modify: ")
// 	fmt.Scan(&id)

// 	var firstName, lastName, phone, address, email string

// 	fmt.Print("Enter new first name: ")
// 	fmt.Scan(&firstName)
// 	fmt.Print("Enter new last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Print("Enter new phone: ")
// 	fmt.Scan(&phone)
// 	fmt.Print("Enter new address: ")
// 	fmt.Scan(&address)
// 	fmt.Print("Enter new email: ")
// 	fmt.Scan(&email)

// 	_, err := db.Exec(`UPDATE clients SET first_name = ?, last_name = ?, phone = ?, address = ?, email = ? WHERE id = ?`, firstName, lastName, phone, address, email, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Client modified successfully!")
// }

// func exportClientsToCSV() {
// 	rows, err := db.Query(`SELECT id, first_name, last_name, phone, address, email FROM clients`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	file, err := os.Create("clients.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	writer.Write([]string{"ID", "First Name", "Last Name", "Phone", "Address", "Email"})

// 	for rows.Next() {
// 		var client Client
// 		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		writer.Write([]string{fmt.Sprintf("%d", client.ID), client.FirstName, client.LastName, client.Phone, client.Address, client.Email})
// 	}
// 	fmt.Println("Clients exported to CSV successfully!")
// }

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Client struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Address   string
	Email     string
}

func tablesSql() {
	initDB()
	defer db.Close()

	createTables()

	addClient()
	viewClients()
	modifyClient()
	exportClientsToCSV()
}

func addClient() {
	var firstName, lastName, phone, address, email string

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter phone: ")
	fmt.Scan(&phone)
	fmt.Print("Enter address: ")
	fmt.Scan(&address)
	fmt.Print("Enter email: ")
	fmt.Scan(&email)

	_, err := db.Exec(`INSERT INTO clients (first_name, last_name, phone, address, email) VALUES (?, ?, ?, ?, ?)`, firstName, lastName, phone, address, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client added successfully!")
}

func viewClients() {
	rows, err := db.Query(`SELECT id, first_name, last_name, phone, address, email FROM clients`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, First Name: %s, Last Name: %s, Phone: %s, Address: %s, Email: %s\n", client.ID, client.FirstName, client.LastName, client.Phone, client.Address, client.Email)
	}
}

func modifyClient() {
	var id int
	fmt.Print("Enter client ID to modify: ")
	fmt.Scan(&id)

	var firstName, lastName, phone, address, email string

	fmt.Print("Enter new first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter new last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter new phone: ")
	fmt.Scan(&phone)
	fmt.Print("Enter new address: ")
	fmt.Scan(&address)
	fmt.Print("Enter new email: ")
	fmt.Scan(&email)

	_, err := db.Exec(`UPDATE clients SET first_name = ?, last_name = ?, phone = ?, address = ?, email = ? WHERE id = ?`, firstName, lastName, phone, address, email, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client modified successfully!")
}

func exportClientsToCSV() {
	rows, err := db.Query(`SELECT id, first_name, last_name, phone, address, email FROM clients`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	file, err := os.Create("clients.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "First Name", "Last Name", "Phone", "Address", "Email"})

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
		if err != nil {
			log.Fatal(err)
		}
		writer.Write([]string{fmt.Sprintf("%d", client.ID), client.FirstName, client.LastName, client.Phone, client.Address, client.Email})
	}
	fmt.Println("Clients exported to CSV successfully!")
}
