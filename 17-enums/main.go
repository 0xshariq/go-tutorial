package main

import "fmt"

type OrderStatus string

const (
	Received   OrderStatus = "received"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Delivered  OrderStatus = "delivered"
	Canceled   OrderStatus = "canceled"
	Confirmed  OrderStatus = "confirmed"
)

func changeOrdeerStatus(status OrderStatus) {
	fmt.Println("Changing the order status to", status)
}

func main() {
	changeOrdeerStatus(Received)
}
