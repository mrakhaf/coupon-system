package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"coupon-system/internal/config"
	"coupon-system/internal/controller"
	"coupon-system/internal/infrastructure"
	"coupon-system/internal/repository"
	"coupon-system/internal/usecase"
	"coupon-system/internal/validator"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize database
	db, err := infrastructure.NewDatabase(cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repository
	couponRepo := repository.NewCouponRepository(db.DB)

	// Initialize use case
	couponUseCase := usecase.NewCouponUseCase(couponRepo)

	// Initialize controller
	couponController := controller.NewCouponController(couponUseCase)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Validator = validator.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Coupon System API",
			"version": "1.0.0",
		})
	})

	// Coupon routes
	couponGroup := e.Group("/api/coupons")
	couponGroup.POST("", couponController.CreateCoupon)
	couponGroup.POST("/claim", couponController.ClaimCoupon)
	couponGroup.GET("/:name", couponController.GetCouponDetails)

	// Start server
	port := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(e.Start(port))
}
