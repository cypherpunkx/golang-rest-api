package app

import (
	"fmt"
	"golang-rest-api/internal/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	config config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{config: cfg}
}

func (a *App) Run() error {
	// Inisialisasi router dan handlers
	router := initializeRouter()

	// Jalankan server
	return router.Run(fmt.Sprintf(":%d", a.config.Server.Port))
}

func initializeRouter() *gin.Engine {
	// Inisialisasi router Gin
	router := gin.Default()

	// Register handlers
	// userHandler := handlers.NewUserHandler()
	// taskHandler := handlers.NewTaskHandler()

	// Register routes
	// v1 := router.Group("/api/v1")
	// {
	// 	userRoutes := v1.Group("/users")
	// 	{
	// 		userRoutes.GET("/:id", userHandler.GetUserByID)
	// 		// Tambahkan rute lain untuk pengguna di sini
	// 	}

	// 	taskRoutes := v1.Group("/tasks")
	// 	{
	// 		taskRoutes.GET("/:id", taskHandler.GetTaskByID)
	// 		// Tambahkan rute lain untuk tugas di sini
	// 	}
	// }

	return router
}
