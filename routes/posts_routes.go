package routes

import (
	"fmt"
)

import("net/http"
	"encoding/json"
"github.com/ashwin-sridhar/database"
"io"
"io/ioutil"
"log"
"github.com/gorilla/mux"
"strconv")

var rep =database.Repository{}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rep.GetPosts())
}

func GetPostsNearMe(w http.ResponseWriter, r *http.Request) {
	requestparams := mux.Vars(r)
	fmt.Print(requestparams["loclon"])
	fmt.Print(requestparams["loclat"])
	loclon, err_loclon := strconv.ParseFloat(requestparams["loclon"], 64)
	loclat, err_loclat := strconv.ParseFloat(requestparams["loclat"], 64)
	radius,err_radius:=strconv.Atoi(requestparams["radius"]);
	if (err_loclon!=nil)||(err_loclat!=nil)||(err_radius)!=nil{
		w.WriteHeader(400)
		return
	}else{
		json.NewEncoder(w).Encode(rep.GetPostsFromNearby(radius,loclon,loclat))
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post database.PostData
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
	 log.Fatalln("Error AddAlbum", err)
	 w.WriteHeader(http.StatusInternalServerError)
	 return
	}
	if err := r.Body.Close(); err != nil {
	 log.Fatalln("Error AddAlbum", err)
	}
	if err := json.Unmarshal(body, &post); err != nil { // unmarshall body contents as a type Candidate
	 w.WriteHeader(422) // unprocessable entity
	 if err := json.NewEncoder(w).Encode(err); err != nil {
	  log.Fatalln("Error AddAlbum unmarshalling data", err)
	  w.WriteHeader(http.StatusInternalServerError)
	  return
	 }
	}
	 success := rep.CreatePost(post) // adds the post to the DB
	 if !success {
	  w.WriteHeader(http.StatusInternalServerError)
	  return
	 }
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusCreated)
	 return
	}