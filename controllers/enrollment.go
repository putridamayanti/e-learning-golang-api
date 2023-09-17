package controllers

import (
	"elearning/config"
	"elearning/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const EnrollmentCollection = "enrollments"

func EnrollCourse(c *gin.Context)  {
	var request models.Enrollment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}

	request.Id			= uuid.NewString()
	request.CreatedAt	= time.Now()

	_, err = config.InsertOne(EnrollmentCollection, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	c.JSON(http.StatusOK, models.Result{Data: "Success"})
}
