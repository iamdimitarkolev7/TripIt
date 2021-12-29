package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" validate:"required,min=3,max=30"`
	Email    string             `json:"email" validate:"required,min=3,max=30"`
	Password string             `json:"password" validate:"required,min=6,max=20"`
	UserType string             `json:"user_type"`
}

func (user *User) Encrypt(password string) {
	hash, bcryptErr := bcrypt.GenerateFromPassword([]byte(password), 14)

	if bcryptErr != nil {
		log.Fatal(bcryptErr)
	}

	user.Password = string(hash)
}

func (user *User) VerifyPassword(providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(user.Password))
	check := true
	msg := ""

	if err != nil {
		msg = "email or passowrd is incorrect"
		check = false
	}

	return check, msg
}

func (user *User) GetSignedJWT(id string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
