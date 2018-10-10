package gorb

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrderListTestSuite struct {
	suite.Suite
	DescOrderList *OrderList
	AscOrderList  *OrderList
}

func (suite *OrderListTestSuite) SetupTest() {
	suite.DescOrderList = NewOrderList(true)
	suite.AscOrderList = NewOrderList(false)
}

func (suite *OrderListTestSuite) TestInsert() {
	order, _ := NewOrder(1, 1, "#1")

	suite.DescOrderList.Insert(order)
	suite.AscOrderList.Insert(order)

	suite.Equal(suite.DescOrderList.AsSlice(), []*Order{order})
	suite.Equal(suite.AscOrderList.AsSlice(), []*Order{order})
}

func (suite *OrderListTestSuite) TestInsertInDirectOrder() {
	order1, _ := NewOrder(1, 1, "#1")
	order2, _ := NewOrder(2, 1, "#2")

	suite.DescOrderList.Insert(order1)
	suite.DescOrderList.Insert(order2)

	suite.AscOrderList.Insert(order1)
	suite.AscOrderList.Insert(order2)

	suite.Equal(suite.DescOrderList.AsSlice(), []*Order{order1, order2})
	suite.Equal(suite.AscOrderList.AsSlice(), []*Order{order2, order1})
}

func (suite *OrderListTestSuite) TestInsertReverseOrder() {
	order1, _ := NewOrder(1, 1, "#1")
	order2, _ := NewOrder(2, 1, "#2")

	suite.DescOrderList.Insert(order1)
	suite.DescOrderList.Insert(order2)

	suite.AscOrderList.Insert(order1)
	suite.AscOrderList.Insert(order2)

	suite.Equal(suite.DescOrderList.AsSlice(), []*Order{order1, order2})
	suite.Equal(suite.AscOrderList.AsSlice(), []*Order{order2, order1})
}

func (suite *OrderListTestSuite) TestInsertEqualPriceOrdededByInsertionOrder() {
	order1, _ := NewOrder(1, 1, "#1")
	order1duplicate, _ := NewOrder(1, 1, "#2")

	suite.DescOrderList.Insert(order1)
	suite.DescOrderList.Insert(order1duplicate)

	suite.AscOrderList.Insert(order1)
	suite.AscOrderList.Insert(order1duplicate)

	suite.Equal(suite.DescOrderList.AsSlice(), []*Order{order1, order1duplicate})
	suite.Equal(suite.AscOrderList.AsSlice(), []*Order{order1, order1duplicate})
}

func (suite *OrderListTestSuite) TestInsertInBetween() {
	order1, _ := NewOrder(1, 1, "#1")
	order2, _ := NewOrder(2, 1, "#2")
	order3, _ := NewOrder(3, 1, "#3")

	suite.DescOrderList.Insert(order1)
	suite.DescOrderList.Insert(order3)
	suite.DescOrderList.Insert(order2)

	suite.AscOrderList.Insert(order1)
	suite.AscOrderList.Insert(order3)
	suite.AscOrderList.Insert(order2)

	suite.Equal(suite.DescOrderList.AsSlice(), []*Order{order1, order2, order3})
	suite.Equal(suite.AscOrderList.AsSlice(), []*Order{order1, order2, order3})
}

func (suite *OrderListTestSuite) TestProcessMarketOrderReturnsErrorOnZeroQuantity() {
	_, err := suite.AscOrderList.ProcessMarketOrder(0)

	suite.NotNil(err)
}

func (suite *OrderListTestSuite) TestProcessMarketOrderReturnsEmptyOnEmptyOrderList() {
	orders, err := suite.AscOrderList.ProcessMarketOrder(1)

	suite.Nil(err)
	suite.Equal(orders, []*Order{})
}

func (suite *OrderListTestSuite) TestProcessMarketOrderExactQuantityMatch() {
	order, _ := NewOrder(1, 10, "#1")

	orders, err := suite.AscOrderList.ProcessMarketOrder(10)

	suite.Nil(err)
	suite.Equal(orders, []*Order{order})
	suite.Equal(10, order.Quantity)
	suite.Nil(suite.AscOrderList.Tip)
}

func (suite *OrderListTestSuite) TestProcessMarketOrderLessThenAvailableQuantity() {
	order, _ := NewOrder(1, 10, "#1")

	orders, err := suite.AscOrderList.ProcessMarketOrder(1)

	suite.Nil(err)
	suite.Equal(orders, []*Order{order})
	suite.Equal(9, order.Quantity)
	suite.Equal(order, suite.AscOrderList.Tip)
}

func (suite *OrderListTestSuite) TestProcessMarketOrderExceedsAvailableQuantity() {
	order, _ := NewOrder(1, 1, "#1")

	orders, err := suite.AscOrderList.ProcessMarketOrder(10)

	suite.Nil(err)
	suite.Equal(orders, []*Order{order})
	suite.Equal(10, order.Quantity)
	suite.Equal(order, suite.AscOrderList.Tip)
}

func (suite *OrderListTestSuite) TestProcessMarketOrderFollowsQuantityChain() {
	order1, _ := NewOrder(1, 1, "#1")
	order2, _ := NewOrder(1, 2, "#1")

	orders, err := suite.AscOrderList.ProcessMarketOrder(2)

	suite.Nil(err)
	suite.Equal(orders, []*Order{order1, order2})
	suite.Equal(1, order1.Quantity)
	suite.Equal(1, order2.Quantity)
	suite.Equal(order2, suite.AscOrderList.Tip)
}

func TestOrderListTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
