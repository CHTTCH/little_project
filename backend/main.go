package main

import (
	mockRepo "github.com/CHTTCH/little_project/backend/adapter/patient/mock_repository"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	createPatientUseCase "github.com/CHTTCH/little_project/backend/adapter/patient/controller"
	"github.com/gin-gonic/gin"
    "net/http"
)

type IndexData struct {
	Title   string
	Content string
}

func homePage(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	repo := &mockRepo.MockPatientRepository{ PatientList: []entityPatient.Patient{} }
	server := gin.Default()
	server.LoadHTMLGlob("../frontend/*")
	server.GET("/", homePage)
	server.POST("/patient/create", createPatientUseCase.CreatePatientController(repo))
	server.Run(":8888")
}