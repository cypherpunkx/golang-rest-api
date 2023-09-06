package app

import (
	"fmt"
	"golang-rest-api/internal/app/handlers"
	"golang-rest-api/internal/app/repositories"
	"golang-rest-api/internal/app/services"
	"golang-rest-api/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	config config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{config: cfg}
}

func (a *App) Run() error {
	// Inisialisasi router dan handlers
	db := config.NewDB(a.config)

	router := initializeRouter(db.ConnectDB())

	// Jalankan server
	return router.Run(fmt.Sprintf(":%d", a.config.Server.Port))
}

func initializeRouter(db *gorm.DB) *gin.Engine {
	// Inisialisasi router Gin
	router := gin.Default()

	// Register handlers
	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)

	// Register routes
	v1 := router.Group("/api/v1")
	{
		customerRoutes := v1.Group("/customers")
		{
			customerRoutes.POST("/", customerHandler.CreateCustomer)
			// Tambahkan rute lain untuk pengguna di sini
		}

	}

	return router
}
