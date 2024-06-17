package main

import (
	"fmt"
)

func main() {
	initDB()
	defer db.Close()

	createTables()

	for {
		fmt.Println("1. Add Product")
		fmt.Println("2. View Products")
		fmt.Println("3. Modify Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Export Products to CSV")
		fmt.Println("6. Add Client")
		fmt.Println("7. View Clients")
		fmt.Println("8. Modify Client")
		fmt.Println("9. Export Clients to CSV")
		fmt.Println("10. Place Order")
		fmt.Println("11. Export Orders")
		fmt.Println("12. Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addProduct()
		case 2:
			viewProducts()
		case 3:
			modifyProduct()
		case 4:
			deleteProduct()
		case 5:
			exportProductsToCSV()
		case 6:
			addClient()
		case 7:
			viewClients()
		case 8:
			modifyClient()
		case 9:
			exportClientsToCSV()
		case 10:
			placeOrder()
		case 11:
			exportOrders()
		case 12:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
