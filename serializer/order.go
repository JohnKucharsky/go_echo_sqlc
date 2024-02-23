package serializer

import (
	"github.com/JohnKucharsky/go_echo_sqlc/internal/database"
	"github.com/JohnKucharsky/go_echo_sqlc/utils"
	"time"
)

type OrderBody struct {
	ProductId int32 `json:"product_id" validate:"required"`
	UserId    int32 `json:"user_id" validate:"required"`
}

type Order struct {
	OrderID       int32
	ProductName   string
	ProductSerial *string
	UserName      string
	UserLastName  *string
	UpdatedAt     time.Time
}

func DatabaseOrderToOrder(dbOrder database.GetOrdersRow) Order {
	return Order{
		OrderID:       dbOrder.OrderID,
		ProductName:   dbOrder.ProductName,
		ProductSerial: utils.SqlNullStringToString(dbOrder.ProductSerial),
		UserName:      dbOrder.UserName,
		UserLastName:  utils.SqlNullStringToString(dbOrder.UserLastName),
		UpdatedAt:     dbOrder.UpdatedAt,
	}
}

func OrderBodyToOrderCreate(body OrderBody) database.CreateOrderParams {
	return database.CreateOrderParams{
		ProductID: body.ProductId,
		UserID:    body.UserId,
		UpdatedAt: time.Now().Local(),
	}
}

func OrderBodyToOrderUpdate(
	body OrderBody,
	id int32,
) database.UpdateOrderParams {
	return database.UpdateOrderParams{
		ProductID: body.ProductId,
		UserID:    body.UserId,
		UpdatedAt: time.Now().Local(),
		ID:        id,
	}
}

func DatabaseOrdersToOrders(dbOrders []database.GetOrdersRow) []Order {
	var orders []Order

	for _, order := range dbOrders {
		orders = append(orders, DatabaseOrderToOrder(order))
	}

	return orders
}
