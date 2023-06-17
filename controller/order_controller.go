package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/simple-transaction-api/common"
	"github.com/hilmiikhsan/simple-transaction-api/configuration"
	"github.com/hilmiikhsan/simple-transaction-api/constants"
	"github.com/hilmiikhsan/simple-transaction-api/model"
	"github.com/hilmiikhsan/simple-transaction-api/service/order"
	"github.com/hilmiikhsan/simple-transaction-api/utils"
)

type OrderController struct {
	order.OrderServiceInterface
	configuration.Config
}

func NewOrderController(orderService *order.OrderServiceInterface, config configuration.Config) *OrderController {
	return &OrderController{
		OrderServiceInterface: *orderService,
		Config:                config,
	}
}

func (controller OrderController) Route(app *fiber.App) {
	app.Post("/orders", controller.CreateOrder)
}

func (controller OrderController) CreateOrder(c *fiber.Ctx) error {
	var req model.CreateOrderModel
	var errMessage []map[string]interface{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	errMessage = common.Validate(req)
	if len(errMessage) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errMessage,
			Data:    nil,
		})
	}

	err = controller.OrderServiceInterface.CreateOrder(c.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrCustomerAndAddressNotFound.Error()) || strings.Contains(err.Error(), constants.ErrProductNotFound.Error()) || strings.Contains(err.Error(), constants.ErrPaymentMethodNotFound.Error()) {
			return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{err.Error()},
				Data:    nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(utils.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.GeneralResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    "",
	})
}
