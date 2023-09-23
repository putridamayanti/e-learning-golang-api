package controllers

import (
	"context"
	"elearning/config"
	"elearning/models"
	"elearning/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

const UserCollection = "users"

func GetUserByQuery(c *gin.Context)  {
	var query models.Query
	err := c.ShouldBindQuery(&query); if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	filter := query.GetQueryFind()

	opts := query.GetOptions()
	opts.SetProjection(bson.M{"password": 0})

	results := make([]models.User, 0)
	cursor := config.Find("users", filter, opts)
	for cursor.Next(context.Background()) {
		var data models.User

		if cursor.Decode(&data) == nil {
			results = append(results, data)
		}
	}

	c.JSON(http.StatusOK, models.Result{Data: results}); return
}

func CreateUser(c *gin.Context)  {
	var request models.User
	err := c.ShouldBindQuery(&request); if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	request.Id			= uuid.NewString()
	request.CreatedAt	= time.Now()

	_, err = config.InsertOne(UserCollection, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error()); return
	}

	c.JSON(http.StatusOK, "Success")
}

func GetUserById(c *gin.Context)  {
	id := c.Param("id")

	cursor := config.FindOne(UserCollection, bson.M{ "id": id }, options.FindOne().SetProjection(bson.M{"password": 0}))
	var result models.Course
	err := cursor.Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}


	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context)  {
	id := c.Param("id")

	cursor := config.FindOne(UserCollection, bson.M{ "id": id }, nil)
	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Result{Data: "Data Not Found"}); return
	}

	var request models.User
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: err.Error()}); return
	}

	if request.Password != "" {
		request.Password = utils.HashAndSalt(request.Password)
	} else {
		request.Password = user.Password
	}

	_, err = config.UpdateOne(UserCollection, bson.M{"id": id}, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: err.Error()}); return
	}

	c.String(200, "Success")
}

func DeleteUser(c *gin.Context)  {
	id := c.Param("id")

	_, err := config.DeleteOne(UserCollection, bson.M{"id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Result{Data: "Failed Delete Data"}); return
	}

	c.String(http.StatusOK, "Success")
}