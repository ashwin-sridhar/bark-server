package database

import("gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
"fmt")
const USERSDOCNAME = "users"
type UserData struct{
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	Username   string `json:"username" bson:"username"`
	Email      string `json:"email" bson:"email"`
}

// GetUsers returns the list of Users
func (r Repository) GetUsers() []UserData {
	session, err := mgo.Dial(SERVER)
	if err != nil {
	 fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(USERSDOCNAME)
	users := make([]UserData, 10, 30)
	if err := c.Find(nil).All(&users); err != nil {
	 fmt.Println("Failed to write results:", err)
	}
   return users
   }