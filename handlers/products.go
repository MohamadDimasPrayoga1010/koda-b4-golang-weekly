package handlers

import "time"


type Menu struct {
	ID int
    Name  string
    Price int
}

type Order struct {
    Item     Menu
    Quantity int
}

type Invoice struct {
    InvoiceNumber string
    Date          time.Time
    Orders        []Order
    Total         int
}

var Orders []Order
var History []Invoice
