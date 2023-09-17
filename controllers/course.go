package controllers

import (
	"elearning/config"
	"elearning/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

const CourseCollection = "courses"

func GetCourseByQuery(c *gin.Context)  {
	var query models.Query
	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	filter := query.GetQueryFind()

	opts := query.GetOptions()

	results := config.Find(CourseCollection, filter, opts)

	c.JSON(http.StatusOK, models.Result{Data: results}); return
}

func CreateCourse(c *gin.Context)  {
	var request models.Course
	err := c.ShouldBindQuery(&request); if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	request.Id			= uuid.NewString()
	request.CreatedAt	= time.Now()

	_, err = config.InsertOne(CourseCollection, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	c.JSON(http.StatusOK, models.Result{Data: "Success"})
}

func GetCourseById(c *gin.Context)  {
	id := c.Param("id")

	cursor := config.FindOne(CourseCollection, bson.M{ "id": id }, nil)
	var course models.Course
	err := cursor.Decode(&course)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}

	c.JSON(http.StatusOK, models.Result{Data: course})
}

func UpdateCourse(c *gin.Context)  {
	id := c.Param("id")

	cursor := config.FindOne(CourseCollection, bson.M{ "id": id }, nil)
	var course models.Course
	err := cursor.Decode(&course)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}

	var request models.Course
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: err.Error()}); return
	}

	request.Id			= course.Id
	request.CreatedAt	= course.CreatedAt

	_, err = config.UpdateOne(CourseCollection, bson.M{"id": id}, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: err.Error()}); return
	}

	c.JSON(200, models.Result{Data: "Success"})
}

func DeleteCourse(c *gin.Context)  {
	id := c.Param("id")

	_, err := config.DeleteOne(CourseCollection, bson.M{"id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: "Failed Delete Data"}); return
	}

	c.JSON(http.StatusOK, models.Result{Data: "Success"})
}

func AddCourseContent(c *gin.Context)  {
	id := c.Param("id")

	type RequestContent struct {
		Content 		string		`json:"content"`
	}
	var request RequestContent
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}

	cursor := config.FindOne(CourseCollection, bson.M{ "id": id }, nil)
	var course models.Course
	err = cursor.Decode(&course)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}

	course.Content = append(course.Content, request.Content)

	_, err = config.UpdateOne(CourseCollection, bson.M{"id": id}, course)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: err.Error()}); return
	}

	c.JSON(http.StatusOK, models.Result{Data: course})
}