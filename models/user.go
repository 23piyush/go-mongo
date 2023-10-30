package models

import "gopkg.in/mgo.v2/bson"

// struct is used for user defined datatype
// defines a User struct that represents a user entity.
type User struct{
	Id		bson.ObjectId	`json:"id" bson:"_id"` // Id will be created automatically for each entry
	Name	string			`json:"name" bson:"name"`
	Gender	string			`json:"gender" bson:"gender"`
	Age		int				`json:"age" bson:"age"`
}
// User is like glu-layer between database and golang program
// bson: format understood by mongodb. mongodb generates id for each entry with column name as "_id"
// json is what we send from postman through request to golang
// In golang functions, we will use "Id"
// While sending from postman, we will use "id"
// "_id" is what goes to database
// bson.ObjectId - It represents the unique identifier used by MongoDB.
// The JSON tags (json:"...") are used to specify the field names when encoding and decoding JSON data. 
// The BSON tags (bson:"...") are used to specify the field names when working with MongoDB.