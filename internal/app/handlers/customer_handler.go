package handlers

import (
	"fmt"
	"golang-rest-api/internal/app/services"
	"golang-rest-api/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	requestBody := domain.Customer{}

	if err := c.ShouldBind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		fmt.Println(err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		fmt.Println(err.Error())
		return
	}

	customer, err := h.customerService.Create(requestBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": customer,
	})

}
