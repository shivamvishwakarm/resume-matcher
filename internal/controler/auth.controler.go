package controler

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shivamvishwakarm/resume-matcher/internal/config"
	"github.com/shivamvishwakarm/resume-matcher/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var round = 2

func CreateUser(user models.User) (models.User, error) {

	coll := config.Client.Database(config.DbName).Collection(config.ColName)

	// Check if user already exists
	filter := bson.M{"email": user.Email}

	var existingUser models.User
	err := coll.FindOne(context.TODO(), filter).Decode(&existingUser)

	if err == nil {
		return models.User{}, errors.New("user already exists")
	}

	if err != mongo.ErrNoDocuments {
		return models.User{}, err
	}

	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), round)
	if err != nil {
		return models.User{}, err
	}

	// Todo: have to findout what is the best way to remove password field in the response
	// Document to insert
	doc := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return models.User{}, err
	}

	// Optional: set inserted ID if using Mongo ObjectID
	doc.ID = result.InsertedID.(bson.ObjectID)

	return doc, nil
}

type LoginReq struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password" `
}

type LoginRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type User struct {
	ID       bson.ObjectID `bson:"_id"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

func LoginUser(user LoginReq) (LoginRes, error) {

	coll := config.Client.Database(config.DbName).Collection(config.ColName)

	filter := bson.M{"email": user.Email}

	var userDB User
	err := coll.FindOne(context.TODO(), filter).Decode(&userDB)
	if err != nil {
		return LoginRes{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(userDB.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return LoginRes{}, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  userDB.Name,
		"email": userDB.Email,
	})

	tokenString, err := token.SignedString([]byte("verysecret123")) //todo: should comes from .env
	if err != nil {
		return LoginRes{}, err
	}

	return LoginRes{
		Name:  userDB.Name,
		Email: userDB.Email,
		Token: tokenString,
	}, nil
}
