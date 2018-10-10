package gorb

import (
	"errors"
	"fmt"
)

type OrderList struct {
	PriceDescending bool
	Tip             *Order
}

func NewOrderList(priceDescending bool) *OrderList {
	return &OrderList{PriceDescending: priceDescending}
}

func (list *OrderList) String() string {
	ret := ""

	for ord := list.Tip; ord != nil; ord = ord.next {
		ret += fmt.Sprintf("%v", ord) + "\n"
	}

	return ret
}

func (list *OrderList) Insert(order *Order) {
	if list.Tip == nil {
		list.SetTip(order)
		return
	}

	var cur, prev *Order

	cur = list.Tip
	for cur != nil &&
		((list.PriceDescending && (cur.Price >= order.Price)) ||
			(!list.PriceDescending && (cur.Price <= order.Price))) {

		prev = cur
		cur = cur.next
	}

	if prev == nil {
		list.SetTip(order)
	} else {
		prev.Attach(order)
	}
}

func (list *OrderList) ProcessMarketOrder(quantity uint) (ret []*Order, err error) {
	ret = make([]*Order, 0)
	err = nil

	if quantity <= 0 {
		err = errors.New("Quantity must be > 0")
		return
	}

	cur := list.Tip

Loop:
	for {
		switch {
		case cur == nil:
			break Loop
		case cur.Quantity <= quantity:
			list.Tip = cur.next
			ret = append(ret, cur)
			quantity -= cur.Quantity

			if cur.Quantity == quantity {
				break Loop
			}
		case cur.Quantity > quantity:
			cur.Quantity -= quantity
			break Loop
		}
	}

	return
}

func (list *OrderList) SetTip(order *Order) {
	if order.next != nil {
		panic("Only new orders can be attached")
	}

	order.next = list.Tip
	list.Tip = order
}

func (list *OrderList) ByIndex(i uint) (*Order, bool) {
	cur := list.Tip
	var j uint = 0

	for ; j < i && cur != nil; j++ {
		cur = cur.next
	}

	return cur, cur != nil
}

func (list *OrderList) AsSlice() []*Order {
	var result []*Order

	for cur := list.Tip; cur != nil; cur = cur.next {
		result = append(result, cur)
	}

	return result
}
