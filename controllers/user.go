package controllers

import(
"fmt"
"encoding/json" // golang doesn't understands json : Whatever data we get from POST request needs to be encoded first
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"github.com/23piyush/go-mongo/models"
"net/http" // to work with http requests in golang
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
return &UserController{s}
}

// GetUser - a struct method, not a function
func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound) // 404
		// w is the response we are sending back to postman, or front-end whoever is sending request to golang
		// we can send header with response, which has status codes
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{} // empty struct, we will store data returned from database in it later

	// "mongo-golang" is database name
	// "users" is collection in the database
	// mongoose -  MongoDB object modeling tool designed to work in an asynchronous environment. Mongoose supports Node.js and Deno (alpha).
	// There is nothing in golang similar to go
  if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil{
	w.WriteHeader(404)
	return
   }

   uj, err :=json.Marshal(u) // Marshal() - converts to json
   if err!= nil{
	   fmt.Println(err)
   }

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK) // 200
fmt.Fprintf(w, "%s\n", uj) // we will see this printed in postman

}

// we use _ when don't use this further in code
// a struct method, not a function
func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}

	// decode the json values received from postman so that golang can understand, and get it into "u"
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u) // convert back to json before sending back to postman

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}
// We are creating data in database and returning it to postman

// We need params to get id to delete user from database
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
