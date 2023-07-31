package controller

import (
	"errors"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	usecasePatient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePatientController(repo repository.Repository[entityPatient.Patient, string]) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := new(usecasePatient.CreatePatientInput)

		if id, isExist := c.GetPostForm("Id"); isExist && id != "" {
			input.SetId(id)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id sholud exist").Error()})
			return
		}

		if name, isExist := c.GetPostForm("Name"); isExist && name != "" {
			input.SetName(name)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("name sholud exist").Error()})
			return
		}

		output := usecasePatient.CreatePatient(repo, input)

		if output.GetResult() {
			c.JSON(http.StatusOK, gin.H{"Id": output.GetId()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("create patient failed").Error()})
		}
	}
}

func FindAllPatientsController(repo repository.Repository[entityPatient.Patient, string]) gin.HandlerFunc {
	return func(c *gin.Context) {
		output := usecasePatient.FindAllPatient(repo)

		if output.GetResult() {
			c.JSON(http.StatusOK, output.GetData())
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("find all patients failed").Error()})
		}
	}
}
