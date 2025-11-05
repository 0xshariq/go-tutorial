package main

import (
	"fmt"
	"time"
)

// constructing a struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer       // struct embedding
}

type customer struct {
	name string
	phone string
}

// constructor function for struct (in object oriented programming)
func NewOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// method to change the status valu of struct
func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	myOrder := NewOrder("1", 30.00, "Confirmed")
	// myOrder := order{
	// 	id:     id,
	// 	amount: amount,
	// 	status: status,
	//  customer: customer {
	// 		name : "john",
	//		phone "1234578825",
	//  }
	// }

	myOrder.createdAt = time.Now()
	myOrder.changeStatus("paid")
	fmt.Println(myOrder)

	// if we have to use struct only one time then we can create struct like this
	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
}
