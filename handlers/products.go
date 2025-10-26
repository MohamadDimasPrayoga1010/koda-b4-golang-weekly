package handlers

import "time"

type UserOrders interface{
    GetID() int
    GetName() string
    GetPrice() int
    CalculateSubtotal(quantity int) int
}


type Menu struct {
	ID int
    Name  string
    Price int
}

func (m Menu) GetID() int{
    return m.ID
}

func (m Menu) GetName() string{
    return m.Name
}

func (m Menu) GetPrice() int{
    return m.Price
}

func (m Menu) CalculateSubtotal(quantity int) int{
    return m.Price * quantity
}

type Order struct {
    Item     UserOrders
    Quantity int
}

func (o Order) GetSubtotal() int{
    return o.Item.CalculateSubtotal(o.Quantity)
}

func (o Order) GetItemName() string{
    return o.Item.GetName()
}

func (o Order) GetItemPrice() int{
    return o.Item.GetPrice()
}

type Invoice struct {
    InvoiceNumber string
    Date          time.Time
    Orders        []Order
    Total         int
}

func (inv Invoice) CalculateTotal() int{
    total := 0
	for _, order := range inv.Orders {
		total += order.GetSubtotal()
	}
	return total
}

func (inv Invoice) GetOrderCount() int {
	return len(inv.Orders)
}

var Orders []Order
var History []Invoice
