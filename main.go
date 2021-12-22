package main

import(
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)
func main(){
	
	godotenv.Load()

	port := os.Getenv("PORT")
	// port = "8080"
    s := http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/")))
	r := mux.NewRouter()
    r.PathPrefix("/ui/").Handler(s)
	r.HandleFunc("/", Home).Methods("GET", "POST")
	r.HandleFunc("/about", About).Methods("GET")
	r.HandleFunc("/techno", Techno).Methods("GET")
	r.HandleFunc("/projects", Projects).Methods("GET")

	r.HandleFunc("/admin", Loggin).Methods("GET", "POST")
	r.HandleFunc("/admin/comments", GetAll).Methods("GET", "VIEW")
	r.HandleFunc("/admin/comments/{commentId}", GetById).Methods("GET", "VIEW")

	log.Println("Server ready at :",port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

