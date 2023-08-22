package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nakulbh/Go-Lang/Go-Lang/go-mongodb/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserControllers struct {
	session *mgo.Session
}

func NewUserControllers(s *mgo.Session) *UserControllers {
	return &UserControllers{s}
}

func (uc UserControllers) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	/*     uc.Session accesses the MongoDB session.
	.DB("mongo-golang") specifies that you want to work with the database named "mongo-golang". It's like selecting a database in MongoDB.*/

	if err := uc.Session.DB("mongo-golang").C("users").FindID(oid).One(&u); err != nill {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nill {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
func (uc UserControllers) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.Session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nill {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", u)
}
