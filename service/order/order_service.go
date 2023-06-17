package order

import (
	"context"
	"time"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"github.com/hilmiikhsan/simple-transaction-api/model"
	"github.com/hilmiikhsan/simple-transaction-api/repository/customer_address"
	"github.com/hilmiikhsan/simple-transaction-api/repository/order"
	"github.com/hilmiikhsan/simple-transaction-api/repository/order_payment"
	"github.com/hilmiikhsan/simple-transaction-api/repository/order_product"
	"github.com/hilmiikhsan/simple-transaction-api/repository/payment_method"
	"github.com/hilmiikhsan/simple-transaction-api/repository/product"
	"gorm.io/gorm"
)

func NewOrderServiceInterface(orderRepository *order.OrderRepositoryInterface, db *gorm.DB, customerAddressRepository *customer_address.CustomerAddressRepositoryInterface, orderProductRepository *order_product.OrderProductRepositoryInterface, productRepository *product.ProductRepositoryInterface, paymentMethodRepository *payment_method.PaymentMenthodRepositoryInterface, orderPaymentRepository *order_payment.OrderPaymentRepositoryInterface) OrderServiceInterface {
	return &orderService{
		OrderRepositoryInterface:           *orderRepository,
		DB:                                 db,
		CustomerAddressRepositoryInterface: *customerAddressRepository,
		OrderProductRepositoryInterface:    *orderProductRepository,
		ProductRepositoryInterface:         *productRepository,
		PaymentMenthodRepositoryInterface:  *paymentMethodRepository,
		OrderPaymentRepositoryInterface:    *orderPaymentRepository,
	}
}

type orderService struct {
	order.OrderRepositoryInterface
	*gorm.DB
	customer_address.CustomerAddressRepositoryInterface
	order_product.OrderProductRepositoryInterface
	product.ProductRepositoryInterface
	payment_method.PaymentMenthodRepositoryInterface
	order_payment.OrderPaymentRepositoryInterface
}

func (orderService *orderService) CreateOrder(ctx context.Context, order model.CreateOrderModel) error {
	tx := orderService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	customerAddressData, err := orderService.CustomerAddressRepositoryInterface.FindByID(ctx, order.CustomerAddressID, order.CustomerID)
	if err != nil {
		tx.Rollback()
		return err
	}

	orderID, err := orderService.OrderRepositoryInterface.Insert(ctx, tx, entity.Order{
		CustomerID:        order.CustomerID,
		CustomerAddressID: customerAddressData.ID,
		OrderDate:         time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	var productIDs, paymentMethodIDs []int
	var orderProducts []entity.OrderProduct
	var orderPayments []entity.OrderPayment

	for _, detail := range order.OrderProducts {
		productIDs = append(productIDs, detail.ProductID)
	}

	for _, detailPayment := range order.OrderPayments {
		paymentMethodIDs = append(paymentMethodIDs, detailPayment.PaymentMethodID)
	}

	productData, err := orderService.ProductRepositoryInterface.FindBatchByID(ctx, productIDs)
	if err != nil {
		tx.Rollback()
		return err
	}

	paymentMethodData, err := orderService.PaymentMenthodRepositoryInterface.FindBatchByID(ctx, paymentMethodIDs)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, detailOrderProduct := range order.OrderProducts {
		var product entity.Product
		for _, x := range productData {
			if x.ID == detailOrderProduct.ProductID {
				product = x
				break
			}
		}

		orderProducts = append(orderProducts, entity.OrderProduct{
			OrderID:   orderID,
			ProductID: product.ID,
		})
	}

	for _, detailOrderPayment := range order.OrderPayments {
		var paymentMethod entity.PaymentMethod
		for _, x := range paymentMethodData {
			if x.ID == detailOrderPayment.PaymentMethodID {
				paymentMethod = x
				break
			}
		}

		orderPayments = append(orderPayments, entity.OrderPayment{
			OrderID:         orderID,
			PaymentMethodID: paymentMethod.ID,
		})
	}

	err = orderService.OrderProductRepositoryInterface.InsertBatch(ctx, tx, orderProducts)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = orderService.OrderPaymentRepositoryInterface.InsertBatch(ctx, tx, orderPayments)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
