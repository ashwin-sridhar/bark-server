package routes

import(
	"fmt"
	"net/http"
	"encoding/json"
"github.com/bark-server/database")

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rep:=database.Repository{}
	fmt.Print("RECO")
	json.NewEncoder(w).Encode(rep.GetUsers())
}