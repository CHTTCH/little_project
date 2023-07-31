package main

import (
	orderController "github.com/CHTTCH/little_project/backend/adapter/order/controller"
	orderRepo "github.com/CHTTCH/little_project/backend/adapter/order/repository"
	patientController "github.com/CHTTCH/little_project/backend/adapter/patient/controller"
	patientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/repository"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=Zachary password=Zachary dbname=little_project port=5432 sslmode=disable"
)

func main() {
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	patientRepo := patientRepo.NewPatientRepository(postgres)
	orderRepo := orderRepo.NewOrderRepository(postgres)

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // 改為你的前端應用程式的地址
	server.Use(cors.New(config))

	server.GET("/patients", patientController.FindAllPatientsController(patientRepo))
	server.POST("/patients/create", patientController.CreatePatientController(patientRepo))
	server.GET("/orders", orderController.FindAllOrdersController(orderRepo))
	server.POST("/orders/create", orderController.CreateOrderController(patientRepo, orderRepo))
	server.PUT("/orders/edit", orderController.EditOrderController(orderRepo))
	server.Run(":8888")
}
