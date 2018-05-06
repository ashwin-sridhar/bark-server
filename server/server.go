package main
import ("fmt"
"github.com/gorilla/mux"
"log"
"net/http"
"github.com/bark-server/routes")


func main(){
	fmt.Println("Testing iteration50")
	router:=mux.NewRouter()
	router.HandleFunc("/getusers/", routes.GetUsers).Methods("GET")
	router.HandleFunc("/getposts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/getpostsnearme/{radius}/{loclon},{loclat}", routes.GetPostsNearMe).Methods("GET")
	router.HandleFunc("/createpost",routes.CreatePost).Methods("POST")
	router.HandleFunc("/createpost",routes.EnableCors).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":9000", (router)))

}



