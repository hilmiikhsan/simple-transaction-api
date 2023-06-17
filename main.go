package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hilmiikhsan/simple-transaction-api/configuration"
	"github.com/hilmiikhsan/simple-transaction-api/controller"
	customerAddressRepo "github.com/hilmiikhsan/simple-transaction-api/repository/customer_address"
	orderRepo "github.com/hilmiikhsan/simple-transaction-api/repository/order"
	orderPaymentRepo "github.com/hilmiikhsan/simple-transaction-api/repository/order_payment"
	orderProductRepo "github.com/hilmiikhsan/simple-transaction-api/repository/order_product"
	paymentMethodRepo "github.com/hilmiikhsan/simple-transaction-api/repository/payment_method"
	productRepo "github.com/hilmiikhsan/simple-transaction-api/repository/product"
	orderService "github.com/hilmiikhsan/simple-transaction-api/service/order"
)

func main() {
	config := configuration.New()
	db := configuration.ConnectDatabase(config)

	// repository
	orderRepository := orderRepo.NewOrderRepositoryInterface(db)
	customerAddressRepository := customerAddressRepo.NewCustomerAddressInterface(db)
	orderProductRepository := orderProductRepo.NewOrderProductRepositoryInterface(db)
	productRepository := productRepo.NewProductRepositoryInterface(db)
	paymentMnethodRepository := paymentMethodRepo.NewPaymentMethodRepositoryInterface(db)
	orderPaymentRepository := orderPaymentRepo.NewOrderPaymentRepositoryInterface(db)

	// service
	orderService := orderService.NewOrderServiceInterface(&orderRepository, db, &customerAddressRepository, &orderProductRepository, &productRepository, &paymentMnethodRepository, &orderPaymentRepository)

	// controller
	orderController := controller.NewOrderController(&orderService, config)

	// setup fiber
	app := fiber.New()
	app.Use(recover.New())

	// routing
	orderController.Route(app)

	err := app.Listen(config.Get("SERVER.PORT"))
	if err != nil {
		log.Fatal("Error running server")
	}
}
