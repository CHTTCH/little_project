package main

import (
	createOrderController "github.com/CHTTCH/little_project/backend/adapter/order/controller"
	orderRepo "github.com/CHTTCH/little_project/backend/adapter/order/repository"
	createPatientController "github.com/CHTTCH/little_project/backend/adapter/patient/controller"
	patientRepo "github.com/CHTTCH/little_project/backend/adapter/patient/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type IndexData struct {
	Title   string
	Content string
}

const (
	dsn = "host=localhost user=Zachary password=Zachary dbname=little_project port=5432 sslmode=disable"
)

func homePage(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	patientRepo := &patientRepo.PatientRepository{DB: postgres}
	orderRepo := &orderRepo.OrderRepository{DB: postgres}

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	server := gin.Default()
	server.LoadHTMLGlob("../frontend/*")
	server.GET("/", homePage)
	server.GET("/patients", createPatientController.FindAllPatientsController(patientRepo))
	server.POST("/patients/create", createPatientController.CreatePatientController(patientRepo))
	server.POST("/orders/create", createOrderController.CreateOrderController(patientRepo, orderRepo))
	server.Run(":8888")
}
