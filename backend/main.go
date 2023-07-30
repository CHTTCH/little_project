package main

import (
	repo "github.com/CHTTCH/little_project/backend/adapter/patient/repository"
	createPatientUseCase "github.com/CHTTCH/little_project/backend/adapter/patient/controller"
	"github.com/gin-gonic/gin"
    "net/http"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type IndexData struct {
	Title   string
	Content string
}

const( dsn = "host=localhost user=Zachary password=Zachary dbname=little_project port=5432 sslmode=disable")

func homePage(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	repo := &repo.Repository{ DB: postgres }

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	
	server := gin.Default()
	server.LoadHTMLGlob("../frontend/*")
	server.GET("/", homePage)
	server.GET("/patients", createPatientUseCase.FindAllPatientController(repo))
	server.POST("/patients/create", createPatientUseCase.CreatePatientController(repo))
	server.Run(":8888")
}