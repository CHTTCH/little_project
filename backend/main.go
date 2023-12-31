package main

import (
	orderController "github.com/CHTTCH/little_project/backend/adapter/order/controller"
	orderRepo "github.com/CHTTCH/little_project/backend/adapter/order/repository"
	patientController "github.com/CHTTCH/little_project/backend/adapter/patient/controller"
	patientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/repository"
	"github.com/CHTTCH/little_project/backend/scripts"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	if err := scripts.CreateFivePatients(postgres, patientRepo); err != nil {
		panic(err.Error())
	}

	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	server.Use(cors.New(config))

	server.GET("/patients", patientController.FindAllPatientsController(patientRepo))
	server.POST("/patients/create", patientController.CreatePatientController(patientRepo))
	server.GET("/orders", orderController.FindAllOrdersController(orderRepo))
	server.POST("/orders/create", orderController.CreateOrderController(patientRepo, orderRepo))
	server.PUT("/orders/edit", orderController.EditOrderController(orderRepo))
	server.Run(":8888")
}
