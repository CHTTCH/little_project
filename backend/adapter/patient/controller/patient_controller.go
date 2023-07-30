package controller

import (
	"errors"
	"net/http"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	usecasePateient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/CHTTCH/little_project/backend/usecase/repository"
	"github.com/gin-gonic/gin"
)

func CreatePatientController(repo repository.Repository[entityPatient.Patient]) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := usecasePateient.CreatePatientInput{}

		if id, isExist := c.GetPostForm("Id"); isExist && id != "" {
			input.SetId(id)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id sholud exist")})
			return
		}

		if name, isExist := c.GetPostForm("Name"); isExist && name != "" {
			input.SetName(name)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("name sholud exist").Error()})
			return
		}
	
		output := usecasePateient.CreatePatient(repo, input)

		if output.GetResult() {
			c.JSON(http.StatusOK, gin.H{"Id": output.GetId()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("save failed")})
		}
	}
}

func FindAllPatientController(repo repository.Repository[entityPatient.Patient]) gin.HandlerFunc {
	return func(c *gin.Context) {
		output:= usecasePateient.FindAllPatient(repo)
		
		if output.GetResult() {
			c.JSON(http.StatusOK, output.GetData())
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("extract failed")})
		}
	}
}