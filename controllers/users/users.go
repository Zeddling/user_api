package users

import (
	"errors"
	"net/http"

	"github.com/Zeddling/user/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var (
	errInvalidBody = errors.New("invalid request body")
	errInvalidID   = errors.New("invalid ID")
	errNotExist    = errors.New("users don't exist")
)

func badRequest(c *gin.Context, err error, msg error) {
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": msg,
			},
		)
		return
	}
}

func DeleteById(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	err := models.Delete(id)
	if err != nil {
		badRequest(c, err, errNotExist)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Success",
		},
	)
}

func FindAllUsers(c *gin.Context) {
	users, err := models.FindAll()
	if err != nil {
		badRequest(c, err, errNotExist)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"users":  &users,
		},
	)
}

func FindUserByID(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))

	user, err := models.FindById(id)
	if err != nil {
		badRequest(c, err, errInvalidID)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"users":  &user,
		},
	)
}

func SaveUser(c *gin.Context) {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		badRequest(c, err, errInvalidBody)
		return
	}

	user, err = models.Save(user)
	if err != nil {
		badRequest(c, err, errNotExist)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusCreated,
			"users":  &user,
		},
	)
}

func UpdateUser(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	existing, err := models.FindById(id)
	if err != nil {
		badRequest(c, err, errNotExist)
		return
	}

	err = c.Bind(&existing)
	if err != nil {
		badRequest(c, err, errInvalidBody)
		return
	}

	user := models.Update(existing)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"user":   &user,
		},
	)
}
