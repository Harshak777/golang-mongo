package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/harshak777/golang-mongo/models"
)

type UserController struct {
	session *mongo.Client
	ctx     context.Context
}

func NewUserController(session *mongo.Client, ctx context.Context) *UserController {
	return &UserController{session, ctx}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	userObject := models.User{}
	filter := bson.M{"_id": oid}

	if err := uc.session.Database("golang").Collection("users").FindOne(uc.ctx, filter).Decode(&userObject); err != nil {
		w.WriteHeader(404)
		return
	}

	userJsonObject, err := json.Marshal(userObject)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprint(w, string(userJsonObject))
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userObject := models.User{}

	json.NewDecoder(r.Body).Decode(&userObject)

	userObject.Id = primitive.NewObjectID()

	uc.session.Database("golang").Collection("users").InsertOne(uc.ctx, userObject)

	userJsonObject, err := json.Marshal(userObject)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, "\n", string(userJsonObject))
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	filter := bson.M{"_id": oid}
	if uc.session.Database("golang").Collection("users").DeleteOne(uc.ctx, filter); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user ", oid, "\n")
}
