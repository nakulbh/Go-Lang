package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nakulbh/Go-Lang/Go-Lang/go-mongodb/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserControllers(getSession)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":9000", r)

}

// this function helps cpnnect to mongodb
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
