package main

import (
	"service-account/internal/controllers"
	"service-account/internal/infrastructure/database"
	appLogger "service-account/internal/infrastructure/logger" // Alias untuk logger aplikasi
	"service-account/internal/repositories"
	"service-account/internal/usecases"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger" // Alias untuk middleware logging Fiber
)

func main() {
	// Initialize logger
	appLogger.InitLogger()

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		appLogger.Log.Error("Failed to initialize database: ", err)
		return
	}
	appLogger.Log.Info("Database connected successfully")

	// Initialize repositories
	nasabahRepo := repositories.NewNasabahRepository(db)
	saldoRepo := repositories.NewSaldoRepository(db)

	// Initialize use cases
	nasabahUseCase := usecases.NewNasabahUseCase(nasabahRepo, saldoRepo)

	// Initialize controllers
	nasabahController := controllers.NewNasabahController(nasabahUseCase)

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(fiberLogger.New()) // Gunakan middleware logging dari Fiber dengan alias

	// Routes
	app.Post("/daftar", nasabahController.DaftarNasabah)
	app.Post("/tabung", nasabahController.Tabung)
	app.Post("/tarik", nasabahController.Tarik)
	app.Get("/saldo/:no_rekening", nasabahController.CekSaldo)

	// Start server
	port := ":8080"
	appLogger.Log.Info("Server started on port ", port)
	if err := app.Listen(port); err != nil {
		appLogger.Log.Error("Failed to start server: ", err)
	}
}