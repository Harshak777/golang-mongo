package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/harshak777/golang-mongo/models"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectId(id)

	userObject := models.User{}

	if err := uc.session.DB("golang").C("users").FindId(oid).One(&userObject); err != nil {
		w.WriteHeader(404)
		return
	}

	userJsonObject, err := json.Marshal(userObject)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Print(w, "%s\n", userJsonObject)
}
