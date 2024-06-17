// package database

// import (
//     "database/sql"
//     "fmt"
//     _ "github.com/go-sql-driver/mysql"
// )

// type DB struct {
//     *sql.DB
// }

// func ConnectDB() (*DB, error) {
//     dsn := "user:password@tcp(localhost:3306)/shop"
//     db, err := sql.Open("mysql", dsn)
//     if err != nil {
//         return nil, fmt.Errorf("erreur lors de la connexion à la base de données : %v", err)
//     }
//     return &DB{db}, nil
// }

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func createTables() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
        id INT AUTO_INCREMENT,
        title VARCHAR(255),
        description TEXT,
        price FLOAT,
        quantity INT,
        active BOOL DEFAULT TRUE,
        PRIMARY KEY (id)
    )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clients (
        id INT AUTO_INCREMENT,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        phone VARCHAR(255),
        address TEXT,
        email VARCHAR(255),
        PRIMARY KEY (id)
    )`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
        id INT AUTO_INCREMENT,
        client_id INT,
        product_id INT,
        quantity INT,
        price FLOAT,
        purchase_date TIMESTAMP,
        PRIMARY KEY (id),
        FOREIGN KEY (client_id) REFERENCES clients(id),
        FOREIGN KEY (product_id) REFERENCES products(id)
    )`)
	if err != nil {
		log.Fatal(err)
	}
}
