package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Order struct {
	ID           int
	ClientID     int
	ProductID    int
	Quantity     int
	Price        float64
	PurchaseDate time.Time
}

func placeOrder() {
	var clientID, productID, quantity int
	fmt.Print("Enter client ID: ")
	fmt.Scan(&clientID)
	fmt.Print("Enter product ID: ")
	fmt.Scan(&productID)
	fmt.Print("Enter quantity: ")
	fmt.Scan(&quantity)

	var product Product
	err := db.QueryRow(`SELECT id, price FROM products WHERE id = ? AND active = TRUE`, productID).Scan(&product.ID, &product.Price)
	if err != nil {
		log.Fatal(err)
	}

	totalPrice := product.Price * float64(quantity)
	purchaseDate := time.Now()

	_, err = db.Exec(`INSERT INTO orders (client_id, product_id, quantity, price, purchase_date) VALUES (?, ?, ?, ?, ?)`, clientID, productID, quantity, totalPrice, purchaseDate)
	if err != nil {
		log.Fatal(err)
	}

	sendOrderEmail(clientID, productID, quantity, totalPrice, purchaseDate)
	generateOrderPDF(clientID, productID, quantity, totalPrice, purchaseDate)

	fmt.Println("Order placed successfully!")
}

func sendOrderEmail(clientID, productID, quantity int, totalPrice float64, purchaseDate time.Time) {
	var client Client
	err := db.QueryRow(`SELECT email FROM clients WHERE id = ?`, clientID).Scan(&client.Email)
	if err != nil {
		log.Fatal(err)
	}

	from := "your_email@example.com"
	pass := "your_password"
	to := client.Email
	msg := "Subject: Order Confirmation\n\n" +
		fmt.Sprintf("Thank you for your order!\n\nProduct ID: %d\nQuantity: %d\nTotal Price: %.2f\nPurchase Date: %s\n",
			productID, quantity, totalPrice, purchaseDate.Format(time.RFC1123))

	err = smtp.SendMail("smtp.example.com:587", smtp.PlainAuth("", from, pass, "smtp.example.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Confirmation email sent successfully!")
}

func generateOrderPDF(clientID, productID, quantity int, totalPrice float64, purchaseDate time.Time) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Order Confirmation")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Client ID: %d", clientID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Product ID: %d", productID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Quantity: %d", quantity))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Total Price: %.2f", totalPrice))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Purchase Date: %s", purchaseDate.Format(time.RFC1123)))
	err := pdf.OutputFileAndClose("order.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Order PDF generated successfully!")
}

func exportOrders() {
	rows, err := db.Query(`SELECT o.id, o.client_id, o.product_id, o.quantity, o.price, o.purchase_date, c.first_name, c.last_name, c.email, p.title FROM orders o JOIN clients c ON o.client_id = c.id JOIN products p ON o.product_id = p.id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	file, err := os.Create("orders.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Order ID", "Client ID", "Client Name", "Client Email", "Product ID", "Product Title", "Quantity", "Total Price", "Purchase Date"})

	for rows.Next() {
		var order Order
		var clientName, clientEmail, productTitle string
		err := rows.Scan(&order.ID, &order.ClientID, &order.ProductID, &order.Quantity, &order.Price, &order.PurchaseDate, &clientName, &clientEmail, &productTitle)
		if err != nil {
			log.Fatal(err)
		}
		writer.Write([]string{fmt.Sprintf("%d", order.ID), fmt.Sprintf("%d", order.ClientID), clientName, clientEmail, fmt.Sprintf("%d", order.ProductID), productTitle, fmt.Sprintf("%d", order.Quantity), fmt.Sprintf("%.2f", order.Price), order.PurchaseDate.Format(time.RFC1123)})
	}
	fmt.Println("Orders exported to CSV successfully!")
}
