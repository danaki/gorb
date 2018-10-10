package gorb

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BookTestSuite struct {
	suite.Suite
	*Book
}

func (suite *BookTestSuite) SetupTest() {
	suite.Book = NewBook()
}

func (suite *BookTestSuite) TestProcessBidLimitOrderInsertsOrderIfBothBooksEmpty() {
	executedOrders, err := suite.Book.ProcessBidLimitOrder("#1", 1, 1)

	suite.Nil(err)
	suite.Nil(executedOrders)
}

func (suite *BookTestSuite) TestProcessBidLimitOrderInsertsOrderIfAskPriceNotMatched() {
	askOrder, _ := NewOrder(10, 1, "#1")

	suite.Book.Asks.Insert(askOrder)

	executedOrders, err := suite.Book.ProcessBidLimitOrder("#2", 1, 1)

	suite.Nil(err)
	suite.Nil(executedOrders)
}

func (suite *BookTestSuite) TestProcessBidLimitOrderExecutesOrderIdAskPriceEquallyMatched() {
	askOrder, _ := NewOrder(10, 10, "#1")

	suite.Book.Asks.Insert(askOrder)

	executedOrders, err := suite.Book.ProcessBidLimitOrder("#2", 10, 10)

	suite.Nil(err)
	suite.Equal([]*Order{askOrder}, executedOrders)
	suite.Nil(suite.Book.Asks.Tip)
}

func (suite *BookTestSuite) TestProcessBidLimitOrderExecutesSeveralexecutedOrdersInListUntilPriceNotMatch() {
	askOrder1, _ := NewOrder(10, 10, "#1")
	askOrder2, _ := NewOrder(10, 10, "#2")

	suite.Book.Asks.Insert(askOrder1)
	suite.Book.Asks.Insert(askOrder2)

	executedOrders, err := suite.Book.ProcessBidLimitOrder("#3", 10, 20)

	suite.Nil(err)
	suite.Equal([]*Order{askOrder1, askOrder2}, executedOrders)
	suite.Nil(suite.Book.Asks.Tip)
}

func (suite *BookTestSuite) TestProcessBidLimitOrderExecutesAndStopsIfQuantityReached() {
	askOrder1, _ := NewOrder(10, 10, "#1")
	askOrder2, _ := NewOrder(10, 10, "#2")

	suite.Book.Asks.Insert(askOrder1)
	suite.Book.Asks.Insert(askOrder2)

	executedOrders, err := suite.Book.ProcessBidLimitOrder("#3", 10, 10)

	suite.Nil(err)
	suite.Equal([]*Order{askOrder1}, executedOrders)
	suite.Equal(suite.Book.Asks.AsSlice(), []*Order{askOrder2})
}

func (suite *BookTestSuite) TestProcessBidLimitOrderExecutesPartially() {
	askOrder, _ := NewOrder(10, 10, "#1")

	suite.Book.Asks.Insert(askOrder)

	executedOrders, err := suite.Book.ProcessBidLimitOrder("#3", 10, 5)

	suite.Nil(err)
	suite.Equal([]*Order{askOrder}, executedOrders)
	suite.Equal(suite.Book.Asks.AsSlice(), []*Order{askOrder})
	suite.Equal(uint(5), askOrder.Quantity)
}

func (suite *BookTestSuite) TestProcessBidLimitOrderExecutesUntilPriceIsMatchedAndLeavesUnexecutedPartialOrder() {
	askOrder1, _ := NewOrder(10, 10, "#1")
	askOrder2, _ := NewOrder(20, 10, "#1")

	suite.Book.Asks.Insert(askOrder1)
	suite.Book.Asks.Insert(askOrder2)

	executedOrders, err := suite.Book.ProcessBidLimitOrder("#3", 15, 15)

	bidOrder := suite.Book.Bids.AsSlice()[0]

	suite.Nil(err)
	suite.Equal([]*Order{askOrder1}, executedOrders)
	suite.Equal(suite.Book.Asks.AsSlice(), []*Order{askOrder2})
	suite.Equal(uint(5), bidOrder.Quantity)
	suite.Equal(uint(15), bidOrder.Price)
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
