package gorb

import (
	"errors"
	"fmt"
)

// Order ...
type Order struct {
	Price    uint
	Quantity uint
	ID       string
	next     *Order
}

func NewOrder(price uint, quantity uint, id string) (*Order, error) {
	if price <= 0 {
		return nil, errors.New("Price < 0")
	}

	if quantity <= 0 {
		return nil, errors.New("Quantity < 0")
	}

	return &Order{
		Price:    price,
		Quantity: quantity,
		ID:       id,
	}, nil
}

func (order *Order) String() string {
	return fmt.Sprintf("%v@%v/%v", order.Quantity, order.Price, order.ID)
}

func (order *Order) Attach(anotherOrder *Order) {
	if order == nil {
		panic("Attempt to attach to a Nil order")
	}

	if anotherOrder == nil {
		panic("Attempt to attach a Nil order")
	}

	if anotherOrder.next != nil {
		panic("Only new orders can be attached")
	}

	old := order.next
	order.next = anotherOrder
	anotherOrder.next = old
}
