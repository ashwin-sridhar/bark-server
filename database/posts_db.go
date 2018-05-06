package database

import("gopkg.in/mgo.v2/bson"
	   "gopkg.in/mgo.v2"
	   "fmt"
	   "log")
const POSTSDOCNAME = "posts"
type PostData struct{
	ID     bson.ObjectId `bson:"_id"`
	Description string `json:"description" bson:"description"`
	Status		string `json:"status" bson:"status"`
	Severity    string `json:"severity" bson:"severity"`
	Location    GeoJson `json:"location" bson:"location"`
	Author		UserData `json:"author" bson:"author"`
	Responder	UserData  `json:"responder,omitempty" bson:"responder,omitempty"`
}

type GeoJson struct{
	Type string `json:"-"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

// GetPosts returns all the posts
func (r Repository) GetPosts() []PostData {
	session, err := mgo.Dial(SERVER)
	if err != nil {
	 fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(POSTSDOCNAME)
	posts := make([]PostData, 10, 30)
	if err := c.Find(nil).All(&posts); err != nil {
	 fmt.Println("Failed to write results:", err)
	}
   return posts
   }

// GetPosts based on proximity
func (r Repository) GetPostsFromNearby(radius int,loclon float64,loclat float64) []PostData {
	session, err := mgo.Dial(SERVER)
	if err != nil {
	 fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(POSTSDOCNAME)
	posts := make([]PostData, 10, 30)
	session.SetMode(mgo.Monotonic, true)
    // Creating the index
    index := mgo.Index{
        Key:  []string{"$2dsphere:location"},
        Bits: 26,
    }
    err = c.EnsureIndex(index)
    if err != nil {
        fmt.Println("Indexing error")
        panic(err)
    }
		//Querying
		err = c.Find(bson.M{
			"location": bson.M{
				"$nearSphere": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": []float64{loclon, loclat},
					},
					"$maxDistance":radius,
				},
			},
		}).All(&posts)
		if err != nil {
			panic(err)
		}
	//ENDS
   return posts
   }

// CreatePost inserts a given post in to the DB
func (r Repository) CreatePost(post PostData) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
   post.ID = bson.NewObjectId()
	session.DB(DBNAME).C(POSTSDOCNAME).Insert(post)
   if err != nil {
	 log.Fatalln("Insert failure",err)
	 return false
	}
	fmt.Println("DONE OH DONE")
	return true
   }