package controllers

import (
	"net/http"
	"tripit/models"
	"tripit/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = utils.OpenCollection(utils.Client, "users")
var validate = validator.New()

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseMessage{
				Success: false,
				Message: validationErr.Error(),
			})
			return
		}

		count, err := userCollection.CountDocuments(c, bson.M{"email": user.Email})

		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, utils.ResponseMessage{
				Success: false,
				Message: "an error occured! email is already used!",
			})
			return
		}

		user.Encrypt(user.Password)
		user.ID = primitive.NewObjectID()
		token, _ := user.GetSignedJWT(user.ID.Hex())

		resultInsertionNumber, insertErr := userCollection.InsertOne(c, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, utils.ResponseMessage{
				Success: false,
				Message: insertErr.Error(),
			})
			return
		}

		c.SetCookie("token", token, 2000, "/", "", false, true)
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, utils.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
		}

		err := userCollection.FindOne(c, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResponseMessage{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		passwordIsValid, msg := user.VerifyPassword(foundUser.Password)

		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, utils.ResponseMessage{
				Success: false,
				Message: msg,
			})
			return
		}

		token, _ := foundUser.GetSignedJWT(foundUser.ID.Hex())

		c.SetCookie("token", token, 2000, "/", "", false, true)
		c.JSON(http.StatusOK, utils.ResponseUser{
			Success: true,
			Message: foundUser,
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, objErr := primitive.ObjectIDFromHex(c.Keys["id"].(string))
		if objErr != nil {
			c.JSON(400, utils.ResponseMessage{
				Success: false,
				Message: "Bad cookie",
			})
		}

		result := models.User{}

		findOneErr := userCollection.FindOne(c.Request.Context(), bson.M{
			"_id": id,
		}).Decode(&result)
		if findOneErr != nil {
			c.JSON(400, utils.ResponseMessage{
				Success: false,
				Message: "Query error for id",
			})
			return
		}

		c.JSON(200, utils.ResponseUser{
			Success: true,
			Message: result,
		})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", 2000, "/", "", false, true)

		c.JSON(200, utils.ResponseMessage{
			Success: true,
			Message: "You have been logged out",
		})
	}
}
