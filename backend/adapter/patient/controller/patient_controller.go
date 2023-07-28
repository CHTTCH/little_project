package controller

import (
	"errors"
	entityPatient "github.com/CHTTCH/little_project/backend/entity/patient"
	usecasePateient "github.com/CHTTCH/little_project/backend/usecase/patient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePatientController(repo usecasePateient.Repository[entityPatient.Patient]) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := &usecasePateient.CreatePatientInput{}

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
			c.JSON(http.StatusOK, gin.H{"error": errors.New("save failed")})
		}
	}
}
