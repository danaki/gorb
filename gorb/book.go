package gorb

import (
	"fmt"
)

// Book ...
type Book struct {
	Bids      *OrderList
	Asks      *OrderList
	LastPrice int
}

func NewBook() *Book {
	return &Book{
		Bids: NewOrderList(true),
		Asks: NewOrderList(false),
	}
}

func (book *Book) String() string {
	ret := ""

	ret += "Bids: \n"
	ret += fmt.Sprintf("%v", book.Bids) + "\n"

	ret += "Asks: \n"
	ret += fmt.Sprintf("%v", book.Asks) + "\n"

	return ret
}

func (book *Book) ProcessBidMarketOrder(quantity uint) ([]*Order, error) {
	return book.Bids.ProcessMarketOrder(quantity)
}

func (book *Book) ProcessAskMarketOrder(quantity uint) ([]*Order, error) {
	return book.Asks.ProcessMarketOrder(quantity)
}

func (book *Book) processLimitOrder(
	list *OrderList,
	oppositeList *OrderList,
	id string,
	price uint,
	quantity uint) (ret []*Order, err error) {

	var order *Order

	if (list.Tip == nil) ||
		((list.PriceDescending && (price > list.Tip.Price)) ||
			(!list.PriceDescending && (price < list.Tip.Price))) {

		order, err = NewOrder(price, quantity, id)
		if err == nil {
			book.Bids.Insert(order)
		}

		return
	}

	ret = make([]*Order, 0)
	err = nil

	cur := list.Tip

Loop:
	for {
		switch {
		case cur == nil:
			break Loop

		case (list.PriceDescending && (cur.Price < price)) ||
			(!list.PriceDescending && (cur.Price > price)):
			order, err = NewOrder(price, quantity, id)
			if err != nil {
				return
			}

			oppositeList.Insert(order)
			break Loop

		case cur.Quantity <= quantity:
			list.Tip = cur.next
			ret = append(ret, cur)
			quantity -= cur.Quantity
			cur = cur.next

			if quantity == 0 {
				break Loop
			}

		case cur.Quantity > quantity:
			cur.Quantity -= quantity
			ret = append(ret, cur)
			break Loop

		default:
			panic("Unexpected state")
		}
	}

	return
}

func (book *Book) ProcessBidLimitOrder(id string, price uint, quantity uint) (ret []*Order, err error) {
	return book.processLimitOrder(book.Asks, book.Bids, id, price, quantity)
}

func (book *Book) ProcessAskLimitOrder(id string, price uint, quantity uint) (ret []*Order, err error) {
	return book.processLimitOrder(book.Bids, book.Asks, id, price, quantity)
}

func (book *Book) ByIndex(bidOrAsk bool, i uint) (*Order, bool) {
	if bidOrAsk {
		return book.Bids.ByIndex(i)
	} else {
		return book.Asks.ByIndex(i)
	}
}
