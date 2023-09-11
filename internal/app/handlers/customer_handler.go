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

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		fmt.Println(err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		// Memeriksa jika ada kesalahan validasi
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Field: %s, Error: %s", fieldError.Field(), fieldError.Tag()),
			})
			return
		}
	}

	data, err := h.customerService.Create(&requestBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Println(err.Error())
		return
	}

	response := domain.Response{
		Message: "Customer Was Created!",
		Data:    data,
	}

	c.JSON(http.StatusCreated, response)

}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers := []domain.Customer{}

	data, err := h.customerService.View(&customers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Println(err.Error())
		return
	}

	response := domain.Response{
		Message: "Get All Data Customers",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)

}
