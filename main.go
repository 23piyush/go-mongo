package main

import (
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"net/http" // helps to create server using ListenAndServe() function

"github.com/23piyush/go-mongo/controllers"
)

func main(){

	r := httprouter.New() 
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser) // GetUser is the function that will handle this route
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}


func getSession() *mgo.Session{
	
	s, err := mgo.Dial("mongodb://localhost:27017") // connection string helps to connect with MongoDB server
	if err != nil{
		panic(err)
	}
	return s
}

// In Go, := is for declaration + assignment, whereas = is for assignment only.
// For example, var foo int = 10 is the same as foo := 10.
// ListenAndServe() creates a golang server for us. As an input, we pass the ip and post where we want the server to run
// We also pass httprouter variable as parameter to specify who will handle the routes 

// This file sets up a server using the httprouter package and connects to a MongoDB server using the mgo package. 
// It defines routes for handling HTTP requests related to user operations.