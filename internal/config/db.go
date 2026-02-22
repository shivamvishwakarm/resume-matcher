package config

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shivamvishwakarm/resume-matcher/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Exported package-level variables — accessible as config.Db, config.Client, etc.
var (
	Db       *mongo.Database
	Client   *mongo.Client
	MongoCtx context.Context
	Cancel   context.CancelFunc
)

const MONGO_URI = "mongodb://admin:password@localhost:27017/?authSource=admin"
const dbName = "resume-matcher"
const colName = "user"

const round = 8

func init() {

	client, err := mongo.Connect(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping to verify the connection is alive
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	Client = client
	Db = client.Database(dbName)
	MongoCtx = context.Background()

	log.Println("Connected to MongoDB successfully")
}
func CreateUser(user models.User) (models.User, error) {

	coll := Client.Database(dbName).Collection(colName)

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

	coll := Client.Database(dbName).Collection(colName)

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
